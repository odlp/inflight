package project_test

import (
	"fmt"
	"time"

	"github.com/odlp/go-tracker"
	"github.com/odlp/inflight/project"
	"github.com/odlp/inflight/project/projectfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Project", func() {
	var (
		p          project.Project
		fakeClient projectfakes.FakeProjectClient
	)

	BeforeEach(func() {
		fakeClient = projectfakes.FakeProjectClient{}
		p = project.Project{
			Client: &fakeClient,
		}
	})

	Describe("FindUserByEmail", func() {
		It("returns the first member with a matching email", func() {
			user1 := createUser("a@example.com")
			user2 := createUser("b@example.com")
			user3 := createUser("c@example.com")
			users := []tracker.ProjectMembership{user1, user2, user3}

			fakeClient.ProjectMembershipsReturns(users, nil)

			foundUser, err := p.FindUserByEmail("b@example.com")

			Expect(err).ToNot(HaveOccurred())
			Expect(foundUser).To(Equal(user2))
		})

		Context("user cannot be found", func() {
			It("returns an error", func() {
				_, err := p.FindUserByEmail("ghost@example.com")
				Expect(err).To(MatchError("Unable to find 'ghost@example.com' in project members"))
			})
		})

		Context("project client returns an error", func() {
			It("propagates the error", func() {
				expectedError := fmt.Errorf("Oops")
				fakeClient.ProjectMembershipsReturns([]tracker.ProjectMembership{}, expectedError)

				_, err := p.FindUserByEmail("ghost@example.com")
				Expect(err).To(Equal(expectedError))
			})
		})
	})

	var _ = Describe("FindCurrentStory", func() {
		var member tracker.ProjectMembership

		BeforeEach(func() {
			member = tracker.ProjectMembership{
				Person: tracker.Person{Initials: "AA", Email: "a@example.com"},
			}
		})

		It("queries for the started stories owned by the user", func() {
			s := tracker.Story{ID: 2, UpdatedAt: timeHoursAgo(1)}
			fakeClient.StoriesReturns([]tracker.Story{s}, tracker.Pagination{}, nil)

			p.FindCurrentStory(member)

			Expect(fakeClient.StoriesCallCount()).To(Equal(1))

			capturedQuery := fakeClient.StoriesArgsForCall(0)
			Expect(capturedQuery.Filter).To(ConsistOf("owner:AA", "state:started"))
		})

		It("filters the returned stories by most recently updated", func() {
			oldestStory := tracker.Story{ID: 1, UpdatedAt: timeHoursAgo(24)}
			newestStory := tracker.Story{ID: 2, UpdatedAt: timeHoursAgo(1)}
			oldStory := tracker.Story{ID: 3, UpdatedAt: timeHoursAgo(12)}

			fakeClient.StoriesReturns([]tracker.Story{oldestStory, newestStory, oldStory}, tracker.Pagination{}, nil)

			foundStory, err := p.FindCurrentStory(member)

			Expect(err).ToNot(HaveOccurred())
			Expect(foundStory).To(Equal(newestStory))
		})

		Context("no stories found", func() {
			It("returns an error", func() {
				_, err := p.FindCurrentStory(member)

				Expect(err).To(MatchError("Unable to find current story for user 'a@example.com'"))
			})
		})

		Context("project client returns an error", func() {
			It("propagates the error", func() {
				expectedError := fmt.Errorf("Oops")
				fakeClient.StoriesReturns([]tracker.Story{}, tracker.Pagination{}, expectedError)

				_, err := p.FindCurrentStory(member)
				Expect(err).To(Equal(expectedError))
			})
		})
	})
})

var _ = Describe("NewProject", func() {
	It("returns a project with a client wired up", func() {
		p := project.NewProject("api-token", 123)
		Expect(p.Client).ToNot(BeNil())
	})
})

func createUser(email string) tracker.ProjectMembership {
	return tracker.ProjectMembership{
		Person: tracker.Person{
			Email: email,
		},
	}
}

func timeHoursAgo(n int64) *time.Time {
	duration := time.Duration(n) * -time.Hour
	t := time.Now().Add(duration)
	return &t
}
