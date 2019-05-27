package bot

import (
	"fmt"
	"sync"
)

var isTesting = false // TODO: check for benchmark?

type testingPanic string

func (t testingPanic) String() string {
	return string(t)
}

type testingHelper struct {
	mu      sync.Mutex
	userIDs map[string]int64
	names   map[int64]string
}

func (t *testingHelper) checkUserNameID(name string, id int64) {
	if t == nil {
		return
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	if t.userIDs == nil {
		t.userIDs = make(map[string]int64)
	}

	if t.names == nil {
		t.names = make(map[int64]string)
	}

	if expectedID, ok := t.userIDs[name]; ok {
		if id != expectedID {
			panic(fmt.Sprintf("%v previously had id %v, now %v", name, expectedID, id))
		}
		return
	}

	if expectedName, ok := t.names[id]; ok {
		if name != expectedName {
			panic(testingPanic(fmt.Sprintf("%v previously had name %v, now %v", id, expectedName, name)))
		}
		return
	}

	t.userIDs[name] = id
	t.names[id] = name
}