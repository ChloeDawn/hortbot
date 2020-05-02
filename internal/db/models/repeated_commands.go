// Code generated by SQLBoiler v4.0.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// RepeatedCommand is an object representing the database table.
type RepeatedCommand struct {
	ID            int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt     time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt     time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	ChannelID     int64     `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	CommandInfoID int64     `boil:"command_info_id" json:"command_info_id" toml:"command_info_id" yaml:"command_info_id"`
	Enabled       bool      `boil:"enabled" json:"enabled" toml:"enabled" yaml:"enabled"`
	Delay         int       `boil:"delay" json:"delay" toml:"delay" yaml:"delay"`
	MessageDiff   int64     `boil:"message_diff" json:"message_diff" toml:"message_diff" yaml:"message_diff"`
	LastCount     int64     `boil:"last_count" json:"last_count" toml:"last_count" yaml:"last_count"`
	InitTimestamp null.Time `boil:"init_timestamp" json:"init_timestamp,omitempty" toml:"init_timestamp" yaml:"init_timestamp,omitempty"`
	Creator       string    `boil:"creator" json:"creator" toml:"creator" yaml:"creator"`
	Editor        string    `boil:"editor" json:"editor" toml:"editor" yaml:"editor"`

	R *repeatedCommandR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L repeatedCommandL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RepeatedCommandColumns = struct {
	ID            string
	CreatedAt     string
	UpdatedAt     string
	ChannelID     string
	CommandInfoID string
	Enabled       string
	Delay         string
	MessageDiff   string
	LastCount     string
	InitTimestamp string
	Creator       string
	Editor        string
}{
	ID:            "id",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	ChannelID:     "channel_id",
	CommandInfoID: "command_info_id",
	Enabled:       "enabled",
	Delay:         "delay",
	MessageDiff:   "message_diff",
	LastCount:     "last_count",
	InitTimestamp: "init_timestamp",
	Creator:       "creator",
	Editor:        "editor",
}

// Generated where

var RepeatedCommandWhere = struct {
	ID            whereHelperint64
	CreatedAt     whereHelpertime_Time
	UpdatedAt     whereHelpertime_Time
	ChannelID     whereHelperint64
	CommandInfoID whereHelperint64
	Enabled       whereHelperbool
	Delay         whereHelperint
	MessageDiff   whereHelperint64
	LastCount     whereHelperint64
	InitTimestamp whereHelpernull_Time
	Creator       whereHelperstring
	Editor        whereHelperstring
}{
	ID:            whereHelperint64{field: "\"repeated_commands\".\"id\""},
	CreatedAt:     whereHelpertime_Time{field: "\"repeated_commands\".\"created_at\""},
	UpdatedAt:     whereHelpertime_Time{field: "\"repeated_commands\".\"updated_at\""},
	ChannelID:     whereHelperint64{field: "\"repeated_commands\".\"channel_id\""},
	CommandInfoID: whereHelperint64{field: "\"repeated_commands\".\"command_info_id\""},
	Enabled:       whereHelperbool{field: "\"repeated_commands\".\"enabled\""},
	Delay:         whereHelperint{field: "\"repeated_commands\".\"delay\""},
	MessageDiff:   whereHelperint64{field: "\"repeated_commands\".\"message_diff\""},
	LastCount:     whereHelperint64{field: "\"repeated_commands\".\"last_count\""},
	InitTimestamp: whereHelpernull_Time{field: "\"repeated_commands\".\"init_timestamp\""},
	Creator:       whereHelperstring{field: "\"repeated_commands\".\"creator\""},
	Editor:        whereHelperstring{field: "\"repeated_commands\".\"editor\""},
}

// RepeatedCommandRels is where relationship names are stored.
var RepeatedCommandRels = struct {
	Channel     string
	CommandInfo string
}{
	Channel:     "Channel",
	CommandInfo: "CommandInfo",
}

// repeatedCommandR is where relationships are stored.
type repeatedCommandR struct {
	Channel     *Channel
	CommandInfo *CommandInfo
}

// NewStruct creates a new relationship struct
func (*repeatedCommandR) NewStruct() *repeatedCommandR {
	return &repeatedCommandR{}
}

// repeatedCommandL is where Load methods for each relationship are stored.
type repeatedCommandL struct{}

var (
	repeatedCommandAllColumns            = []string{"id", "created_at", "updated_at", "channel_id", "command_info_id", "enabled", "delay", "message_diff", "last_count", "init_timestamp", "creator", "editor"}
	repeatedCommandColumnsWithoutDefault = []string{"channel_id", "command_info_id", "enabled", "delay", "last_count", "init_timestamp", "creator", "editor"}
	repeatedCommandColumnsWithDefault    = []string{"id", "created_at", "updated_at", "message_diff"}
	repeatedCommandPrimaryKeyColumns     = []string{"id"}
)

type (
	// RepeatedCommandSlice is an alias for a slice of pointers to RepeatedCommand.
	// This should generally be used opposed to []RepeatedCommand.
	RepeatedCommandSlice []*RepeatedCommand

	repeatedCommandQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	repeatedCommandType                 = reflect.TypeOf(&RepeatedCommand{})
	repeatedCommandMapping              = queries.MakeStructMapping(repeatedCommandType)
	repeatedCommandPrimaryKeyMapping, _ = queries.BindMapping(repeatedCommandType, repeatedCommandMapping, repeatedCommandPrimaryKeyColumns)
	repeatedCommandInsertCacheMut       sync.RWMutex
	repeatedCommandInsertCache          = make(map[string]insertCache)
	repeatedCommandUpdateCacheMut       sync.RWMutex
	repeatedCommandUpdateCache          = make(map[string]updateCache)
	repeatedCommandUpsertCacheMut       sync.RWMutex
	repeatedCommandUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single repeatedCommand record from the query.
func (q repeatedCommandQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RepeatedCommand, error) {
	o := &RepeatedCommand{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for repeated_commands")
	}

	return o, nil
}

// All returns all RepeatedCommand records from the query.
func (q repeatedCommandQuery) All(ctx context.Context, exec boil.ContextExecutor) (RepeatedCommandSlice, error) {
	var o []*RepeatedCommand

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RepeatedCommand slice")
	}

	return o, nil
}

// Count returns the count of all RepeatedCommand records in the query.
func (q repeatedCommandQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count repeated_commands rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q repeatedCommandQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if repeated_commands exists")
	}

	return count > 0, nil
}

// Channel pointed to by the foreign key.
func (o *RepeatedCommand) Channel(mods ...qm.QueryMod) channelQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ChannelID),
	}

	queryMods = append(queryMods, mods...)

	query := Channels(queryMods...)
	queries.SetFrom(query.Query, "\"channels\"")

	return query
}

// CommandInfo pointed to by the foreign key.
func (o *RepeatedCommand) CommandInfo(mods ...qm.QueryMod) commandInfoQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.CommandInfoID),
	}

	queryMods = append(queryMods, mods...)

	query := CommandInfos(queryMods...)
	queries.SetFrom(query.Query, "\"command_infos\"")

	return query
}

// LoadChannel allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (repeatedCommandL) LoadChannel(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRepeatedCommand interface{}, mods queries.Applicator) error {
	var slice []*RepeatedCommand
	var object *RepeatedCommand

	if singular {
		object = maybeRepeatedCommand.(*RepeatedCommand)
	} else {
		slice = *maybeRepeatedCommand.(*[]*RepeatedCommand)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &repeatedCommandR{}
		}
		args = append(args, object.ChannelID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &repeatedCommandR{}
			}

			for _, a := range args {
				if a == obj.ChannelID {
					continue Outer
				}
			}

			args = append(args, obj.ChannelID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`channels`),
		qm.WhereIn(`channels.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Channel")
	}

	var resultSlice []*Channel
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Channel")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for channels")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for channels")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Channel = foreign
		if foreign.R == nil {
			foreign.R = &channelR{}
		}
		foreign.R.RepeatedCommands = append(foreign.R.RepeatedCommands, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ChannelID == foreign.ID {
				local.R.Channel = foreign
				if foreign.R == nil {
					foreign.R = &channelR{}
				}
				foreign.R.RepeatedCommands = append(foreign.R.RepeatedCommands, local)
				break
			}
		}
	}

	return nil
}

// LoadCommandInfo allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (repeatedCommandL) LoadCommandInfo(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRepeatedCommand interface{}, mods queries.Applicator) error {
	var slice []*RepeatedCommand
	var object *RepeatedCommand

	if singular {
		object = maybeRepeatedCommand.(*RepeatedCommand)
	} else {
		slice = *maybeRepeatedCommand.(*[]*RepeatedCommand)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &repeatedCommandR{}
		}
		args = append(args, object.CommandInfoID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &repeatedCommandR{}
			}

			for _, a := range args {
				if a == obj.CommandInfoID {
					continue Outer
				}
			}

			args = append(args, obj.CommandInfoID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`command_infos`),
		qm.WhereIn(`command_infos.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load CommandInfo")
	}

	var resultSlice []*CommandInfo
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice CommandInfo")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for command_infos")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for command_infos")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.CommandInfo = foreign
		if foreign.R == nil {
			foreign.R = &commandInfoR{}
		}
		foreign.R.RepeatedCommand = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CommandInfoID == foreign.ID {
				local.R.CommandInfo = foreign
				if foreign.R == nil {
					foreign.R = &commandInfoR{}
				}
				foreign.R.RepeatedCommand = local
				break
			}
		}
	}

	return nil
}

// SetChannel of the repeatedCommand to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.RepeatedCommands.
func (o *RepeatedCommand) SetChannel(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Channel) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"repeated_commands\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"channel_id"}),
		strmangle.WhereClause("\"", "\"", 2, repeatedCommandPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ChannelID = related.ID
	if o.R == nil {
		o.R = &repeatedCommandR{
			Channel: related,
		}
	} else {
		o.R.Channel = related
	}

	if related.R == nil {
		related.R = &channelR{
			RepeatedCommands: RepeatedCommandSlice{o},
		}
	} else {
		related.R.RepeatedCommands = append(related.R.RepeatedCommands, o)
	}

	return nil
}

// SetCommandInfo of the repeatedCommand to the related item.
// Sets o.R.CommandInfo to related.
// Adds o to related.R.RepeatedCommand.
func (o *RepeatedCommand) SetCommandInfo(ctx context.Context, exec boil.ContextExecutor, insert bool, related *CommandInfo) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"repeated_commands\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"command_info_id"}),
		strmangle.WhereClause("\"", "\"", 2, repeatedCommandPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CommandInfoID = related.ID
	if o.R == nil {
		o.R = &repeatedCommandR{
			CommandInfo: related,
		}
	} else {
		o.R.CommandInfo = related
	}

	if related.R == nil {
		related.R = &commandInfoR{
			RepeatedCommand: o,
		}
	} else {
		related.R.RepeatedCommand = o
	}

	return nil
}

// RepeatedCommands retrieves all the records using an executor.
func RepeatedCommands(mods ...qm.QueryMod) repeatedCommandQuery {
	mods = append(mods, qm.From("\"repeated_commands\""))
	return repeatedCommandQuery{NewQuery(mods...)}
}

// FindRepeatedCommand retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRepeatedCommand(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*RepeatedCommand, error) {
	repeatedCommandObj := &RepeatedCommand{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"repeated_commands\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, repeatedCommandObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from repeated_commands")
	}

	return repeatedCommandObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RepeatedCommand) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no repeated_commands provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(repeatedCommandColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	repeatedCommandInsertCacheMut.RLock()
	cache, cached := repeatedCommandInsertCache[key]
	repeatedCommandInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			repeatedCommandAllColumns,
			repeatedCommandColumnsWithDefault,
			repeatedCommandColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(repeatedCommandType, repeatedCommandMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(repeatedCommandType, repeatedCommandMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"repeated_commands\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"repeated_commands\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into repeated_commands")
	}

	if !cached {
		repeatedCommandInsertCacheMut.Lock()
		repeatedCommandInsertCache[key] = cache
		repeatedCommandInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the RepeatedCommand.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RepeatedCommand) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	repeatedCommandUpdateCacheMut.RLock()
	cache, cached := repeatedCommandUpdateCache[key]
	repeatedCommandUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			repeatedCommandAllColumns,
			repeatedCommandPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update repeated_commands, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"repeated_commands\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, repeatedCommandPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(repeatedCommandType, repeatedCommandMapping, append(wl, repeatedCommandPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	_, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update repeated_commands row")
	}

	if !cached {
		repeatedCommandUpdateCacheMut.Lock()
		repeatedCommandUpdateCache[key] = cache
		repeatedCommandUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q repeatedCommandQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for repeated_commands")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RepeatedCommandSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), repeatedCommandPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"repeated_commands\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, repeatedCommandPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in repeatedCommand slice")
	}

	return nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RepeatedCommand) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no repeated_commands provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(repeatedCommandColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	repeatedCommandUpsertCacheMut.RLock()
	cache, cached := repeatedCommandUpsertCache[key]
	repeatedCommandUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			repeatedCommandAllColumns,
			repeatedCommandColumnsWithDefault,
			repeatedCommandColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			repeatedCommandAllColumns,
			repeatedCommandPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert repeated_commands, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(repeatedCommandPrimaryKeyColumns))
			copy(conflict, repeatedCommandPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"repeated_commands\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(repeatedCommandType, repeatedCommandMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(repeatedCommandType, repeatedCommandMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert repeated_commands")
	}

	if !cached {
		repeatedCommandUpsertCacheMut.Lock()
		repeatedCommandUpsertCache[key] = cache
		repeatedCommandUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single RepeatedCommand record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RepeatedCommand) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no RepeatedCommand provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), repeatedCommandPrimaryKeyMapping)
	sql := "DELETE FROM \"repeated_commands\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from repeated_commands")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q repeatedCommandQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no repeatedCommandQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from repeated_commands")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RepeatedCommandSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), repeatedCommandPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"repeated_commands\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, repeatedCommandPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from repeatedCommand slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *RepeatedCommand) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRepeatedCommand(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RepeatedCommandSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RepeatedCommandSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), repeatedCommandPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"repeated_commands\".* FROM \"repeated_commands\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, repeatedCommandPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RepeatedCommandSlice")
	}

	*o = slice

	return nil
}

// RepeatedCommandExists checks if the RepeatedCommand row exists.
func RepeatedCommandExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"repeated_commands\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if repeated_commands exists")
	}

	return exists, nil
}
