package web

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gofrs/uuid"
	"github.com/gorilla/sessions"
	"github.com/hortbot/hortbot/internal/confimport"
	"github.com/hortbot/hortbot/internal/db/models"
	"github.com/hortbot/hortbot/internal/db/modelsx"
	"github.com/hortbot/hortbot/internal/db/redis"
	"github.com/hortbot/hortbot/internal/pkg/apis/twitch"
	"github.com/hortbot/hortbot/internal/pkg/ctxlog"
	"github.com/hortbot/hortbot/internal/pkg/jsonx"
	"github.com/hortbot/hortbot/internal/web/mid"
	"github.com/hortbot/hortbot/internal/web/static"
	"github.com/hortbot/hortbot/internal/web/templates"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tomwright/queryparam/v4"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"go.uber.org/zap"
)

var botScopes = []string{
	"user_follows_edit",
	"channel:moderate",
	"chat:edit",
	"chat:read",
	"whispers:read",
	"whispers:edit",
}

type App struct {
	Addr       string
	RealIP     bool
	SessionKey []byte
	AdminAuth  map[string]string

	Brand    string
	BrandMap map[string]string

	Debug bool

	Redis  *redis.DB
	DB     *sql.DB
	Twitch *twitch.Twitch

	store *sessions.CookieStore
}

func (a *App) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if len(a.SessionKey) == 0 {
		panic("empty session key")
	}

	a.store = sessions.NewCookieStore(a.SessionKey)

	r := chi.NewRouter()

	logger := ctxlog.FromContext(ctx)
	r.Use(mid.Logger(logger))
	r.Use(mid.RequestID)

	if a.RealIP {
		r.Use(middleware.RealIP)
	}

	r.Use(func(next http.Handler) http.Handler {
		return promhttp.InstrumentHandlerCounter(metricRequest, next)
	})

	r.Use(mid.RequestLogger)
	r.Use(mid.Tracer)
	r.Use(mid.Recoverer)

	r.Group(func(r chi.Router) {
		r.Use(middleware.RedirectSlashes)

		r.Get("/", a.index)
		r.Get("/about", a.about)
		r.Get("/docs", a.docs)
		r.Get("/channels", a.channels)

		const paramChannel = "channel"
		r.Route("/c/{"+paramChannel+"}", func(r chi.Router) {
			r.Use(func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					p := r.URL.Path
					lp := strings.ToLower(p)

					if p == lp {
						next.ServeHTTP(w, r)
						return
					}

					if r.URL.RawQuery != "" {
						lp += "?" + r.URL.RawQuery
					}

					http.Redirect(w, r, lp, http.StatusMovedPermanently)
				})
			})

			r.Use(a.channelMiddleware(paramChannel))
			r.Get("/", a.channel)
			r.Get("/commands", a.channelCommands)
			r.Get("/quotes", a.channelQuotes)
			r.Get("/autoreplies", a.channelAutoreplies)
			r.Get("/lists", a.channelLists)
			r.Get("/regulars", a.channelRegulars)
			r.Get("/chatrules", a.channelChatRules)
			r.Get("/scheduled", a.channelScheduled)
		})

		r.Get("/login", a.login)
		r.Get("/logout", a.logout)
		r.Get("/auth/twitch", a.authTwitchNormal)
		r.Get("/auth/twitch/bot", a.authTwitchBot)
		r.Get("/auth/twitch/callback", a.authTwitchCallback)

		routeDebug := func(r chi.Router) {
			r.Use(middleware.NoCache)
			r.Get("/request", dumpRequest)
		}

		if a.Debug {
			r.Route("/debug", routeDebug)
		}

		r.Route("/admin", func(r chi.Router) {
			r.Use(middleware.NoCache)
			r.Use(a.adminAuth)

			r.Route("/debug", routeDebug)

			r.Get("/import", a.adminImport)
			r.Post("/import", a.adminImportPost)
			r.Get("/export/{channel}", a.adminExport)
		})
	})

	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(static.FS(false))))
	r.Handle("/favicon.ico", http.RedirectHandler("/static/icons/favicon.ico", http.StatusFound))

	srv := http.Server{
		Addr:    a.Addr,
		Handler: r,
	}

	go func() {
		<-ctx.Done()
		if err := srv.Shutdown(context.Background()); err != nil {
			ctxlog.Error(ctx, "error shutting down server", zap.Error(err))
		}
	}()

	ctxlog.Info(ctx, "web server listening", zap.String("addr", srv.Addr))

	return srv.ListenAndServe()
}

func (a *App) getBrand(r *http.Request) string {
	if a.BrandMap == nil {
		return a.Brand
	}

	host, _, _ := net.SplitHostPort(r.Host)
	host = normalizeHost(host)

	if host != "" {
		if brand := a.BrandMap[host]; brand != "" {
			return brand
		}
	}

	return a.Brand
}

type authState struct {
	Host     string
	Bot      bool
	Redirect string
}

func (a *App) authTwitch(w http.ResponseWriter, r *http.Request, bot bool) {
	ctx := r.Context()

	state := uuid.Must(uuid.NewV4()).String()

	stateVal := &authState{
		Host: r.Host, // Not normalized; needed for redirects.
		Bot:  bot,
	}

	query := struct {
		Redirect string `queryparam:"redirect"`
	}{}

	if err := queryparam.Parse(r.URL.Query(), &query); err != nil {
		httpError(w, http.StatusBadRequest)
		return
	}

	stateVal.Redirect = query.Redirect

	if err := a.Redis.SetAuthState(r.Context(), state, stateVal, time.Minute); err != nil {
		ctxlog.Error(ctx, "error setting auth state", zap.Error(err))
		httpError(w, http.StatusInternalServerError)
		return
	}

	var extraScopes []string
	if bot {
		extraScopes = botScopes
	}

	url := a.Twitch.AuthCodeURL(state, extraScopes...)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (a *App) authTwitchNormal(w http.ResponseWriter, r *http.Request) {
	a.authTwitch(w, r, false)
}

func (a *App) authTwitchBot(w http.ResponseWriter, r *http.Request) {
	a.authTwitch(w, r, true)
}

func (a *App) authTwitchCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	state := r.FormValue("state")
	if state == "" {
		httpError(w, http.StatusBadRequest)
		return
	}

	var stateVal authState

	ok, err := a.Redis.GetAuthState(ctx, state, &stateVal)
	if err != nil {
		ctxlog.Error(ctx, "error checking auth state", zap.Error(err))
		httpError(w, http.StatusInternalServerError)
		return
	}

	if !ok {
		httpError(w, http.StatusBadRequest)
		return
	}

	if normalizeHost(stateVal.Host) != normalizeHost(r.Host) {
		// This came to the wrong host. Put the state back and redirect.
		if err := a.Redis.SetAuthState(r.Context(), state, &stateVal, time.Minute); err != nil {
			ctxlog.Error(ctx, "error setting auth state", zap.Error(err))
			httpError(w, http.StatusInternalServerError)
			return
		}

		u := *r.URL
		u.Host = stateVal.Host
		templates.WriteMetaRedirect(w, u.String())
		return
	}

	tok, err := a.Twitch.Exchange(ctx, r.FormValue("code"))
	if err != nil {
		ctxlog.Error(ctx, "error exchanging code", zap.Error(err))
		httpError(w, http.StatusBadRequest)
		return
	}

	user, newToken, err := a.Twitch.GetUserForToken(ctx, tok)
	if err != nil {
		ctxlog.Error(ctx, "error getting user for token", zap.Error(err))
		httpError(w, http.StatusInternalServerError)
		return
	}
	if newToken != nil {
		tok = newToken
	}

	tt := modelsx.TokenToModel(user.ID, tok)
	if stateVal.Bot {
		tt.BotName = null.StringFrom(user.Name)
	}

	if err := modelsx.FullUpsertToken(ctx, a.DB, tt); err != nil {
		ctxlog.Error(ctx, "error upserting token", zap.Error(err))
		httpError(w, http.StatusInternalServerError)
		return
	}

	if err := a.clearSession(w, r); err != nil {
		ctxlog.Error(ctx, "error saving session", zap.Error(err))
		httpError(w, http.StatusInternalServerError)
		return
	}

	session := a.getSession(r)
	session.setTwitchID(user.ID)
	session.setUsername(user.Name)

	if err := session.save(w, r); err != nil {
		ctxlog.Error(ctx, "error saving session", zap.Error(err))
		httpError(w, http.StatusInternalServerError)
		return
	}

	if stateVal.Redirect != "" {
		http.Redirect(w, r, stateVal.Redirect, http.StatusSeeOther)
		return
	}

	page := &templates.LoginSuccessPage{
		Name: user.Name,
		ID:   user.ID,
		Bot:  stateVal.Bot,
	}
	page.Brand = a.getBrand(r)
	page.User = user.Name

	templates.WritePageTemplate(w, page)
}

func (a *App) index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	channels, err := models.Channels(models.ChannelWhere.Active.EQ(true)).Count(ctx, a.DB)
	if err != nil {
		ctxlog.Error(ctx, "error querying channels", zap.Error(err))
		httpError(w, http.StatusInternalServerError)
		return
	}

	var row struct {
		BotCount int64
	}

	if err := queries.Raw("SELECT COUNT(DISTINCT bot_name) AS bot_count FROM channels WHERE active").Bind(ctx, a.DB, &row); err != nil {
		ctxlog.Error(ctx, "error querying bot names", zap.Error(err))
		httpError(w, http.StatusInternalServerError)
		return
	}

	page := &templates.IndexPage{
		ChannelCount: channels,
		BotCount:     row.BotCount,
	}
	page.Brand = a.getBrand(r)
	page.User = a.getSession(r).getUsername()

	templates.WritePageTemplate(w, page)
}

func (a *App) about(w http.ResponseWriter, r *http.Request) {
	page := &templates.AboutPage{}
	page.Brand = a.getBrand(r)
	page.User = a.getSession(r).getUsername()
	templates.WritePageTemplate(w, page)
}

func (a *App) docs(w http.ResponseWriter, r *http.Request) {
	page := &templates.DocsPage{}
	page.Brand = a.getBrand(r)
	page.User = a.getSession(r).getUsername()
	templates.WritePageTemplate(w, page)
}

func (a *App) channels(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	channels, err := models.Channels(
		models.ChannelWhere.Active.EQ(true),
		qm.OrderBy(models.ChannelColumns.Name),
	).All(ctx, a.DB)
	if err != nil {
		ctxlog.Error(ctx, "error querying channels", zap.Error(err))
		httpError(w, http.StatusInternalServerError)
		return
	}

	page := &templates.ChannelsPage{
		Channels: channels,
	}
	page.Brand = a.getBrand(r)
	page.User = a.getSession(r).getUsername()

	templates.WritePageTemplate(w, page)
}

func (a *App) channel(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	channel := getChannel(ctx)

	page := &templates.ChannelPage{
		Channel: channel,
	}
	page.Brand = a.getBrand(r)
	page.User = a.getSession(r).getUsername()
	templates.WritePageTemplate(w, page)
}

func (a *App) channelCommands(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	channel := getChannel(ctx)

	commands, err := channel.CustomCommands(qm.Load(models.CustomCommandRels.CommandInfo)).All(ctx, a.DB)
	if err != nil {
		ctxlog.Error(ctx, "error querying custom commands", zap.Error(err))
		httpError(w, http.StatusInternalServerError)
		return
	}

	sort.Slice(commands, func(i, j int) bool {
		return commands[i].R.CommandInfo.Name < commands[j].R.CommandInfo.Name
	})

	page := &templates.ChannelCommandsPage{
		Commands: commands,
	}
	page.Brand = a.getBrand(r)
	page.Channel = channel
	page.User = a.getSession(r).getUsername()

	templates.WritePageTemplate(w, page)
}

func (a *App) channelQuotes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	channel := getChannel(ctx)

	quotes, err := channel.Quotes(qm.OrderBy(models.QuoteColumns.Num)).All(ctx, a.DB)
	if err != nil {
		ctxlog.Error(ctx, "error querying quotes", zap.Error(err))
		httpError(w, http.StatusInternalServerError)
		return
	}

	page := &templates.ChannelQuotesPage{
		Quotes: quotes,
	}
	page.Brand = a.getBrand(r)
	page.Channel = channel
	page.User = a.getSession(r).getUsername()

	templates.WritePageTemplate(w, page)
}

func (a *App) channelAutoreplies(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	channel := getChannel(ctx)

	autoreplies, err := channel.Autoreplies(qm.OrderBy(models.AutoreplyColumns.Num)).All(ctx, a.DB)
	if err != nil {
		ctxlog.Error(ctx, "error querying autoreplies", zap.Error(err))
		httpError(w, http.StatusInternalServerError)
		return
	}

	page := &templates.ChannelAutorepliesPage{
		Autoreplies: autoreplies,
	}
	page.Brand = a.getBrand(r)
	page.Channel = channel
	page.User = a.getSession(r).getUsername()

	templates.WritePageTemplate(w, page)
}

func (a *App) channelLists(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	channel := getChannel(ctx)

	lists, err := channel.CommandLists(qm.Load(models.CommandListRels.CommandInfo)).All(ctx, a.DB)
	if err != nil {
		ctxlog.Error(ctx, "error querying command lists", zap.Error(err))
		httpError(w, http.StatusInternalServerError)
		return
	}

	sort.Slice(lists, func(i, j int) bool {
		return lists[i].R.CommandInfo.Name < lists[j].R.CommandInfo.Name
	})

	page := &templates.ChannelListsPage{
		Lists: lists,
	}
	page.Brand = a.getBrand(r)
	page.Channel = channel
	page.User = a.getSession(r).getUsername()

	templates.WritePageTemplate(w, page)
}

func (a *App) channelRegulars(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	channel := getChannel(ctx)

	page := &templates.ChannelRegularsPage{}
	page.Brand = a.getBrand(r)
	page.Channel = channel
	page.User = a.getSession(r).getUsername()

	templates.WritePageTemplate(w, page)
}

func (a *App) channelChatRules(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	channel := getChannel(ctx)

	page := &templates.ChannelRulesPage{}
	page.Brand = a.getBrand(r)
	page.Channel = channel
	page.User = a.getSession(r).getUsername()

	templates.WritePageTemplate(w, page)
}

func (a *App) channelScheduled(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	channel := getChannel(ctx)

	repeated, err := channel.RepeatedCommands(qm.Load(models.RepeatedCommandRels.CommandInfo)).All(ctx, a.DB)
	if err != nil {
		ctxlog.Error(ctx, "error querying repeated commands", zap.Error(err))
		httpError(w, http.StatusInternalServerError)
		return
	}

	scheduled, err := channel.ScheduledCommands(qm.Load(models.ScheduledCommandRels.CommandInfo)).All(ctx, a.DB)
	if err != nil {
		ctxlog.Error(ctx, "error querying scheduled commands", zap.Error(err))
		httpError(w, http.StatusInternalServerError)
		return
	}

	sort.Slice(repeated, func(i, j int) bool {
		if repeated[i].Enabled != repeated[j].Enabled {
			return repeated[i].Enabled
		}

		return repeated[i].R.CommandInfo.Name < repeated[j].R.CommandInfo.Name
	})

	sort.Slice(scheduled, func(i, j int) bool {
		if scheduled[i].Enabled != scheduled[j].Enabled {
			return scheduled[i].Enabled
		}

		return scheduled[i].R.CommandInfo.Name < scheduled[j].R.CommandInfo.Name
	})

	page := &templates.ChannelScheduledPage{
		Repeated:  repeated,
		Scheduled: scheduled,
	}
	page.Brand = a.getBrand(r)
	page.Channel = channel
	page.User = a.getSession(r).getUsername()

	templates.WritePageTemplate(w, page)
}

func (a *App) login(w http.ResponseWriter, r *http.Request) {
	page := &templates.LoginPage{}
	page.Brand = a.getBrand(r)
	page.User = a.getSession(r).getUsername()

	templates.WritePageTemplate(w, page)
}

func (a *App) adminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(a.AdminAuth) == 0 {
			notAuthorized(w, false)
			return
		}

		user, pass, ok := r.BasicAuth()
		if !ok {
			notAuthorized(w, true)
			return
		}

		expected := a.AdminAuth[user]
		if expected == "" || pass != expected {
			notAuthorized(w, true)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a *App) adminExport(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	channelName := chi.URLParam(r, "channel")

	config, err := confimport.ExportByName(ctx, a.DB, strings.ToLower(channelName))
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
		} else {
			ctxlog.Error(ctx, "error exporting channel", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	enc := json.NewEncoder(w)

	query := struct {
		Pretty bool `queryparam:"pretty"`
	}{}

	if err := queryparam.Parse(r.URL.Query(), &query); err != nil {
		httpError(w, http.StatusBadRequest)
		return
	}

	if query.Pretty {
		enc.SetIndent("", "    ")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := enc.Encode(config); err != nil {
		ctxlog.Error(ctx, "error encoding exported config", zap.Error(err))
	}
}

func (a *App) adminImport(w http.ResponseWriter, r *http.Request) {
	page := &templates.AdminImportPage{}
	page.Brand = a.getBrand(r)
	page.User = a.getSession(r).getUsername()
	templates.WritePageTemplate(w, page)
}

func (a *App) adminImportPost(w http.ResponseWriter, r *http.Request) {
	config := &confimport.Config{}

	if err := jsonx.DecodeSingle(r.Body, config); err != nil {
		http.Error(w, "decoding body: "+err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	tx, err := a.DB.BeginTx(ctx, nil)
	if err != nil {
		http.Error(w, "beginning transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rolledBack := false

	defer func() {
		if rolledBack {
			return
		}

		if err := tx.Commit(); err != nil {
			fmt.Fprintln(w, "committing transaction:", err)
		}
	}()

	if err := config.Insert(ctx, tx); err != nil {
		http.Error(w, "inserting config: "+err.Error(), http.StatusBadRequest)
		if err := tx.Rollback(); err != nil {
			fmt.Fprintln(w, "rolling back transaction:", err)
		}
		rolledBack = true
		return
	}

	fmt.Fprintln(w, "Successfully inserted channel", config.Channel.ID)
}

func (a *App) logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := a.clearSession(w, r); err != nil {
		ctxlog.Error(ctx, "error clearing session", zap.Error(err))
		httpError(w, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
