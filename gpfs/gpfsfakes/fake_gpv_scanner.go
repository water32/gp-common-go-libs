// Code generated by counterfeiter. DO NOT EDIT.
package gpfsfakes

import (
	"sync"

	"github.com/greenplum-db/gp-common-go-libs/gpfs"
)

type FakeGpvScanner struct {
	ScanStub        func() bool
	scanMutex       sync.RWMutex
	scanArgsForCall []struct {
	}
	scanReturns struct {
		result1 bool
	}
	scanReturnsOnCall map[int]struct {
		result1 bool
	}
	TextStub        func() string
	textMutex       sync.RWMutex
	textArgsForCall []struct {
	}
	textReturns struct {
		result1 string
	}
	textReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGpvScanner) Scan() bool {
	fake.scanMutex.Lock()
	ret, specificReturn := fake.scanReturnsOnCall[len(fake.scanArgsForCall)]
	fake.scanArgsForCall = append(fake.scanArgsForCall, struct {
	}{})
	stub := fake.ScanStub
	fakeReturns := fake.scanReturns
	fake.recordInvocation("Scan", []interface{}{})
	fake.scanMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGpvScanner) ScanCallCount() int {
	fake.scanMutex.RLock()
	defer fake.scanMutex.RUnlock()
	return len(fake.scanArgsForCall)
}

func (fake *FakeGpvScanner) ScanCalls(stub func() bool) {
	fake.scanMutex.Lock()
	defer fake.scanMutex.Unlock()
	fake.ScanStub = stub
}

func (fake *FakeGpvScanner) ScanReturns(result1 bool) {
	fake.scanMutex.Lock()
	defer fake.scanMutex.Unlock()
	fake.ScanStub = nil
	fake.scanReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeGpvScanner) ScanReturnsOnCall(i int, result1 bool) {
	fake.scanMutex.Lock()
	defer fake.scanMutex.Unlock()
	fake.ScanStub = nil
	if fake.scanReturnsOnCall == nil {
		fake.scanReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.scanReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeGpvScanner) Text() string {
	fake.textMutex.Lock()
	ret, specificReturn := fake.textReturnsOnCall[len(fake.textArgsForCall)]
	fake.textArgsForCall = append(fake.textArgsForCall, struct {
	}{})
	stub := fake.TextStub
	fakeReturns := fake.textReturns
	fake.recordInvocation("Text", []interface{}{})
	fake.textMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGpvScanner) TextCallCount() int {
	fake.textMutex.RLock()
	defer fake.textMutex.RUnlock()
	return len(fake.textArgsForCall)
}

func (fake *FakeGpvScanner) TextCalls(stub func() string) {
	fake.textMutex.Lock()
	defer fake.textMutex.Unlock()
	fake.TextStub = stub
}

func (fake *FakeGpvScanner) TextReturns(result1 string) {
	fake.textMutex.Lock()
	defer fake.textMutex.Unlock()
	fake.TextStub = nil
	fake.textReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeGpvScanner) TextReturnsOnCall(i int, result1 string) {
	fake.textMutex.Lock()
	defer fake.textMutex.Unlock()
	fake.TextStub = nil
	if fake.textReturnsOnCall == nil {
		fake.textReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.textReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeGpvScanner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.scanMutex.RLock()
	defer fake.scanMutex.RUnlock()
	fake.textMutex.RLock()
	defer fake.textMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGpvScanner) recordInvocation(key string, args []interface{}) {
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

var _ gpfs.GpvScanner = new(FakeGpvScanner)