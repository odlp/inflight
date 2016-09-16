// This fake has been hand edited while Counterfeiter
// cannot fake interfaces from the vendor folder:
// https://github.com/maxbrunsfeld/counterfeiter/issues/50

package projectfakes

import (
	"sync"

	"github.com/odlp/go-tracker"
	"github.com/odlp/inflight/project"
)

type FakeProjectClient struct {
	ProjectMembershipsStub        func() ([]tracker.ProjectMembership, error)
	projectMembershipsMutex       sync.RWMutex
	projectMembershipsArgsForCall []struct{}
	projectMembershipsReturns     struct {
		result1 []tracker.ProjectMembership
		result2 error
	}
	StoriesStub        func(query tracker.StoriesQuery) ([]tracker.Story, tracker.Pagination, error)
	storiesMutex       sync.RWMutex
	storiesArgsForCall []struct {
		query tracker.StoriesQuery
	}
	storiesReturns struct {
		result1 []tracker.Story
		result2 tracker.Pagination
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProjectClient) ProjectMemberships() ([]tracker.ProjectMembership, error) {
	fake.projectMembershipsMutex.Lock()
	fake.projectMembershipsArgsForCall = append(fake.projectMembershipsArgsForCall, struct{}{})
	fake.recordInvocation("ProjectMemberships", []interface{}{})
	fake.projectMembershipsMutex.Unlock()
	if fake.ProjectMembershipsStub != nil {
		return fake.ProjectMembershipsStub()
	} else {
		return fake.projectMembershipsReturns.result1, fake.projectMembershipsReturns.result2
	}
}

func (fake *FakeProjectClient) ProjectMembershipsCallCount() int {
	fake.projectMembershipsMutex.RLock()
	defer fake.projectMembershipsMutex.RUnlock()
	return len(fake.projectMembershipsArgsForCall)
}

func (fake *FakeProjectClient) ProjectMembershipsReturns(result1 []tracker.ProjectMembership, result2 error) {
	fake.ProjectMembershipsStub = nil
	fake.projectMembershipsReturns = struct {
		result1 []tracker.ProjectMembership
		result2 error
	}{result1, result2}
}

func (fake *FakeProjectClient) Stories(query tracker.StoriesQuery) ([]tracker.Story, tracker.Pagination, error) {
	fake.storiesMutex.Lock()
	fake.storiesArgsForCall = append(fake.storiesArgsForCall, struct {
		query tracker.StoriesQuery
	}{query})
	fake.recordInvocation("Stories", []interface{}{query})
	fake.storiesMutex.Unlock()
	if fake.StoriesStub != nil {
		return fake.StoriesStub(query)
	} else {
		return fake.storiesReturns.result1, fake.storiesReturns.result2, fake.storiesReturns.result3
	}
}

func (fake *FakeProjectClient) StoriesCallCount() int {
	fake.storiesMutex.RLock()
	defer fake.storiesMutex.RUnlock()
	return len(fake.storiesArgsForCall)
}

func (fake *FakeProjectClient) StoriesArgsForCall(i int) tracker.StoriesQuery {
	fake.storiesMutex.RLock()
	defer fake.storiesMutex.RUnlock()
	return fake.storiesArgsForCall[i].query
}

func (fake *FakeProjectClient) StoriesReturns(result1 []tracker.Story, result2 tracker.Pagination, result3 error) {
	fake.StoriesStub = nil
	fake.storiesReturns = struct {
		result1 []tracker.Story
		result2 tracker.Pagination
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeProjectClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.projectMembershipsMutex.RLock()
	defer fake.projectMembershipsMutex.RUnlock()
	fake.storiesMutex.RLock()
	defer fake.storiesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeProjectClient) recordInvocation(key string, args []interface{}) {
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

var _ project.ProjectClient = new(FakeProjectClient)
