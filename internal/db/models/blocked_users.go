// Code generated by SQLBoiler v4.3.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// BlockedUser is an object representing the database table.
type BlockedUser struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	TwitchID  int64     `boil:"twitch_id" json:"twitch_id" toml:"twitch_id" yaml:"twitch_id"`

	R *blockedUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L blockedUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BlockedUserColumns = struct {
	ID        string
	CreatedAt string
	TwitchID  string
}{
	ID:        "id",
	CreatedAt: "created_at",
	TwitchID:  "twitch_id",
}

// Generated where

var BlockedUserWhere = struct {
	ID        whereHelperint64
	CreatedAt whereHelpertime_Time
	TwitchID  whereHelperint64
}{
	ID:        whereHelperint64{field: "\"blocked_users\".\"id\""},
	CreatedAt: whereHelpertime_Time{field: "\"blocked_users\".\"created_at\""},
	TwitchID:  whereHelperint64{field: "\"blocked_users\".\"twitch_id\""},
}

// BlockedUserRels is where relationship names are stored.
var BlockedUserRels = struct {
}{}

// blockedUserR is where relationships are stored.
type blockedUserR struct {
}

// NewStruct creates a new relationship struct
func (*blockedUserR) NewStruct() *blockedUserR {
	return &blockedUserR{}
}

// blockedUserL is where Load methods for each relationship are stored.
type blockedUserL struct{}

var (
	blockedUserAllColumns            = []string{"id", "created_at", "twitch_id"}
	blockedUserColumnsWithoutDefault = []string{"twitch_id"}
	blockedUserColumnsWithDefault    = []string{"id", "created_at"}
	blockedUserPrimaryKeyColumns     = []string{"id"}
)

type (
	// BlockedUserSlice is an alias for a slice of pointers to BlockedUser.
	// This should generally be used opposed to []BlockedUser.
	BlockedUserSlice []*BlockedUser

	blockedUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	blockedUserType                 = reflect.TypeOf(&BlockedUser{})
	blockedUserMapping              = queries.MakeStructMapping(blockedUserType)
	blockedUserPrimaryKeyMapping, _ = queries.BindMapping(blockedUserType, blockedUserMapping, blockedUserPrimaryKeyColumns)
	blockedUserInsertCacheMut       sync.RWMutex
	blockedUserInsertCache          = make(map[string]insertCache)
	blockedUserUpdateCacheMut       sync.RWMutex
	blockedUserUpdateCache          = make(map[string]updateCache)
	blockedUserUpsertCacheMut       sync.RWMutex
	blockedUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single blockedUser record from the query.
func (q blockedUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BlockedUser, error) {
	o := &BlockedUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for blocked_users")
	}

	return o, nil
}

// All returns all BlockedUser records from the query.
func (q blockedUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (BlockedUserSlice, error) {
	var o []*BlockedUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BlockedUser slice")
	}

	return o, nil
}

// Count returns the count of all BlockedUser records in the query.
func (q blockedUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count blocked_users rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q blockedUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if blocked_users exists")
	}

	return count > 0, nil
}

// BlockedUsers retrieves all the records using an executor.
func BlockedUsers(mods ...qm.QueryMod) blockedUserQuery {
	mods = append(mods, qm.From("\"blocked_users\""))
	return blockedUserQuery{NewQuery(mods...)}
}

// FindBlockedUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBlockedUser(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*BlockedUser, error) {
	blockedUserObj := &BlockedUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"blocked_users\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, blockedUserObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from blocked_users")
	}

	return blockedUserObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BlockedUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no blocked_users provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(blockedUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	blockedUserInsertCacheMut.RLock()
	cache, cached := blockedUserInsertCache[key]
	blockedUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			blockedUserAllColumns,
			blockedUserColumnsWithDefault,
			blockedUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(blockedUserType, blockedUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(blockedUserType, blockedUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"blocked_users\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"blocked_users\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into blocked_users")
	}

	if !cached {
		blockedUserInsertCacheMut.Lock()
		blockedUserInsertCache[key] = cache
		blockedUserInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the BlockedUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BlockedUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	blockedUserUpdateCacheMut.RLock()
	cache, cached := blockedUserUpdateCache[key]
	blockedUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			blockedUserAllColumns,
			blockedUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update blocked_users, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"blocked_users\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, blockedUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(blockedUserType, blockedUserMapping, append(wl, blockedUserPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update blocked_users row")
	}

	if !cached {
		blockedUserUpdateCacheMut.Lock()
		blockedUserUpdateCache[key] = cache
		blockedUserUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q blockedUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for blocked_users")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BlockedUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), blockedUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"blocked_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, blockedUserPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in blockedUser slice")
	}

	return nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BlockedUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no blocked_users provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(blockedUserColumnsWithDefault, o)

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

	blockedUserUpsertCacheMut.RLock()
	cache, cached := blockedUserUpsertCache[key]
	blockedUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			blockedUserAllColumns,
			blockedUserColumnsWithDefault,
			blockedUserColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			blockedUserAllColumns,
			blockedUserPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert blocked_users, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(blockedUserPrimaryKeyColumns))
			copy(conflict, blockedUserPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"blocked_users\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(blockedUserType, blockedUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(blockedUserType, blockedUserMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert blocked_users")
	}

	if !cached {
		blockedUserUpsertCacheMut.Lock()
		blockedUserUpsertCache[key] = cache
		blockedUserUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single BlockedUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BlockedUser) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no BlockedUser provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), blockedUserPrimaryKeyMapping)
	sql := "DELETE FROM \"blocked_users\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from blocked_users")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q blockedUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no blockedUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from blocked_users")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BlockedUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), blockedUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"blocked_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, blockedUserPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from blockedUser slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BlockedUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBlockedUser(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BlockedUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BlockedUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), blockedUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"blocked_users\".* FROM \"blocked_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, blockedUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BlockedUserSlice")
	}

	*o = slice

	return nil
}

// BlockedUserExists checks if the BlockedUser row exists.
func BlockedUserExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"blocked_users\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if blocked_users exists")
	}

	return exists, nil
}
