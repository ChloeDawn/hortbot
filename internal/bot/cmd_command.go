package bot

import (
	"context"
	"database/sql"
	"strings"

	"github.com/gobuffalo/flect"
	"github.com/hortbot/hortbot/internal/cbp"
	"github.com/hortbot/hortbot/internal/db/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

var ccCommands = newHandlerMap(map[string]handlerFunc{
	"add":             {fn: cmdCommandAddNormal, minLevel: levelModerator},
	"addb":            {fn: cmdCommandAddBroadcaster, minLevel: levelModerator},
	"addbroadcaster":  {fn: cmdCommandAddBroadcaster, minLevel: levelModerator},
	"addbroadcasters": {fn: cmdCommandAddBroadcaster, minLevel: levelModerator},
	"addo":            {fn: cmdCommandAddBroadcaster, minLevel: levelModerator},
	"addowner":        {fn: cmdCommandAddBroadcaster, minLevel: levelModerator},
	"addowners":       {fn: cmdCommandAddBroadcaster, minLevel: levelModerator},
	"addstreamer":     {fn: cmdCommandAddBroadcaster, minLevel: levelModerator},
	"addstreamers":    {fn: cmdCommandAddBroadcaster, minLevel: levelModerator},
	"addm":            {fn: cmdCommandAddModerator, minLevel: levelModerator},
	"addmod":          {fn: cmdCommandAddModerator, minLevel: levelModerator},
	"addmods":         {fn: cmdCommandAddModerator, minLevel: levelModerator},
	"adds":            {fn: cmdCommandAddSubscriber, minLevel: levelModerator},
	"addsub":          {fn: cmdCommandAddSubscriber, minLevel: levelModerator},
	"addsubs":         {fn: cmdCommandAddSubscriber, minLevel: levelModerator},
	"adde":            {fn: cmdCommandAddEveryone, minLevel: levelModerator},
	"adda":            {fn: cmdCommandAddEveryone, minLevel: levelModerator},
	"addeveryone":     {fn: cmdCommandAddEveryone, minLevel: levelModerator},
	"addall":          {fn: cmdCommandAddEveryone, minLevel: levelModerator},
	"delete":          {fn: cmdCommandDelete, minLevel: levelModerator},
	"remove":          {fn: cmdCommandDelete, minLevel: levelModerator},
	"restrict":        {fn: cmdCommandRestrict, minLevel: levelModerator},
	"editor":          {fn: cmdCommandProperty, minLevel: levelModerator},
	"author":          {fn: cmdCommandProperty, minLevel: levelModerator},
	"count":           {fn: cmdCommandProperty, minLevel: levelModerator},
	"rename":          {fn: cmdCommandRename, minLevel: levelModerator},
	"get":             {fn: cmdCommandGet, minLevel: levelModerator},
	// TODO: clone
})

func cmdCommand(ctx context.Context, s *session, cmd string, args string) error {
	subcommand, args := splitSpace(args)
	subcommand = strings.ToLower(subcommand)

	ok, err := ccCommands.run(ctx, s, subcommand, args)
	if err != nil {
		return err
	}

	if !ok {
		return s.ReplyUsage("add|delete|restrict|...")
	}

	return nil
}

func cmdCommandAddNormal(ctx context.Context, s *session, cmd string, args string) error {
	return cmdCommandAdd(ctx, s, args, levelSubscriber, false)
}

func cmdCommandAddBroadcaster(ctx context.Context, s *session, cmd string, args string) error {
	return cmdCommandAdd(ctx, s, args, levelBroadcaster, true)
}

func cmdCommandAddModerator(ctx context.Context, s *session, cmd string, args string) error {
	return cmdCommandAdd(ctx, s, args, levelModerator, true)
}

func cmdCommandAddSubscriber(ctx context.Context, s *session, cmd string, args string) error {
	return cmdCommandAdd(ctx, s, args, levelSubscriber, true)
}

func cmdCommandAddEveryone(ctx context.Context, s *session, cmd string, args string) error {
	return cmdCommandAdd(ctx, s, args, levelEveryone, true)
}

func cmdCommandAdd(ctx context.Context, s *session, args string, level accessLevel, forceLevel bool) error {
	usage := func() error {
		return s.ReplyUsage("<name> <text>")
	}

	name, text := splitSpace(args)
	name = cleanCommandName(name)

	if name == "" || text == "" {
		return usage()
	}

	if reservedCommandNames[name] {
		return s.Replyf("Command name '%s' is reserved.", name)
	}

	// TODO: remove this warning
	var warning string
	if _, ok := builtinCommands[name]; ok {
		warning = " Warning: '" + name + "' is a builtin command and will now only be accessible via " + s.Channel.Prefix + "builtin " + name
	}

	_, err := cbp.Parse(text)
	if err != nil {
		return s.Replyf("Error parsing command.%s", warning)
	}

	command, err := s.Channel.CustomCommands(
		models.CustomCommandWhere.Name.EQ(name),
		qm.For("UPDATE"),
	).One(ctx, s.Tx)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	update := err != sql.ErrNoRows

	if !s.UserLevel.CanAccess(level) {
		a := "add"
		if update {
			a = "update"
		}

		return s.Replyf("Your level is %s; you cannot %s a command with level %s.", s.UserLevel.PGEnum(), a, level.PGEnum())
	}

	if update {
		if !s.UserLevel.CanAccess(newAccessLevel(command.AccessLevel)) {
			al := flect.Pluralize(command.AccessLevel)
			return s.Replyf("Command '%s' is restricted to %s; only %s and above can update it.", name, al, al)
		}

		command.Message = text
		command.Editor = s.User

		if forceLevel {
			command.AccessLevel = level.PGEnum()
		}

		if err := command.Update(ctx, s.Tx, boil.Whitelist(models.CustomCommandColumns.UpdatedAt, models.CustomCommandColumns.Message, models.CustomCommandColumns.Editor)); err != nil {
			return err
		}

		al := flect.Pluralize(command.AccessLevel)
		return s.Replyf("Command '%s' updated, restricted to %s and above.%s", name, al, warning)
	}

	command = &models.CustomCommand{
		Name:        name,
		ChannelID:   s.Channel.ID,
		Message:     text,
		AccessLevel: level.PGEnum(),
		Creator:     s.User,
		Editor:      s.User,
	}

	if err := command.Insert(ctx, s.Tx, boil.Infer()); err != nil {
		return err
	}

	al := flect.Pluralize(command.AccessLevel)
	return s.Replyf("Command '%s' added, restricted to %s and above.%s", name, al, warning)
}

func cmdCommandDelete(ctx context.Context, s *session, cmd string, args string) error {
	usage := func() error {
		return s.ReplyUsage("<name>")
	}

	name, _ := splitSpace(args)
	name = cleanCommandName(name)

	if name == "" {
		return usage()
	}

	command, err := s.Channel.CustomCommands(
		models.CustomCommandWhere.Name.EQ(name),
		qm.For("UPDATE"),
		qm.Load(models.CustomCommandRels.RepeatedCommand),
		qm.Load(models.CustomCommandRels.ScheduledCommand),
	).One(ctx, s.Tx)

	if err == sql.ErrNoRows {
		return s.Replyf("Command '%s' does not exist.", name)
	}

	if err != nil {
		return err
	}

	level := newAccessLevel(command.AccessLevel)
	if !s.UserLevel.CanAccess(level) {
		return s.Replyf("Your level is %s; you cannot delete a command with level %s.", s.UserLevel.PGEnum(), command.AccessLevel)
	}

	deletedRepeat := false

	if command.R.RepeatedCommand != nil {
		deletedRepeat = true
		s.Deps.UpdateRepeat(command.R.RepeatedCommand.ID, false, 0, 0)

		if err := command.R.RepeatedCommand.Delete(ctx, s.Tx); err != nil {
			return err
		}
	}

	if command.R.ScheduledCommand != nil {
		deletedRepeat = true
		s.Deps.UpdateSchedule(command.R.ScheduledCommand.ID, false, nil)

		if err := command.R.ScheduledCommand.Delete(ctx, s.Tx); err != nil {
			return err
		}
	}

	if err := command.Delete(ctx, s.Tx); err != nil {
		return err
	}

	if deletedRepeat {
		return s.Replyf("Command '%s' and its repeat/schedule have been deleted.", name)
	}

	return s.Replyf("Command '%s' deleted.", name)
}

func cmdCommandRestrict(ctx context.Context, s *session, cmd string, args string) error {
	usage := func() error {
		return s.ReplyUsage("<name> everyone|regulars|subs|mods|broadcaster|admin")
	}

	name, level := splitSpace(args)
	name = cleanCommandName(name)

	if name == "" {
		return usage()
	}

	command, err := s.Channel.CustomCommands(
		models.CustomCommandWhere.Name.EQ(name),
		qm.For("UPDATE"),
	).One(ctx, s.Tx)

	if err == sql.ErrNoRows {
		return s.Replyf("Command '%s' does not exist.", name)
	}

	if err != nil {
		return err
	}

	if level == "" {
		return s.Replyf("Command '%s' is restricted to %s and above.", name, flect.Pluralize(command.AccessLevel))
	}

	level = strings.ToLower(level)

	var newLevel string
	switch level {
	case "everyone", "all", "everybody", "normal":
		newLevel = models.AccessLevelEveryone
	case "default", "sub", "subs", "subscriber", "subscrbers", "regular", "regulars", "reg", "regs":
		newLevel = models.AccessLevelSubscriber
	case "mod", "mods", "moderator", "moderators":
		newLevel = models.AccessLevelModerator
	case "broadcaster", "broadcasters", "owner", "owners", "streamer", "streamers":
		newLevel = models.AccessLevelBroadcaster
	case "admin", "admins":
		newLevel = models.AccessLevelAdmin
	default:
		return usage()
	}

	if !s.UserLevel.CanAccess(newAccessLevel(command.AccessLevel)) {
		return s.Replyf("Your level is %s; you cannot restrict a command with level %s.", s.UserLevel.PGEnum(), command.AccessLevel)
	}

	if !s.UserLevel.CanAccess(newAccessLevel(newLevel)) {
		return s.Replyf("Your level is %s; you cannot restrict a command to level %s.", s.UserLevel.PGEnum(), newLevel)
	}

	command.AccessLevel = newLevel
	command.Editor = s.User

	if err := command.Update(ctx, s.Tx, boil.Whitelist(models.CustomCommandColumns.UpdatedAt, models.CustomCommandColumns.AccessLevel, models.CustomCommandColumns.Editor)); err != nil {
		return err
	}

	return s.Replyf("Command '%s' restricted to %s and above.", name, flect.Pluralize(command.AccessLevel))
}

func cmdCommandProperty(ctx context.Context, s *session, prop string, args string) error {
	name, _ := splitSpace(args)
	name = cleanCommandName(name)

	if name == "" {
		return s.ReplyUsage("<name>")
	}

	command, err := s.Channel.CustomCommands(
		models.CustomCommandWhere.Name.EQ(name),
	).One(ctx, s.Tx)

	if err == sql.ErrNoRows {
		return s.Replyf("Command '%s' does not exist.", name)
	}

	if err != nil {
		return err
	}

	switch prop {
	case "editor", "author":
		return s.Replyf("Command '%s' was last modified by %s.", name, command.Editor) // TODO: include the date/time?
	case "count":
		u := "times"

		if command.Count == 1 {
			u = "time"
		}

		return s.Replyf("Command '%s' has been used %d %s.", name, command.Count, u)
	}

	panic("unreachable")
}

func cmdCommandRename(ctx context.Context, s *session, cmd string, args string) error {
	usage := func() error {
		return s.ReplyUsage("<old> <new>")
	}

	oldName, args := splitSpace(args)
	newName, _ := splitSpace(args)

	oldName = cleanCommandName(oldName)
	newName = cleanCommandName(newName)

	if oldName == "" || newName == "" {
		return usage()
	}

	if oldName == newName {
		return s.Replyf("'%s' is already called '%s'!", oldName, oldName)
	}

	command, err := s.Channel.CustomCommands(
		models.CustomCommandWhere.Name.EQ(oldName),
		qm.For("UPDATE"),
	).One(ctx, s.Tx)

	if err == sql.ErrNoRows {
		return s.Replyf("Command '%s' does not exist.", oldName)
	}

	if err != nil {
		return err
	}

	level := newAccessLevel(command.AccessLevel)
	if !s.UserLevel.CanAccess(level) {
		return s.Replyf("Your level is %s; you cannot rename a command with level %s.", s.UserLevel.PGEnum(), command.AccessLevel)
	}

	exists, err := s.Channel.CustomCommands(
		models.CustomCommandWhere.Name.EQ(newName),
	).Exists(ctx, s.Tx)

	if err != nil {
		return err
	}

	if exists {
		return s.Replyf("Command '%s' already exists.", newName)
	}

	command.Name = newName
	command.Editor = s.User

	if err := command.Update(ctx, s.Tx, boil.Whitelist(models.CustomCommandColumns.UpdatedAt, models.CustomCommandColumns.Name, models.CustomCommandColumns.Editor)); err != nil {
		return err
	}

	return s.Replyf("Command '%s' has been renamed to '%s'.", oldName, newName)
}

func cmdCommandGet(ctx context.Context, s *session, cmd string, args string) error {
	usage := func() error {
		return s.ReplyUsage("<name>")
	}

	name, _ := splitSpace(args)
	name = cleanCommandName(name)

	if name == "" {
		return usage()
	}

	command, err := s.Channel.CustomCommands(
		models.CustomCommandWhere.Name.EQ(name),
	).One(ctx, s.Tx)

	if err == sql.ErrNoRows {
		return s.Replyf("Command '%s' does not exist.", name)
	}

	if err != nil {
		return err
	}

	return s.Replyf("Command '%s': %s", name, command.Message)
}

func init() {
	flect.AddPlural("everyone", "everyone")
}