// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/api"
	"github.com/cloudfoundry/loggregatorlib/logmessage"
)

type FakeLogsRepository struct {
	RecentLogsForStub        func(appGuid string) ([]*logmessage.LogMessage, error)
	recentLogsForMutex       sync.RWMutex
	recentLogsForArgsForCall []struct {
		appGuid string
	}
	recentLogsForReturns struct {
		result1 []*logmessage.LogMessage
		result2 error
	}
	TailLogsForStub        func(appGuid string, onConnect func()) (<-chan *logmessage.LogMessage, error)
	tailLogsForMutex       sync.RWMutex
	tailLogsForArgsForCall []struct {
		appGuid   string
		onConnect func()
	}
	tailLogsForReturns struct {
		result1 <-chan *logmessage.LogMessage
		result2 error
	}
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
}

func (fake *FakeLogsRepository) RecentLogsFor(appGuid string) ([]*logmessage.LogMessage, error) {
	fake.recentLogsForMutex.Lock()
	fake.recentLogsForArgsForCall = append(fake.recentLogsForArgsForCall, struct {
		appGuid string
	}{appGuid})
	fake.recentLogsForMutex.Unlock()
	if fake.RecentLogsForStub != nil {
		return fake.RecentLogsForStub(appGuid)
	} else {
		return fake.recentLogsForReturns.result1, fake.recentLogsForReturns.result2
	}
}

func (fake *FakeLogsRepository) RecentLogsForCallCount() int {
	fake.recentLogsForMutex.RLock()
	defer fake.recentLogsForMutex.RUnlock()
	return len(fake.recentLogsForArgsForCall)
}

func (fake *FakeLogsRepository) RecentLogsForArgsForCall(i int) string {
	fake.recentLogsForMutex.RLock()
	defer fake.recentLogsForMutex.RUnlock()
	return fake.recentLogsForArgsForCall[i].appGuid
}

func (fake *FakeLogsRepository) RecentLogsForReturns(result1 []*logmessage.LogMessage, result2 error) {
	fake.RecentLogsForStub = nil
	fake.recentLogsForReturns = struct {
		result1 []*logmessage.LogMessage
		result2 error
	}{result1, result2}
}

func (fake *FakeLogsRepository) TailLogsFor(appGuid string, onConnect func()) (<-chan *logmessage.LogMessage, error) {
	fake.tailLogsForMutex.Lock()
	fake.tailLogsForArgsForCall = append(fake.tailLogsForArgsForCall, struct {
		appGuid   string
		onConnect func()
	}{appGuid, onConnect})
	fake.tailLogsForMutex.Unlock()
	if fake.TailLogsForStub != nil {
		return fake.TailLogsForStub(appGuid, onConnect)
	} else {
		return fake.tailLogsForReturns.result1, fake.tailLogsForReturns.result2
	}
}

func (fake *FakeLogsRepository) TailLogsForCallCount() int {
	fake.tailLogsForMutex.RLock()
	defer fake.tailLogsForMutex.RUnlock()
	return len(fake.tailLogsForArgsForCall)
}

func (fake *FakeLogsRepository) TailLogsForArgsForCall(i int) (string, func()) {
	fake.tailLogsForMutex.RLock()
	defer fake.tailLogsForMutex.RUnlock()
	return fake.tailLogsForArgsForCall[i].appGuid, fake.tailLogsForArgsForCall[i].onConnect
}

func (fake *FakeLogsRepository) TailLogsForReturns(result1 <-chan *logmessage.LogMessage, result2 error) {
	fake.TailLogsForStub = nil
	fake.tailLogsForReturns = struct {
		result1 <-chan *logmessage.LogMessage
		result2 error
	}{result1, result2}
}

func (fake *FakeLogsRepository) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		fake.CloseStub()
	}
}

func (fake *FakeLogsRepository) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

var _ api.LogsRepository = new(FakeLogsRepository)
