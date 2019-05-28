package bot

import (
	"context"
)

type builtinCommand struct {
	fn       func(ctx context.Context, s *Session, cmd string, args string) error
	minLevel AccessLevel
}

func (b builtinCommand) run(ctx context.Context, s *Session, cmd string, args string) error {
	if !s.UserLevel.CanAccess(b.minLevel) {
		return errNotAuthorized
	}

	return b.fn(ctx, s, cmd, args)
}

var builtinCommands map[string]builtinCommand

func init() {
	// To prevent initialization loop.

	builtinCommands = map[string]builtinCommand{
		"command": {fn: cmdSimpleCommand, minLevel: LevelModerator},
		"coemand": {fn: cmdSimpleCommand, minLevel: LevelModerator},
		"bullet":  {fn: cmdBullet, minLevel: LevelBroadcaster},
		"prefix":  {fn: cmdPrefix, minLevel: LevelBroadcaster},
		"owner":   {fn: cmdOwnerModRegularIgnore, minLevel: LevelBroadcaster},
		"mod":     {fn: cmdOwnerModRegularIgnore, minLevel: LevelBroadcaster},
		"regular": {fn: cmdOwnerModRegularIgnore, minLevel: LevelModerator},
		"ignore":  {fn: cmdOwnerModRegularIgnore, minLevel: LevelModerator},
	}
}

var reservedCommandNames = map[string]bool{
	"builtin": true,
	"command": true,
}
