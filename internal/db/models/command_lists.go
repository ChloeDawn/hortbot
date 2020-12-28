// Code generated by SQLBoiler v4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// CommandList is an object representing the database table.
type CommandList struct {
	ID        int64             `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time         `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time         `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	ChannelID int64             `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	Items     types.StringArray `boil:"items" json:"items" toml:"items" yaml:"items"`

	R *commandListR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L commandListL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CommandListColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	ChannelID string
	Items     string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	ChannelID: "channel_id",
	Items:     "items",
}

// Generated where

var CommandListWhere = struct {
	ID        whereHelperint64
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	ChannelID whereHelperint64
	Items     whereHelpertypes_StringArray
}{
	ID:        whereHelperint64{field: "\"command_lists\".\"id\""},
	CreatedAt: whereHelpertime_Time{field: "\"command_lists\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"command_lists\".\"updated_at\""},
	ChannelID: whereHelperint64{field: "\"command_lists\".\"channel_id\""},
	Items:     whereHelpertypes_StringArray{field: "\"command_lists\".\"items\""},
}

// CommandListRels is where relationship names are stored.
var CommandListRels = struct {
	Channel     string
	CommandInfo string
}{
	Channel:     "Channel",
	CommandInfo: "CommandInfo",
}

// commandListR is where relationships are stored.
type commandListR struct {
	Channel     *Channel     `boil:"Channel" json:"Channel" toml:"Channel" yaml:"Channel"`
	CommandInfo *CommandInfo `boil:"CommandInfo" json:"CommandInfo" toml:"CommandInfo" yaml:"CommandInfo"`
}

// NewStruct creates a new relationship struct
func (*commandListR) NewStruct() *commandListR {
	return &commandListR{}
}

// commandListL is where Load methods for each relationship are stored.
type commandListL struct{}

var (
	commandListAllColumns            = []string{"id", "created_at", "updated_at", "channel_id", "items"}
	commandListColumnsWithoutDefault = []string{"channel_id"}
	commandListColumnsWithDefault    = []string{"id", "created_at", "updated_at", "items"}
	commandListPrimaryKeyColumns     = []string{"id"}
)

type (
	// CommandListSlice is an alias for a slice of pointers to CommandList.
	// This should generally be used opposed to []CommandList.
	CommandListSlice []*CommandList

	commandListQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	commandListType                 = reflect.TypeOf(&CommandList{})
	commandListMapping              = queries.MakeStructMapping(commandListType)
	commandListPrimaryKeyMapping, _ = queries.BindMapping(commandListType, commandListMapping, commandListPrimaryKeyColumns)
	commandListInsertCacheMut       sync.RWMutex
	commandListInsertCache          = make(map[string]insertCache)
	commandListUpdateCacheMut       sync.RWMutex
	commandListUpdateCache          = make(map[string]updateCache)
	commandListUpsertCacheMut       sync.RWMutex
	commandListUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single commandList record from the query.
func (q commandListQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CommandList, error) {
	o := &CommandList{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for command_lists")
	}

	return o, nil
}

// All returns all CommandList records from the query.
func (q commandListQuery) All(ctx context.Context, exec boil.ContextExecutor) (CommandListSlice, error) {
	var o []*CommandList

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CommandList slice")
	}

	return o, nil
}

// Count returns the count of all CommandList records in the query.
func (q commandListQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count command_lists rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q commandListQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if command_lists exists")
	}

	return count > 0, nil
}

// Channel pointed to by the foreign key.
func (o *CommandList) Channel(mods ...qm.QueryMod) channelQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ChannelID),
	}

	queryMods = append(queryMods, mods...)

	query := Channels(queryMods...)
	queries.SetFrom(query.Query, "\"channels\"")

	return query
}

// CommandInfo pointed to by the foreign key.
func (o *CommandList) CommandInfo(mods ...qm.QueryMod) commandInfoQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"command_list_id\" = ?", o.ID),
	}

	queryMods = append(queryMods, mods...)

	query := CommandInfos(queryMods...)
	queries.SetFrom(query.Query, "\"command_infos\"")

	return query
}

// LoadChannel allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (commandListL) LoadChannel(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCommandList interface{}, mods queries.Applicator) error {
	var slice []*CommandList
	var object *CommandList

	if singular {
		object = maybeCommandList.(*CommandList)
	} else {
		slice = *maybeCommandList.(*[]*CommandList)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &commandListR{}
		}
		args = append(args, object.ChannelID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &commandListR{}
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
		foreign.R.CommandLists = append(foreign.R.CommandLists, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ChannelID == foreign.ID {
				local.R.Channel = foreign
				if foreign.R == nil {
					foreign.R = &channelR{}
				}
				foreign.R.CommandLists = append(foreign.R.CommandLists, local)
				break
			}
		}
	}

	return nil
}

// LoadCommandInfo allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-1 relationship.
func (commandListL) LoadCommandInfo(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCommandList interface{}, mods queries.Applicator) error {
	var slice []*CommandList
	var object *CommandList

	if singular {
		object = maybeCommandList.(*CommandList)
	} else {
		slice = *maybeCommandList.(*[]*CommandList)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &commandListR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &commandListR{}
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
		qm.WhereIn(`command_infos.command_list_id in ?`, args...),
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
		foreign.R.CommandList = object
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ID, foreign.CommandListID) {
				local.R.CommandInfo = foreign
				if foreign.R == nil {
					foreign.R = &commandInfoR{}
				}
				foreign.R.CommandList = local
				break
			}
		}
	}

	return nil
}

// SetChannel of the commandList to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.CommandLists.
func (o *CommandList) SetChannel(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Channel) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"command_lists\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"channel_id"}),
		strmangle.WhereClause("\"", "\"", 2, commandListPrimaryKeyColumns),
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
		o.R = &commandListR{
			Channel: related,
		}
	} else {
		o.R.Channel = related
	}

	if related.R == nil {
		related.R = &channelR{
			CommandLists: CommandListSlice{o},
		}
	} else {
		related.R.CommandLists = append(related.R.CommandLists, o)
	}

	return nil
}

// SetCommandInfo of the commandList to the related item.
// Sets o.R.CommandInfo to related.
// Adds o to related.R.CommandList.
func (o *CommandList) SetCommandInfo(ctx context.Context, exec boil.ContextExecutor, insert bool, related *CommandInfo) error {
	var err error

	if insert {
		queries.Assign(&related.CommandListID, o.ID)

		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"command_infos\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"command_list_id"}),
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

		queries.Assign(&related.CommandListID, o.ID)
	}

	if o.R == nil {
		o.R = &commandListR{
			CommandInfo: related,
		}
	} else {
		o.R.CommandInfo = related
	}

	if related.R == nil {
		related.R = &commandInfoR{
			CommandList: o,
		}
	} else {
		related.R.CommandList = o
	}
	return nil
}

// RemoveCommandInfo relationship.
// Sets o.R.CommandInfo to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *CommandList) RemoveCommandInfo(ctx context.Context, exec boil.ContextExecutor, related *CommandInfo) error {
	var err error

	queries.SetScanner(&related.CommandListID, nil)
	if err = related.Update(ctx, exec, boil.Whitelist("command_list_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.CommandInfo = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	related.R.CommandList = nil
	return nil
}

// CommandLists retrieves all the records using an executor.
func CommandLists(mods ...qm.QueryMod) commandListQuery {
	mods = append(mods, qm.From("\"command_lists\""))
	return commandListQuery{NewQuery(mods...)}
}

// FindCommandList retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCommandList(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*CommandList, error) {
	commandListObj := &CommandList{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"command_lists\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, commandListObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from command_lists")
	}

	return commandListObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CommandList) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no command_lists provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(commandListColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	commandListInsertCacheMut.RLock()
	cache, cached := commandListInsertCache[key]
	commandListInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			commandListAllColumns,
			commandListColumnsWithDefault,
			commandListColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(commandListType, commandListMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(commandListType, commandListMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"command_lists\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"command_lists\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into command_lists")
	}

	if !cached {
		commandListInsertCacheMut.Lock()
		commandListInsertCache[key] = cache
		commandListInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the CommandList.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CommandList) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	commandListUpdateCacheMut.RLock()
	cache, cached := commandListUpdateCache[key]
	commandListUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			commandListAllColumns,
			commandListPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update command_lists, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"command_lists\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, commandListPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(commandListType, commandListMapping, append(wl, commandListPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update command_lists row")
	}

	if !cached {
		commandListUpdateCacheMut.Lock()
		commandListUpdateCache[key] = cache
		commandListUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q commandListQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for command_lists")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CommandListSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commandListPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"command_lists\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, commandListPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in commandList slice")
	}

	return nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CommandList) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no command_lists provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(commandListColumnsWithDefault, o)

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

	commandListUpsertCacheMut.RLock()
	cache, cached := commandListUpsertCache[key]
	commandListUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			commandListAllColumns,
			commandListColumnsWithDefault,
			commandListColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			commandListAllColumns,
			commandListPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert command_lists, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(commandListPrimaryKeyColumns))
			copy(conflict, commandListPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"command_lists\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(commandListType, commandListMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(commandListType, commandListMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert command_lists")
	}

	if !cached {
		commandListUpsertCacheMut.Lock()
		commandListUpsertCache[key] = cache
		commandListUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single CommandList record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CommandList) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no CommandList provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), commandListPrimaryKeyMapping)
	sql := "DELETE FROM \"command_lists\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from command_lists")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q commandListQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no commandListQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from command_lists")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CommandListSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commandListPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"command_lists\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, commandListPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from commandList slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CommandList) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCommandList(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CommandListSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CommandListSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commandListPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"command_lists\".* FROM \"command_lists\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, commandListPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CommandListSlice")
	}

	*o = slice

	return nil
}

// CommandListExists checks if the CommandList row exists.
func CommandListExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"command_lists\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if command_lists exists")
	}

	return exists, nil
}
