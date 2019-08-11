package bot

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/hortbot/hortbot/internal/db/models"
	"github.com/hortbot/hortbot/internal/pkg/ctxlog"
	"github.com/jakebailey/irc"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/queries"
	"go.uber.org/zap"
)

var (
	errInvalidMessage  = errors.New("bot: invalid message")
	errNotAuthorized   = errors.New("bot: user is not authorized to use this command")
	errBuiltinDisabled = errors.New("bot: builtin disabled")
)

func (b *Bot) Handle(ctx context.Context, origin string, m *irc.Message) {
	ctx = withCorrelation(ctx)
	logger := ctxlog.FromContext(ctx)

	if m == nil {
		logger.Error("nil message")
		return
	}

	ctx, logger = ctxlog.FromContextWith(ctx, zap.String("irc_command", m.Command))

	var start time.Time

	if !isTesting {
		start = time.Now()
	}

	err := b.handle(ctx, origin, m)

	if !isTesting {
		logger.Debug("handled message", zap.Duration("took", time.Since(start)))
	}

	if err != nil {
		logger.Error("error during handle", zap.Error(err), zap.Any("message", m))
	}
}

func (b *Bot) handle(ctx context.Context, origin string, m *irc.Message) error {
	start := b.deps.Clock.Now()

	if m.Command != "PRIVMSG" {
		// TODO: handle other types of messages
		return nil
	}

	if len(m.Tags) == 0 {
		return errInvalidMessage
	}

	if len(m.Params) == 0 {
		return errInvalidMessage
	}

	id := m.Tags["id"]
	if id == "" {
		return errInvalidMessage
	}

	logger := ctxlog.FromContext(ctx)

	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(testingPanic); ok {
				panic(r)
			}
			logger.Error("panic during handle", zap.Any("value", r), zap.Stack("stack"))
		}
	}()

	seen, err := b.deps.Dedupe.CheckAndMark(id)
	if err != nil {
		logger.Error("error checking for duplicate", zap.Error(err), zap.String("id", id))
		return err
	}

	if seen {
		logger.Debug("message already seen", zap.String("id", id))
		return nil
	}

	user := strings.ToLower(m.Prefix.Name)

	if !b.deps.IsAllowed(user) {
		return nil
	}

	message := m.Trailing

	if message == "" {
		return nil
	}

	me := false
	if c, a, ok := irc.ParseCTCP(message); ok {
		if c != "ACTION" {
			logger.Warn("unknown CTCP", zap.String("ctcpCommand", c), zap.String("ctcpArgs", a))
			return nil
		}

		message = a
		me = true
	}

	message = strings.TrimSpace(message)

	if message == "" {
		return nil
	}

	s := &session{
		Type:    sessionNormal,
		Origin:  origin,
		M:       m,
		Deps:    b.deps,
		Start:   start,
		ID:      id,
		User:    user,
		Message: message,
		Me:      me,
	}

	if displayName := m.Tags["display-name"]; displayName != "" {
		s.UserDisplay = displayName
	} else {
		s.UserDisplay = s.User
	}

	roomID := m.Tags["room-id"]
	if roomID == "" {
		logger.Debug("no room ID")
		return errInvalidMessage
	}
	s.RoomIDStr = roomID

	s.RoomID, err = strconv.ParseInt(roomID, 10, 64)
	if err != nil {
		logger.Debug("error parsing room ID", zap.String("parsed", roomID), zap.Error(err))
		return err
	}

	if s.RoomID == 0 {
		logger.Debug("room ID cannot be zero")
		return errInvalidMessage
	}

	userID := m.Tags["user-id"]
	if userID == "" {
		logger.Debug("no user ID")
		return errInvalidMessage
	}

	s.UserID, err = strconv.ParseInt(userID, 10, 64)
	if err != nil {
		logger.Debug("error parsing user ID", zap.String("parsed", userID), zap.Error(err))
		return err
	}

	if s.UserID == 0 {
		logger.Debug("user ID cannot be zero")
		return errInvalidMessage
	}

	tmiSentTs, _ := strconv.ParseInt(m.Tags["tmi-sent-ts"], 10, 64)
	s.TMISent = time.Unix(tmiSentTs/1000, 0)

	// TODO: atomic locking on the channel

	channelName := m.Params[0]
	if channelName == "" || channelName[0] != '#' || len(channelName) == 1 {
		logger.Debug("bad channel name", zap.Strings("params", m.Params))
		return errInvalidMessage
	}

	s.IRCChannel = channelName[1:]

	b.testingHelper.checkUserNameID(s.User, s.UserID)
	b.testingHelper.checkUserNameID(s.IRCChannel, s.RoomID)

	if s.User == s.Origin {
		return nil
	}

	ctx, _ = ctxlog.FromContextWith(ctx,
		zap.Int64("roomID", s.RoomID),
		zap.String("channel", s.IRCChannel),
	)

	return transact(b.db, func(tx *sql.Tx) error {
		s.Tx = tx
		return handleSession(ctx, s)
	})
}

func handleSession(ctx context.Context, s *session) error {
	logger := ctxlog.FromContext(ctx)

	s.SetUserLevel()

	if s.Origin == s.IRCChannel {
		return handleManagement(ctx, s)
	}

	// This is the most frequent query; speed it up by executing a hand written query.
	channel := &models.Channel{}
	err := queries.Raw(`SELECT * FROM channels WHERE user_id = $1`, s.RoomID).Bind(ctx, s.Tx, channel)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Debug("channel not found in database")
			return nil
		}
		return err
	}

	if !channel.Active {
		logger.Warn("channel is not active")
		return nil
	}

	if channel.Name != s.IRCChannel {
		logger.Error("channel name mismatch", zap.String("fromMessage", s.IRCChannel), zap.String("fromDB", channel.Name))
		return errors.New("channel name mismatch") // TODO
	}

	if channel.BotName != s.Origin {
		logger.Warn("bot name mismatch",
			zap.String("expected", channel.BotName),
			zap.String("origin", s.Origin),
		)
		return nil
	}

	s.Channel = channel
	s.SetUserLevel()

	_, ignored := stringSliceIndex(channel.Ignored, s.User)

	if ignored {
		if s.IsAdmin() || s.IRCChannel == s.User {
			ignored = false
		} else {
			s.UserLevel = levelEveryone
		}
	}

	if filtered, err := tryFilter(ctx, s); filtered || err != nil {
		return err
	}

	// Ignoring does not exempt messages from filters.

	if ignored {
		return nil
	}

	if !s.UserLevel.CanAccessPG(s.Channel.Mode) {
		return nil
	}

	s.N, err = s.IncrementMessageCount()
	if err != nil {
		return err
	}

	wasCommand, err := tryCommand(ctx, s)
	if wasCommand {
		switch err {
		case errNotAuthorized, errBuiltinDisabled, errInCooldown:
		default:
			return err
		}
	}

	if s.Channel.ParseYoutube && s.Deps.YouTube != nil {
		for _, u := range s.Links() {
			title := s.Deps.YouTube.VideoTitle(u)
			if title != "" {
				return s.Replyf("Linked YouTube video: \"%s\"", title)
			}
		}
	}

	if ok, err := tryAutoreplies(ctx, s); ok || err != nil {
		return err
	}

	return nil
}

func tryCommand(ctx context.Context, s *session) (bool, error) {
	if s.Me {
		return false, nil
	}

	channel := s.Channel
	prefix := channel.Prefix
	message := s.Message

	name, params := splitSpace(message)

	if s.Channel.ShouldModerate {
		ok, err := moderationCommands.Run(ctx, s, strings.ToLower(name), params)
		switch {
		case err == errNotAuthorized:
		case err != nil:
			return true, err
		case ok:
			return true, nil
		}
	}

	if !strings.HasPrefix(name, prefix) {
		return false, nil
	}

	name = cleanCommandName(name[len(prefix):])

	if name == "" {
		return false, nil
	}

	s.CommandParams = params
	s.OrigCommandParams = params

	ctx, logger := ctxlog.FromContextWith(ctx, zap.String("name", name), zap.String("params", params))

	infoAndCommand := struct {
		CommandInfo models.CommandInfo `boil:"command_infos,bind"`
		Message     null.String        `boil:"message"`
	}{}

	// This is much faster than using qm.Load, as SQLBoiler's loading does multiple
	// queries to fetch 1:1 relationships rather than joins.
	err := queries.Raw(`
		SELECT command_infos.*, custom_commands.message
		FROM command_infos
		LEFT OUTER JOIN custom_commands on custom_commands.id = command_infos.custom_command_id
		WHERE ("command_infos"."channel_id" = $1) AND ("command_infos"."name" = $2)
		FOR UPDATE OF command_infos
		`, s.Channel.ID, name).Bind(ctx, s.Tx, &infoAndCommand)

	info := &infoAndCommand.CommandInfo

	switch err {
	case sql.ErrNoRows:
		return tryBuiltinCommand(ctx, s, name, params)
	case nil:
	default:
		logger.Error("error looking up command name in database", zap.Error(err))
		return true, err
	}

	if !s.UserLevel.CanAccessPG(info.AccessLevel) {
		return false, errNotAuthorized
	}

	if msg := infoAndCommand.Message; msg.Valid {
		return handleCustomCommand(ctx, s, info, msg.String)
	}

	return handleList(ctx, s, info)
}

func tryBuiltinCommand(ctx context.Context, s *session, cmd string, args string) (bool, error) {
	if cmd == "builtin" {
		cmd, args = splitSpace(args)
		cmd = cleanCommandName(cmd)
	}
	return builtinCommands.RunWithCooldown(ctx, s, cmd, args)
}
