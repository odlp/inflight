// This fake has been hand edited while Counterfeiter
// cannot fake interfaces from the vendor folder:
// https://github.com/maxbrunsfeld/counterfeiter/issues/50

package runnerfakes

import (
	"sync"

	"github.com/odlp/go-tracker"
	"github.com/odlp/inflight/runner"
)

type FakeProjectInterface struct {
	FindUserByEmailStub        func(email string) (tracker.ProjectMembership, error)
	findUserByEmailMutex       sync.RWMutex
	findUserByEmailArgsForCall []struct {
		email string
	}
	findUserByEmailReturns struct {
		result1 tracker.ProjectMembership
		result2 error
	}
	FindCurrentStoryStub        func(member tracker.ProjectMembership) (tracker.Story, error)
	findCurrentStoryMutex       sync.RWMutex
	findCurrentStoryArgsForCall []struct {
		member tracker.ProjectMembership
	}
	findCurrentStoryReturns struct {
		result1 tracker.Story
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProjectInterface) FindUserByEmail(email string) (tracker.ProjectMembership, error) {
	fake.findUserByEmailMutex.Lock()
	fake.findUserByEmailArgsForCall = append(fake.findUserByEmailArgsForCall, struct {
		email string
	}{email})
	fake.recordInvocation("FindUserByEmail", []interface{}{email})
	fake.findUserByEmailMutex.Unlock()
	if fake.FindUserByEmailStub != nil {
		return fake.FindUserByEmailStub(email)
	} else {
		return fake.findUserByEmailReturns.result1, fake.findUserByEmailReturns.result2
	}
}

func (fake *FakeProjectInterface) FindUserByEmailCallCount() int {
	fake.findUserByEmailMutex.RLock()
	defer fake.findUserByEmailMutex.RUnlock()
	return len(fake.findUserByEmailArgsForCall)
}

func (fake *FakeProjectInterface) FindUserByEmailArgsForCall(i int) string {
	fake.findUserByEmailMutex.RLock()
	defer fake.findUserByEmailMutex.RUnlock()
	return fake.findUserByEmailArgsForCall[i].email
}

func (fake *FakeProjectInterface) FindUserByEmailReturns(result1 tracker.ProjectMembership, result2 error) {
	fake.FindUserByEmailStub = nil
	fake.findUserByEmailReturns = struct {
		result1 tracker.ProjectMembership
		result2 error
	}{result1, result2}
}

func (fake *FakeProjectInterface) FindCurrentStory(member tracker.ProjectMembership) (tracker.Story, error) {
	fake.findCurrentStoryMutex.Lock()
	fake.findCurrentStoryArgsForCall = append(fake.findCurrentStoryArgsForCall, struct {
		member tracker.ProjectMembership
	}{member})
	fake.recordInvocation("FindCurrentStory", []interface{}{member})
	fake.findCurrentStoryMutex.Unlock()
	if fake.FindCurrentStoryStub != nil {
		return fake.FindCurrentStoryStub(member)
	} else {
		return fake.findCurrentStoryReturns.result1, fake.findCurrentStoryReturns.result2
	}
}

func (fake *FakeProjectInterface) FindCurrentStoryCallCount() int {
	fake.findCurrentStoryMutex.RLock()
	defer fake.findCurrentStoryMutex.RUnlock()
	return len(fake.findCurrentStoryArgsForCall)
}

func (fake *FakeProjectInterface) FindCurrentStoryArgsForCall(i int) tracker.ProjectMembership {
	fake.findCurrentStoryMutex.RLock()
	defer fake.findCurrentStoryMutex.RUnlock()
	return fake.findCurrentStoryArgsForCall[i].member
}

func (fake *FakeProjectInterface) FindCurrentStoryReturns(result1 tracker.Story, result2 error) {
	fake.FindCurrentStoryStub = nil
	fake.findCurrentStoryReturns = struct {
		result1 tracker.Story
		result2 error
	}{result1, result2}
}

func (fake *FakeProjectInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findUserByEmailMutex.RLock()
	defer fake.findUserByEmailMutex.RUnlock()
	fake.findCurrentStoryMutex.RLock()
	defer fake.findCurrentStoryMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeProjectInterface) recordInvocation(key string, args []interface{}) {
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

var _ runner.ProjectInterface = new(FakeProjectInterface)
