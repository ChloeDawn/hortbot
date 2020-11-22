// Code generated by SQLBoiler v4.3.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// CustomCommand is an object representing the database table.
type CustomCommand struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	ChannelID int64     `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	Message   string    `boil:"message" json:"message" toml:"message" yaml:"message"`

	R *customCommandR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L customCommandL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CustomCommandColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	ChannelID string
	Message   string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	ChannelID: "channel_id",
	Message:   "message",
}

// Generated where

var CustomCommandWhere = struct {
	ID        whereHelperint64
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	ChannelID whereHelperint64
	Message   whereHelperstring
}{
	ID:        whereHelperint64{field: "\"custom_commands\".\"id\""},
	CreatedAt: whereHelpertime_Time{field: "\"custom_commands\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"custom_commands\".\"updated_at\""},
	ChannelID: whereHelperint64{field: "\"custom_commands\".\"channel_id\""},
	Message:   whereHelperstring{field: "\"custom_commands\".\"message\""},
}

// CustomCommandRels is where relationship names are stored.
var CustomCommandRels = struct {
	Channel     string
	CommandInfo string
}{
	Channel:     "Channel",
	CommandInfo: "CommandInfo",
}

// customCommandR is where relationships are stored.
type customCommandR struct {
	Channel     *Channel     `boil:"Channel" json:"Channel" toml:"Channel" yaml:"Channel"`
	CommandInfo *CommandInfo `boil:"CommandInfo" json:"CommandInfo" toml:"CommandInfo" yaml:"CommandInfo"`
}

// NewStruct creates a new relationship struct
func (*customCommandR) NewStruct() *customCommandR {
	return &customCommandR{}
}

// customCommandL is where Load methods for each relationship are stored.
type customCommandL struct{}

var (
	customCommandAllColumns            = []string{"id", "created_at", "updated_at", "channel_id", "message"}
	customCommandColumnsWithoutDefault = []string{"channel_id", "message"}
	customCommandColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	customCommandPrimaryKeyColumns     = []string{"id"}
)

type (
	// CustomCommandSlice is an alias for a slice of pointers to CustomCommand.
	// This should generally be used opposed to []CustomCommand.
	CustomCommandSlice []*CustomCommand

	customCommandQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	customCommandType                 = reflect.TypeOf(&CustomCommand{})
	customCommandMapping              = queries.MakeStructMapping(customCommandType)
	customCommandPrimaryKeyMapping, _ = queries.BindMapping(customCommandType, customCommandMapping, customCommandPrimaryKeyColumns)
	customCommandInsertCacheMut       sync.RWMutex
	customCommandInsertCache          = make(map[string]insertCache)
	customCommandUpdateCacheMut       sync.RWMutex
	customCommandUpdateCache          = make(map[string]updateCache)
	customCommandUpsertCacheMut       sync.RWMutex
	customCommandUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single customCommand record from the query.
func (q customCommandQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CustomCommand, error) {
	o := &CustomCommand{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for custom_commands")
	}

	return o, nil
}

// All returns all CustomCommand records from the query.
func (q customCommandQuery) All(ctx context.Context, exec boil.ContextExecutor) (CustomCommandSlice, error) {
	var o []*CustomCommand

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CustomCommand slice")
	}

	return o, nil
}

// Count returns the count of all CustomCommand records in the query.
func (q customCommandQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count custom_commands rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q customCommandQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if custom_commands exists")
	}

	return count > 0, nil
}

// Channel pointed to by the foreign key.
func (o *CustomCommand) Channel(mods ...qm.QueryMod) channelQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ChannelID),
	}

	queryMods = append(queryMods, mods...)

	query := Channels(queryMods...)
	queries.SetFrom(query.Query, "\"channels\"")

	return query
}

// CommandInfo pointed to by the foreign key.
func (o *CustomCommand) CommandInfo(mods ...qm.QueryMod) commandInfoQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"custom_command_id\" = ?", o.ID),
	}

	queryMods = append(queryMods, mods...)

	query := CommandInfos(queryMods...)
	queries.SetFrom(query.Query, "\"command_infos\"")

	return query
}

// LoadChannel allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (customCommandL) LoadChannel(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCustomCommand interface{}, mods queries.Applicator) error {
	var slice []*CustomCommand
	var object *CustomCommand

	if singular {
		object = maybeCustomCommand.(*CustomCommand)
	} else {
		slice = *maybeCustomCommand.(*[]*CustomCommand)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &customCommandR{}
		}
		args = append(args, object.ChannelID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &customCommandR{}
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
		foreign.R.CustomCommands = append(foreign.R.CustomCommands, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ChannelID == foreign.ID {
				local.R.Channel = foreign
				if foreign.R == nil {
					foreign.R = &channelR{}
				}
				foreign.R.CustomCommands = append(foreign.R.CustomCommands, local)
				break
			}
		}
	}

	return nil
}

// LoadCommandInfo allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-1 relationship.
func (customCommandL) LoadCommandInfo(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCustomCommand interface{}, mods queries.Applicator) error {
	var slice []*CustomCommand
	var object *CustomCommand

	if singular {
		object = maybeCustomCommand.(*CustomCommand)
	} else {
		slice = *maybeCustomCommand.(*[]*CustomCommand)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &customCommandR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &customCommandR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`command_infos`),
		qm.WhereIn(`command_infos.custom_command_id in ?`, args...),
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
		foreign.R.CustomCommand = object
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ID, foreign.CustomCommandID) {
				local.R.CommandInfo = foreign
				if foreign.R == nil {
					foreign.R = &commandInfoR{}
				}
				foreign.R.CustomCommand = local
				break
			}
		}
	}

	return nil
}

// SetChannel of the customCommand to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.CustomCommands.
func (o *CustomCommand) SetChannel(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Channel) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"custom_commands\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"channel_id"}),
		strmangle.WhereClause("\"", "\"", 2, customCommandPrimaryKeyColumns),
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
		o.R = &customCommandR{
			Channel: related,
		}
	} else {
		o.R.Channel = related
	}

	if related.R == nil {
		related.R = &channelR{
			CustomCommands: CustomCommandSlice{o},
		}
	} else {
		related.R.CustomCommands = append(related.R.CustomCommands, o)
	}

	return nil
}

// SetCommandInfo of the customCommand to the related item.
// Sets o.R.CommandInfo to related.
// Adds o to related.R.CustomCommand.
func (o *CustomCommand) SetCommandInfo(ctx context.Context, exec boil.ContextExecutor, insert bool, related *CommandInfo) error {
	var err error

	if insert {
		queries.Assign(&related.CustomCommandID, o.ID)

		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"command_infos\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"custom_command_id"}),
			strmangle.WhereClause("\"", "\"", 2, commandInfoPrimaryKeyColumns),
		)
		values := []interface{}{o.ID, related.ID}

		if boil.IsDebug(ctx) {
			writer := boil.DebugWriterFrom(ctx)
			fmt.Fprintln(writer, updateQuery)
			fmt.Fprintln(writer, values)
		}
		if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		queries.Assign(&related.CustomCommandID, o.ID)
	}

	if o.R == nil {
		o.R = &customCommandR{
			CommandInfo: related,
		}
	} else {
		o.R.CommandInfo = related
	}

	if related.R == nil {
		related.R = &commandInfoR{
			CustomCommand: o,
		}
	} else {
		related.R.CustomCommand = o
	}
	return nil
}

// RemoveCommandInfo relationship.
// Sets o.R.CommandInfo to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *CustomCommand) RemoveCommandInfo(ctx context.Context, exec boil.ContextExecutor, related *CommandInfo) error {
	var err error

	queries.SetScanner(&related.CustomCommandID, nil)
	if err = related.Update(ctx, exec, boil.Whitelist("custom_command_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.CommandInfo = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	related.R.CustomCommand = nil
	return nil
}

// CustomCommands retrieves all the records using an executor.
func CustomCommands(mods ...qm.QueryMod) customCommandQuery {
	mods = append(mods, qm.From("\"custom_commands\""))
	return customCommandQuery{NewQuery(mods...)}
}

// FindCustomCommand retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCustomCommand(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*CustomCommand, error) {
	customCommandObj := &CustomCommand{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"custom_commands\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, customCommandObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from custom_commands")
	}

	return customCommandObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CustomCommand) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no custom_commands provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(customCommandColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	customCommandInsertCacheMut.RLock()
	cache, cached := customCommandInsertCache[key]
	customCommandInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			customCommandAllColumns,
			customCommandColumnsWithDefault,
			customCommandColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(customCommandType, customCommandMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(customCommandType, customCommandMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"custom_commands\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"custom_commands\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into custom_commands")
	}

	if !cached {
		customCommandInsertCacheMut.Lock()
		customCommandInsertCache[key] = cache
		customCommandInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the CustomCommand.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CustomCommand) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	customCommandUpdateCacheMut.RLock()
	cache, cached := customCommandUpdateCache[key]
	customCommandUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			customCommandAllColumns,
			customCommandPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update custom_commands, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"custom_commands\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, customCommandPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(customCommandType, customCommandMapping, append(wl, customCommandPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update custom_commands row")
	}

	if !cached {
		customCommandUpdateCacheMut.Lock()
		customCommandUpdateCache[key] = cache
		customCommandUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q customCommandQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for custom_commands")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CustomCommandSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), customCommandPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"custom_commands\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, customCommandPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in customCommand slice")
	}

	return nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CustomCommand) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no custom_commands provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(customCommandColumnsWithDefault, o)

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

	customCommandUpsertCacheMut.RLock()
	cache, cached := customCommandUpsertCache[key]
	customCommandUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			customCommandAllColumns,
			customCommandColumnsWithDefault,
			customCommandColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			customCommandAllColumns,
			customCommandPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert custom_commands, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(customCommandPrimaryKeyColumns))
			copy(conflict, customCommandPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"custom_commands\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(customCommandType, customCommandMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(customCommandType, customCommandMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert custom_commands")
	}

	if !cached {
		customCommandUpsertCacheMut.Lock()
		customCommandUpsertCache[key] = cache
		customCommandUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single CustomCommand record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CustomCommand) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no CustomCommand provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), customCommandPrimaryKeyMapping)
	sql := "DELETE FROM \"custom_commands\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from custom_commands")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q customCommandQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no customCommandQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from custom_commands")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CustomCommandSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), customCommandPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"custom_commands\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, customCommandPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from customCommand slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CustomCommand) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCustomCommand(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CustomCommandSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CustomCommandSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), customCommandPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"custom_commands\".* FROM \"custom_commands\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, customCommandPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CustomCommandSlice")
	}

	*o = slice

	return nil
}

// CustomCommandExists checks if the CustomCommand row exists.
func CustomCommandExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"custom_commands\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if custom_commands exists")
	}

	return exists, nil
}
