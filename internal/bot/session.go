package bot

import (
	"database/sql"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/hortbot/hortbot/internal/db/models"
	"github.com/jakebailey/irc"
	"github.com/leononame/clock"
)

type Session struct {
	Origin string
	M      *irc.Message

	Bot      *Bot
	Tx       *sql.Tx
	Sender   Sender
	Notifier Notifier
	Clock    clock.Clock

	Start   time.Time
	TMISent time.Time

	ID         string
	RoomID     int64
	IRCChannel string
	Message    string
	Me         bool

	User        string
	UserDisplay string
	UserID      int64
	UserLevel   AccessLevel
	Ignored     bool

	Links []*url.URL

	Channel *models.Channel

	CommandParams     string
	OrigCommandParams string
}

func (s *Session) formatResponse(response string) string {
	response = strings.TrimSpace(response)

	if len(response) >= 4 {
		switch response[:4] {
		case "/me ", ".me ":
			return response
		}
	}

	bullet := s.Bot.bullet

	if s.Channel != nil && s.Channel.Bullet.Valid {
		bullet = s.Channel.Bullet.String
	}

	return bullet + " " + response
}

func (s *Session) Reply(response string) error {
	return s.Sender.SendMessage(s.Origin, "#"+s.IRCChannel, s.formatResponse(response))
}

func (s *Session) Replyf(format string, args ...interface{}) error {
	response := fmt.Sprintf(format, args...)
	return s.Reply(response)
}

func (s *Session) ReplyUsage(usage string) error {
	return s.Replyf("usage: %s%s", s.Channel.Prefix, usage)
}

func (s *Session) SetUserLevel() {
	s.UserLevel = s.parseUserLevel()
}

func (s *Session) parseUserLevel() AccessLevel {
	if s.Bot.admins[s.User] {
		return LevelAdmin
	}

	// Tags are present, safe to not check for nil

	tags := s.M.Tags

	if isTesting && tags["testing-admin"] != "" {
		return LevelAdmin
	}

	if s.User == s.IRCChannel {
		return LevelBroadcaster
	}

	if tags["mod"] == "1" {
		return LevelModerator
	}

	badges := parseBadges(tags["badges"])

	switch {
	case badges["broadcaster"] != "":
		return LevelBroadcaster
	case badges["moderator"] != "":
		return LevelModerator
	case badges["subscriber"] != "", badges["vip"] != "", tags["subscriber"] == "1":
		return LevelSubscriber
	}

	if tags["user-type"] == "mod" {
		return LevelModerator
	}

	if s.Channel != nil {
		for _, owner := range s.Channel.CustomOwners {
			if s.User == owner {
				return LevelBroadcaster
			}
		}

		for _, mod := range s.Channel.CustomMods {
			if s.User == mod {
				return LevelModerator
			}
		}

		for _, reg := range s.Channel.CustomRegulars {
			if s.User == reg {
				return LevelSubscriber
			}
		}
	}

	return LevelEveryone
}

func (s *Session) IsAdmin() bool {
	return s.UserLevel.CanAccess(LevelAdmin)
}

func (s *Session) IsInCooldown() bool {
	seconds := s.Clock.Since(s.Channel.LastCommandAt).Seconds()
	cooldown := s.Bot.cooldown

	if s.Channel.Cooldown.Valid {
		cooldown = s.Channel.Cooldown.Int
	}

	return seconds < float64(cooldown)
}
