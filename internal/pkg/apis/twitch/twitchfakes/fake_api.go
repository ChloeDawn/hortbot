// Code generated by counterfeiter. DO NOT EDIT.
package twitchfakes

import (
	"context"
	"sync"

	"github.com/hortbot/hortbot/internal/pkg/apis/twitch"
	"golang.org/x/oauth2"
)

type FakeAPI struct {
	GetChannelByIDStub        func(context.Context, int64) (*twitch.Channel, error)
	getChannelByIDMutex       sync.RWMutex
	getChannelByIDArgsForCall []struct {
		arg1 context.Context
		arg2 int64
	}
	getChannelByIDReturns struct {
		result1 *twitch.Channel
		result2 error
	}
	getChannelByIDReturnsOnCall map[int]struct {
		result1 *twitch.Channel
		result2 error
	}
	GetChattersStub        func(context.Context, string) (int64, error)
	getChattersMutex       sync.RWMutex
	getChattersArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getChattersReturns struct {
		result1 int64
		result2 error
	}
	getChattersReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	GetCurrentStreamStub        func(context.Context, int64) (*twitch.Stream, error)
	getCurrentStreamMutex       sync.RWMutex
	getCurrentStreamArgsForCall []struct {
		arg1 context.Context
		arg2 int64
	}
	getCurrentStreamReturns struct {
		result1 *twitch.Stream
		result2 error
	}
	getCurrentStreamReturnsOnCall map[int]struct {
		result1 *twitch.Stream
		result2 error
	}
	GetIDForTokenStub        func(context.Context, *oauth2.Token) (int64, *oauth2.Token, error)
	getIDForTokenMutex       sync.RWMutex
	getIDForTokenArgsForCall []struct {
		arg1 context.Context
		arg2 *oauth2.Token
	}
	getIDForTokenReturns struct {
		result1 int64
		result2 *oauth2.Token
		result3 error
	}
	getIDForTokenReturnsOnCall map[int]struct {
		result1 int64
		result2 *oauth2.Token
		result3 error
	}
	GetIDForUsernameStub        func(context.Context, string) (int64, error)
	getIDForUsernameMutex       sync.RWMutex
	getIDForUsernameArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getIDForUsernameReturns struct {
		result1 int64
		result2 error
	}
	getIDForUsernameReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	SetChannelGameStub        func(context.Context, int64, *oauth2.Token, string) (string, *oauth2.Token, error)
	setChannelGameMutex       sync.RWMutex
	setChannelGameArgsForCall []struct {
		arg1 context.Context
		arg2 int64
		arg3 *oauth2.Token
		arg4 string
	}
	setChannelGameReturns struct {
		result1 string
		result2 *oauth2.Token
		result3 error
	}
	setChannelGameReturnsOnCall map[int]struct {
		result1 string
		result2 *oauth2.Token
		result3 error
	}
	SetChannelStatusStub        func(context.Context, int64, *oauth2.Token, string) (string, *oauth2.Token, error)
	setChannelStatusMutex       sync.RWMutex
	setChannelStatusArgsForCall []struct {
		arg1 context.Context
		arg2 int64
		arg3 *oauth2.Token
		arg4 string
	}
	setChannelStatusReturns struct {
		result1 string
		result2 *oauth2.Token
		result3 error
	}
	setChannelStatusReturnsOnCall map[int]struct {
		result1 string
		result2 *oauth2.Token
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAPI) GetChannelByID(arg1 context.Context, arg2 int64) (*twitch.Channel, error) {
	fake.getChannelByIDMutex.Lock()
	ret, specificReturn := fake.getChannelByIDReturnsOnCall[len(fake.getChannelByIDArgsForCall)]
	fake.getChannelByIDArgsForCall = append(fake.getChannelByIDArgsForCall, struct {
		arg1 context.Context
		arg2 int64
	}{arg1, arg2})
	fake.recordInvocation("GetChannelByID", []interface{}{arg1, arg2})
	fake.getChannelByIDMutex.Unlock()
	if fake.GetChannelByIDStub != nil {
		return fake.GetChannelByIDStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getChannelByIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPI) GetChannelByIDCallCount() int {
	fake.getChannelByIDMutex.RLock()
	defer fake.getChannelByIDMutex.RUnlock()
	return len(fake.getChannelByIDArgsForCall)
}

func (fake *FakeAPI) GetChannelByIDCalls(stub func(context.Context, int64) (*twitch.Channel, error)) {
	fake.getChannelByIDMutex.Lock()
	defer fake.getChannelByIDMutex.Unlock()
	fake.GetChannelByIDStub = stub
}

func (fake *FakeAPI) GetChannelByIDArgsForCall(i int) (context.Context, int64) {
	fake.getChannelByIDMutex.RLock()
	defer fake.getChannelByIDMutex.RUnlock()
	argsForCall := fake.getChannelByIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPI) GetChannelByIDReturns(result1 *twitch.Channel, result2 error) {
	fake.getChannelByIDMutex.Lock()
	defer fake.getChannelByIDMutex.Unlock()
	fake.GetChannelByIDStub = nil
	fake.getChannelByIDReturns = struct {
		result1 *twitch.Channel
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetChannelByIDReturnsOnCall(i int, result1 *twitch.Channel, result2 error) {
	fake.getChannelByIDMutex.Lock()
	defer fake.getChannelByIDMutex.Unlock()
	fake.GetChannelByIDStub = nil
	if fake.getChannelByIDReturnsOnCall == nil {
		fake.getChannelByIDReturnsOnCall = make(map[int]struct {
			result1 *twitch.Channel
			result2 error
		})
	}
	fake.getChannelByIDReturnsOnCall[i] = struct {
		result1 *twitch.Channel
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetChatters(arg1 context.Context, arg2 string) (int64, error) {
	fake.getChattersMutex.Lock()
	ret, specificReturn := fake.getChattersReturnsOnCall[len(fake.getChattersArgsForCall)]
	fake.getChattersArgsForCall = append(fake.getChattersArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetChatters", []interface{}{arg1, arg2})
	fake.getChattersMutex.Unlock()
	if fake.GetChattersStub != nil {
		return fake.GetChattersStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getChattersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPI) GetChattersCallCount() int {
	fake.getChattersMutex.RLock()
	defer fake.getChattersMutex.RUnlock()
	return len(fake.getChattersArgsForCall)
}

func (fake *FakeAPI) GetChattersCalls(stub func(context.Context, string) (int64, error)) {
	fake.getChattersMutex.Lock()
	defer fake.getChattersMutex.Unlock()
	fake.GetChattersStub = stub
}

func (fake *FakeAPI) GetChattersArgsForCall(i int) (context.Context, string) {
	fake.getChattersMutex.RLock()
	defer fake.getChattersMutex.RUnlock()
	argsForCall := fake.getChattersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPI) GetChattersReturns(result1 int64, result2 error) {
	fake.getChattersMutex.Lock()
	defer fake.getChattersMutex.Unlock()
	fake.GetChattersStub = nil
	fake.getChattersReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetChattersReturnsOnCall(i int, result1 int64, result2 error) {
	fake.getChattersMutex.Lock()
	defer fake.getChattersMutex.Unlock()
	fake.GetChattersStub = nil
	if fake.getChattersReturnsOnCall == nil {
		fake.getChattersReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.getChattersReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetCurrentStream(arg1 context.Context, arg2 int64) (*twitch.Stream, error) {
	fake.getCurrentStreamMutex.Lock()
	ret, specificReturn := fake.getCurrentStreamReturnsOnCall[len(fake.getCurrentStreamArgsForCall)]
	fake.getCurrentStreamArgsForCall = append(fake.getCurrentStreamArgsForCall, struct {
		arg1 context.Context
		arg2 int64
	}{arg1, arg2})
	fake.recordInvocation("GetCurrentStream", []interface{}{arg1, arg2})
	fake.getCurrentStreamMutex.Unlock()
	if fake.GetCurrentStreamStub != nil {
		return fake.GetCurrentStreamStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getCurrentStreamReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPI) GetCurrentStreamCallCount() int {
	fake.getCurrentStreamMutex.RLock()
	defer fake.getCurrentStreamMutex.RUnlock()
	return len(fake.getCurrentStreamArgsForCall)
}

func (fake *FakeAPI) GetCurrentStreamCalls(stub func(context.Context, int64) (*twitch.Stream, error)) {
	fake.getCurrentStreamMutex.Lock()
	defer fake.getCurrentStreamMutex.Unlock()
	fake.GetCurrentStreamStub = stub
}

func (fake *FakeAPI) GetCurrentStreamArgsForCall(i int) (context.Context, int64) {
	fake.getCurrentStreamMutex.RLock()
	defer fake.getCurrentStreamMutex.RUnlock()
	argsForCall := fake.getCurrentStreamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPI) GetCurrentStreamReturns(result1 *twitch.Stream, result2 error) {
	fake.getCurrentStreamMutex.Lock()
	defer fake.getCurrentStreamMutex.Unlock()
	fake.GetCurrentStreamStub = nil
	fake.getCurrentStreamReturns = struct {
		result1 *twitch.Stream
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetCurrentStreamReturnsOnCall(i int, result1 *twitch.Stream, result2 error) {
	fake.getCurrentStreamMutex.Lock()
	defer fake.getCurrentStreamMutex.Unlock()
	fake.GetCurrentStreamStub = nil
	if fake.getCurrentStreamReturnsOnCall == nil {
		fake.getCurrentStreamReturnsOnCall = make(map[int]struct {
			result1 *twitch.Stream
			result2 error
		})
	}
	fake.getCurrentStreamReturnsOnCall[i] = struct {
		result1 *twitch.Stream
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetIDForToken(arg1 context.Context, arg2 *oauth2.Token) (int64, *oauth2.Token, error) {
	fake.getIDForTokenMutex.Lock()
	ret, specificReturn := fake.getIDForTokenReturnsOnCall[len(fake.getIDForTokenArgsForCall)]
	fake.getIDForTokenArgsForCall = append(fake.getIDForTokenArgsForCall, struct {
		arg1 context.Context
		arg2 *oauth2.Token
	}{arg1, arg2})
	fake.recordInvocation("GetIDForToken", []interface{}{arg1, arg2})
	fake.getIDForTokenMutex.Unlock()
	if fake.GetIDForTokenStub != nil {
		return fake.GetIDForTokenStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getIDForTokenReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeAPI) GetIDForTokenCallCount() int {
	fake.getIDForTokenMutex.RLock()
	defer fake.getIDForTokenMutex.RUnlock()
	return len(fake.getIDForTokenArgsForCall)
}

func (fake *FakeAPI) GetIDForTokenCalls(stub func(context.Context, *oauth2.Token) (int64, *oauth2.Token, error)) {
	fake.getIDForTokenMutex.Lock()
	defer fake.getIDForTokenMutex.Unlock()
	fake.GetIDForTokenStub = stub
}

func (fake *FakeAPI) GetIDForTokenArgsForCall(i int) (context.Context, *oauth2.Token) {
	fake.getIDForTokenMutex.RLock()
	defer fake.getIDForTokenMutex.RUnlock()
	argsForCall := fake.getIDForTokenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPI) GetIDForTokenReturns(result1 int64, result2 *oauth2.Token, result3 error) {
	fake.getIDForTokenMutex.Lock()
	defer fake.getIDForTokenMutex.Unlock()
	fake.GetIDForTokenStub = nil
	fake.getIDForTokenReturns = struct {
		result1 int64
		result2 *oauth2.Token
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAPI) GetIDForTokenReturnsOnCall(i int, result1 int64, result2 *oauth2.Token, result3 error) {
	fake.getIDForTokenMutex.Lock()
	defer fake.getIDForTokenMutex.Unlock()
	fake.GetIDForTokenStub = nil
	if fake.getIDForTokenReturnsOnCall == nil {
		fake.getIDForTokenReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 *oauth2.Token
			result3 error
		})
	}
	fake.getIDForTokenReturnsOnCall[i] = struct {
		result1 int64
		result2 *oauth2.Token
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAPI) GetIDForUsername(arg1 context.Context, arg2 string) (int64, error) {
	fake.getIDForUsernameMutex.Lock()
	ret, specificReturn := fake.getIDForUsernameReturnsOnCall[len(fake.getIDForUsernameArgsForCall)]
	fake.getIDForUsernameArgsForCall = append(fake.getIDForUsernameArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetIDForUsername", []interface{}{arg1, arg2})
	fake.getIDForUsernameMutex.Unlock()
	if fake.GetIDForUsernameStub != nil {
		return fake.GetIDForUsernameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getIDForUsernameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPI) GetIDForUsernameCallCount() int {
	fake.getIDForUsernameMutex.RLock()
	defer fake.getIDForUsernameMutex.RUnlock()
	return len(fake.getIDForUsernameArgsForCall)
}

func (fake *FakeAPI) GetIDForUsernameCalls(stub func(context.Context, string) (int64, error)) {
	fake.getIDForUsernameMutex.Lock()
	defer fake.getIDForUsernameMutex.Unlock()
	fake.GetIDForUsernameStub = stub
}

func (fake *FakeAPI) GetIDForUsernameArgsForCall(i int) (context.Context, string) {
	fake.getIDForUsernameMutex.RLock()
	defer fake.getIDForUsernameMutex.RUnlock()
	argsForCall := fake.getIDForUsernameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPI) GetIDForUsernameReturns(result1 int64, result2 error) {
	fake.getIDForUsernameMutex.Lock()
	defer fake.getIDForUsernameMutex.Unlock()
	fake.GetIDForUsernameStub = nil
	fake.getIDForUsernameReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetIDForUsernameReturnsOnCall(i int, result1 int64, result2 error) {
	fake.getIDForUsernameMutex.Lock()
	defer fake.getIDForUsernameMutex.Unlock()
	fake.GetIDForUsernameStub = nil
	if fake.getIDForUsernameReturnsOnCall == nil {
		fake.getIDForUsernameReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.getIDForUsernameReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) SetChannelGame(arg1 context.Context, arg2 int64, arg3 *oauth2.Token, arg4 string) (string, *oauth2.Token, error) {
	fake.setChannelGameMutex.Lock()
	ret, specificReturn := fake.setChannelGameReturnsOnCall[len(fake.setChannelGameArgsForCall)]
	fake.setChannelGameArgsForCall = append(fake.setChannelGameArgsForCall, struct {
		arg1 context.Context
		arg2 int64
		arg3 *oauth2.Token
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("SetChannelGame", []interface{}{arg1, arg2, arg3, arg4})
	fake.setChannelGameMutex.Unlock()
	if fake.SetChannelGameStub != nil {
		return fake.SetChannelGameStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.setChannelGameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeAPI) SetChannelGameCallCount() int {
	fake.setChannelGameMutex.RLock()
	defer fake.setChannelGameMutex.RUnlock()
	return len(fake.setChannelGameArgsForCall)
}

func (fake *FakeAPI) SetChannelGameCalls(stub func(context.Context, int64, *oauth2.Token, string) (string, *oauth2.Token, error)) {
	fake.setChannelGameMutex.Lock()
	defer fake.setChannelGameMutex.Unlock()
	fake.SetChannelGameStub = stub
}

func (fake *FakeAPI) SetChannelGameArgsForCall(i int) (context.Context, int64, *oauth2.Token, string) {
	fake.setChannelGameMutex.RLock()
	defer fake.setChannelGameMutex.RUnlock()
	argsForCall := fake.setChannelGameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeAPI) SetChannelGameReturns(result1 string, result2 *oauth2.Token, result3 error) {
	fake.setChannelGameMutex.Lock()
	defer fake.setChannelGameMutex.Unlock()
	fake.SetChannelGameStub = nil
	fake.setChannelGameReturns = struct {
		result1 string
		result2 *oauth2.Token
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAPI) SetChannelGameReturnsOnCall(i int, result1 string, result2 *oauth2.Token, result3 error) {
	fake.setChannelGameMutex.Lock()
	defer fake.setChannelGameMutex.Unlock()
	fake.SetChannelGameStub = nil
	if fake.setChannelGameReturnsOnCall == nil {
		fake.setChannelGameReturnsOnCall = make(map[int]struct {
			result1 string
			result2 *oauth2.Token
			result3 error
		})
	}
	fake.setChannelGameReturnsOnCall[i] = struct {
		result1 string
		result2 *oauth2.Token
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAPI) SetChannelStatus(arg1 context.Context, arg2 int64, arg3 *oauth2.Token, arg4 string) (string, *oauth2.Token, error) {
	fake.setChannelStatusMutex.Lock()
	ret, specificReturn := fake.setChannelStatusReturnsOnCall[len(fake.setChannelStatusArgsForCall)]
	fake.setChannelStatusArgsForCall = append(fake.setChannelStatusArgsForCall, struct {
		arg1 context.Context
		arg2 int64
		arg3 *oauth2.Token
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("SetChannelStatus", []interface{}{arg1, arg2, arg3, arg4})
	fake.setChannelStatusMutex.Unlock()
	if fake.SetChannelStatusStub != nil {
		return fake.SetChannelStatusStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.setChannelStatusReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeAPI) SetChannelStatusCallCount() int {
	fake.setChannelStatusMutex.RLock()
	defer fake.setChannelStatusMutex.RUnlock()
	return len(fake.setChannelStatusArgsForCall)
}

func (fake *FakeAPI) SetChannelStatusCalls(stub func(context.Context, int64, *oauth2.Token, string) (string, *oauth2.Token, error)) {
	fake.setChannelStatusMutex.Lock()
	defer fake.setChannelStatusMutex.Unlock()
	fake.SetChannelStatusStub = stub
}

func (fake *FakeAPI) SetChannelStatusArgsForCall(i int) (context.Context, int64, *oauth2.Token, string) {
	fake.setChannelStatusMutex.RLock()
	defer fake.setChannelStatusMutex.RUnlock()
	argsForCall := fake.setChannelStatusArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeAPI) SetChannelStatusReturns(result1 string, result2 *oauth2.Token, result3 error) {
	fake.setChannelStatusMutex.Lock()
	defer fake.setChannelStatusMutex.Unlock()
	fake.SetChannelStatusStub = nil
	fake.setChannelStatusReturns = struct {
		result1 string
		result2 *oauth2.Token
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAPI) SetChannelStatusReturnsOnCall(i int, result1 string, result2 *oauth2.Token, result3 error) {
	fake.setChannelStatusMutex.Lock()
	defer fake.setChannelStatusMutex.Unlock()
	fake.SetChannelStatusStub = nil
	if fake.setChannelStatusReturnsOnCall == nil {
		fake.setChannelStatusReturnsOnCall = make(map[int]struct {
			result1 string
			result2 *oauth2.Token
			result3 error
		})
	}
	fake.setChannelStatusReturnsOnCall[i] = struct {
		result1 string
		result2 *oauth2.Token
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getChannelByIDMutex.RLock()
	defer fake.getChannelByIDMutex.RUnlock()
	fake.getChattersMutex.RLock()
	defer fake.getChattersMutex.RUnlock()
	fake.getCurrentStreamMutex.RLock()
	defer fake.getCurrentStreamMutex.RUnlock()
	fake.getIDForTokenMutex.RLock()
	defer fake.getIDForTokenMutex.RUnlock()
	fake.getIDForUsernameMutex.RLock()
	defer fake.getIDForUsernameMutex.RUnlock()
	fake.setChannelGameMutex.RLock()
	defer fake.setChannelGameMutex.RUnlock()
	fake.setChannelStatusMutex.RLock()
	defer fake.setChannelStatusMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAPI) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ twitch.API = new(FakeAPI)