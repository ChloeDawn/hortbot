package fakeirc

import (
	"context"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/hortbot/hortbot/internal/x/errgroupx"
	"github.com/jakebailey/irc"
	"gotest.tools/assert"
)

const DefaultSleepDur = 50 * time.Millisecond

func getSleepDur() (time.Duration, error) {
	s, ok := os.LookupEnv("TEST_HELPER_SLEEP_DUR")
	if ok {
		return time.ParseDuration(s)
	}

	return DefaultSleepDur, nil
}

type Helper struct {
	SleepDur time.Duration

	stopOnce sync.Once

	t *testing.T
	g *errgroupx.Group
	s *Server
}

func NewHelper(ctx context.Context, t *testing.T, opts ...Option) *Helper {
	t.Helper()

	dur, err := getSleepDur()
	assert.NilError(t, err)

	server, err := Start(opts...)
	assert.NilError(t, err)
	assert.Assert(t, server != nil)

	return &Helper{
		SleepDur: dur,
		t:        t,
		g:        errgroupx.FromContext(ctx),
		s:        server,
	}
}

func (h *Helper) CollectFromChannel(ch <-chan *irc.Message) *[]*irc.Message {
	h.t.Helper()
	messages := []*irc.Message{}

	h.g.Go(func(ctx context.Context) error {
		for {
			select {
			case m, ok := <-ch:
				if !ok {
					return nil
				}

				m.Raw = ""
				messages = append(messages, m)

			case <-ctx.Done():
				return ctx.Err()
			}
		}
	})

	return &messages
}

func (h *Helper) CollectFromServer() *[]*irc.Message {
	h.t.Helper()
	return h.CollectFromChannel(h.s.Incoming())
}

func (h *Helper) CollectFromConn(conn irc.Decoder) *[]*irc.Message {
	h.t.Helper()
	messages := []*irc.Message{}

	h.g.Go(func(ctx context.Context) error {
		h.t.Helper()

		for {
			m := &irc.Message{}
			if err := conn.Decode(m); err != nil {
				err = ignoreClose(err)
				assert.NilError(h.t, err)
				return nil
			}

			m.Raw = ""
			messages = append(messages, m)
		}
	})

	return &messages
}

func (h *Helper) ServerMessages() <-chan *irc.Message {
	return h.s.Incoming()
}

func (h *Helper) StopServer() {
	h.t.Helper()

	h.stopOnce.Do(func() {
		h.t.Helper()
		assert.NilError(h.t, h.s.Stop())
	})
}

func (h *Helper) StopServerErr() (err error) {
	h.t.Helper()

	h.stopOnce.Do(func() {
		h.t.Helper()
		err = h.s.Stop()
	})

	return err
}

func (h *Helper) Wait() {
	h.t.Helper()
	assert.NilError(h.t, h.g.Wait())
}

func (h *Helper) Sleep() {
	h.t.Helper()
	time.Sleep(h.SleepDur)
}

func (h *Helper) SendToServer(ctx context.Context, m *irc.Message) {
	h.t.Helper()
	assert.NilError(h.t, h.s.Send(ctx, m))
	h.Sleep()
}

func (h *Helper) Addr() string {
	h.t.Helper()
	return h.s.Addr()
}

func (h *Helper) Dial() irc.Conn {
	h.t.Helper()
	conn, err := h.s.Dial()
	assert.NilError(h.t, err)
	return conn
}

func (h *Helper) CloseConn(conn irc.Conn) {
	h.t.Helper()
	assert.NilError(h.t, ignoreClose(conn.Close()))
}

func (h *Helper) SendWithConn(conn irc.Encoder, m *irc.Message) {
	h.t.Helper()
	assert.NilError(h.t, conn.Encode(m))
	h.Sleep()
}

func (h *Helper) AssertMessages(gotP *[]*irc.Message, want ...*irc.Message) {
	h.t.Helper()

	assert.Assert(h.t, gotP != nil)

	got := *gotP

	if want == nil {
		want = []*irc.Message{}
	}

	if got == nil {
		got = []*irc.Message{}
	}

	assert.DeepEqual(h.t, want, got)
}