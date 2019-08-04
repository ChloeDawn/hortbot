package btest

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hortbot/hortbot/internal/db/models"
	"github.com/hortbot/hortbot/internal/db/modelsx"
	"github.com/volatiletech/sqlboiler/boil"
	"gotest.tools/assert"
)

func (st *scriptTester) insertChannel(t testing.TB, _, args string) {
	lineNum := st.lineNum

	var channel models.Channel
	assert.NilError(t, json.Unmarshal([]byte(args), &channel), "line %d", lineNum)

	st.addAction(func(ctx context.Context) {
		assert.NilError(t, channel.Insert(ctx, st.db, boil.Infer()), "line %d", lineNum)
	})
}

func (st *scriptTester) insertCustomCommand(t testing.TB, _, args string) {
	lineNum := st.lineNum

	var sc models.CustomCommand
	assert.NilError(t, json.Unmarshal([]byte(args), &sc), "line %d", lineNum)

	st.addAction(func(ctx context.Context) {
		assert.NilError(t, sc.Insert(ctx, st.db, boil.Infer()), "line %d", lineNum)
	})
}

func (st *scriptTester) insertRepeatedCommand(t testing.TB, _, args string) {
	lineNum := st.lineNum

	var rc models.RepeatedCommand
	assert.NilError(t, json.Unmarshal([]byte(args), &rc), "line %d", lineNum)

	st.addAction(func(ctx context.Context) {
		ctx = boil.SkipTimestamps(ctx)
		assert.NilError(t, rc.Insert(ctx, st.db, boil.Infer()), "line %d", lineNum)
	})
}

func (st *scriptTester) insertScheduledCommand(t testing.TB, _, args string) {
	lineNum := st.lineNum

	var sc models.ScheduledCommand
	assert.NilError(t, json.Unmarshal([]byte(args), &sc), "line %d", lineNum)

	st.addAction(func(ctx context.Context) {
		ctx = boil.SkipTimestamps(ctx)
		assert.NilError(t, sc.Insert(ctx, st.db, boil.Infer()), "line %d", lineNum)
	})
}

func (st *scriptTester) upsertTwitchToken(t testing.TB, _, args string) {
	lineNum := st.lineNum

	var tt models.TwitchToken
	assert.NilError(t, json.Unmarshal([]byte(args), &tt), "line %d", lineNum)

	st.addAction(func(ctx context.Context) {
		ctx = boil.SkipTimestamps(ctx)
		assert.NilError(t, modelsx.UpsertToken(ctx, st.db, &tt), "line %d", lineNum)
	})
}