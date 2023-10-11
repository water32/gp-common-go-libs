// Code generated by counterfeiter. DO NOT EDIT.
package gpfsfakes

import (
	"sync"

	"github.com/greenplum-db/gp-common-go-libs/gpfs"
)

type FakeGpRegexp struct {
	MatchStringStub        func(string) bool
	matchStringMutex       sync.RWMutex
	matchStringArgsForCall []struct {
		arg1 string
	}
	matchStringReturns struct {
		result1 bool
	}
	matchStringReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGpRegexp) MatchString(arg1 string) bool {
	fake.matchStringMutex.Lock()
	ret, specificReturn := fake.matchStringReturnsOnCall[len(fake.matchStringArgsForCall)]
	fake.matchStringArgsForCall = append(fake.matchStringArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.MatchStringStub
	fakeReturns := fake.matchStringReturns
	fake.recordInvocation("MatchString", []interface{}{arg1})
	fake.matchStringMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGpRegexp) MatchStringCallCount() int {
	fake.matchStringMutex.RLock()
	defer fake.matchStringMutex.RUnlock()
	return len(fake.matchStringArgsForCall)
}

func (fake *FakeGpRegexp) MatchStringCalls(stub func(string) bool) {
	fake.matchStringMutex.Lock()
	defer fake.matchStringMutex.Unlock()
	fake.MatchStringStub = stub
}

func (fake *FakeGpRegexp) MatchStringArgsForCall(i int) string {
	fake.matchStringMutex.RLock()
	defer fake.matchStringMutex.RUnlock()
	argsForCall := fake.matchStringArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGpRegexp) MatchStringReturns(result1 bool) {
	fake.matchStringMutex.Lock()
	defer fake.matchStringMutex.Unlock()
	fake.MatchStringStub = nil
	fake.matchStringReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeGpRegexp) MatchStringReturnsOnCall(i int, result1 bool) {
	fake.matchStringMutex.Lock()
	defer fake.matchStringMutex.Unlock()
	fake.MatchStringStub = nil
	if fake.matchStringReturnsOnCall == nil {
		fake.matchStringReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.matchStringReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeGpRegexp) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.matchStringMutex.RLock()
	defer fake.matchStringMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGpRegexp) recordInvocation(key string, args []interface{}) {
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

var _ gpfs.GpRegexp = new(FakeGpRegexp)