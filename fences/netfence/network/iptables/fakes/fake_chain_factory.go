// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-linux/fences/netfence/network/iptables"
)

type FakeChainFactory struct {
	CreateChainStub        func(name string) iptables.Chain
	createChainMutex       sync.RWMutex
	createChainArgsForCall []struct {
		name string
	}
	createChainReturns struct {
		result1 iptables.Chain
	}
}

func (fake *FakeChainFactory) CreateChain(name string) iptables.Chain {
	fake.createChainMutex.Lock()
	fake.createChainArgsForCall = append(fake.createChainArgsForCall, struct {
		name string
	}{name})
	fake.createChainMutex.Unlock()
	if fake.CreateChainStub != nil {
		return fake.CreateChainStub(name)
	} else {
		return fake.createChainReturns.result1
	}
}

func (fake *FakeChainFactory) CreateChainCallCount() int {
	fake.createChainMutex.RLock()
	defer fake.createChainMutex.RUnlock()
	return len(fake.createChainArgsForCall)
}

func (fake *FakeChainFactory) CreateChainArgsForCall(i int) string {
	fake.createChainMutex.RLock()
	defer fake.createChainMutex.RUnlock()
	return fake.createChainArgsForCall[i].name
}

func (fake *FakeChainFactory) CreateChainReturns(result1 iptables.Chain) {
	fake.CreateChainStub = nil
	fake.createChainReturns = struct {
		result1 iptables.Chain
	}{result1}
}

var _ iptables.ChainFactory = new(FakeChainFactory)
