package runner_test

import (
	"os"

	"github.com/odlp/go-tracker"
	. "github.com/odlp/inflight/runner"
	"github.com/odlp/inflight/runner/runnerfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewRunner", func() {
	BeforeEach(func() {
		os.Setenv(EnvGitAuthorEmail, "a@example.com")
		os.Setenv(EnvTrackerAPIToken, "api-token")
		os.Setenv(EnvTrackerProjectID, "123")
	})

	It("loads the config from the environment", func() {
		r := NewRunner("./tmp/foo.txt")

		Expect(r.Config.GitAuthorEmail).To(Equal("a@example.com"))
		Expect(r.Config.TrackerAPIToken).To(Equal("api-token"))
		Expect(r.Config.TrackerProjectID).To(Equal(123))
	})
})

var _ = Describe("Exec", func() {
	var (
		r           Runner
		c           Config
		fakeProject runnerfakes.FakeProjectInterface
		fakeWriter  runnerfakes.FakeWriterInterface
		user        tracker.ProjectMembership
		story       tracker.Story
	)

	BeforeEach(func() {
		fakeProject = runnerfakes.FakeProjectInterface{}
		fakeWriter = runnerfakes.FakeWriterInterface{}

		c = Config{
			GitAuthorEmail: "a@example.com",
			OutputPath:     "./tmp/example.txt",
		}

		r = Runner{
			Project: &fakeProject,
			Writer:  &fakeWriter,
			Config:  c,
		}

		user = tracker.ProjectMembership{
			Person: tracker.Person{Initials: "DV"},
		}
		fakeProject.FindUserByEmailReturns(user, nil)

		story = tracker.Story{
			ID: 123,
		}
		fakeProject.FindCurrentStoryReturns(story, nil)

		r.Exec()
	})

	It("finds the user", func() {
		Expect(fakeProject.FindUserByEmailCallCount()).To(Equal(1))
		Expect(fakeProject.FindUserByEmailArgsForCall(0)).To(Equal(c.GitAuthorEmail))
	})

	It("finds the current story", func() {
		Expect(fakeProject.FindCurrentStoryCallCount()).To(Equal(1))
		Expect(fakeProject.FindCurrentStoryArgsForCall(0)).To(Equal(user))
	})

	It("passes the story ID to be written", func() {
		Expect(fakeWriter.WriteToFileCallCount()).To(Equal(1))

		capturedOutputPath, capturedOutputText := fakeWriter.WriteToFileArgsForCall(0)

		Expect(capturedOutputPath).To(Equal(c.OutputPath))
		Expect(capturedOutputText).To(Equal("[#123]\n"))
	})

})
