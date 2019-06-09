package bot

import (
	"database/sql"

	"github.com/hortbot/hortbot/internal/pkg/dedupe"
	"github.com/leononame/clock"
)

const (
	DefaultPrefix = "!"
	DefaultBullet = "[HB]"
)

type Config struct {
	DB       *sql.DB
	Dedupe   dedupe.Deduplicator
	Sender   Sender
	Notifier Notifier
	Clock    clock.Clock

	Prefix   string
	Bullet   string
	Cooldown int

	Admins []string

	WhitelistEnabled bool
	Whitelist        []string
}

type Bot struct {
	db       *sql.DB
	dedupe   dedupe.Deduplicator
	sender   Sender
	notifier Notifier
	clock    clock.Clock

	prefix   string
	bullet   string
	cooldown int

	admins    map[string]bool
	whitelist map[string]bool

	testingHelper testingHelper
}

func New(config *Config) *Bot {
	// TODO: don't panic, return errors.
	switch {
	case config.DB == nil:
		panic("db is nil")
	case config.Dedupe == nil:
		panic("dedupe is nil")
	case config.Sender == nil:
		panic("sender is nil")
	case config.Notifier == nil:
		panic("notifier is nil")
	}

	b := &Bot{
		db:       config.DB,
		dedupe:   config.Dedupe,
		sender:   config.Sender,
		notifier: config.Notifier,
		prefix:   config.Prefix,
		bullet:   config.Bullet,
		cooldown: config.Cooldown,
		admins:   make(map[string]bool),
	}

	if b.bullet == "" {
		b.bullet = DefaultBullet
	}

	if b.prefix == "" {
		b.prefix = DefaultPrefix
	}

	if config.Clock != nil {
		b.clock = config.Clock
	} else {
		b.clock = clock.New()
	}

	if isTesting {
		b.testingHelper = testingHelper{}
	}

	for _, name := range config.Admins {
		b.admins[name] = true
	}

	if config.WhitelistEnabled {
		b.whitelist = make(map[string]bool)
		for _, name := range config.Whitelist {
			b.whitelist[name] = true
		}
	}

	return b
}

func (b *Bot) isAllowed(name string) bool {
	if b.whitelist == nil {
		return true
	}

	if b.admins[name] {
		return true
	}

	if b.whitelist[name] {
		return true
	}

	return false
}
