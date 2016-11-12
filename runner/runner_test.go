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
		outputPath := "./tmp/foo.txt"
		r := NewRunner(outputPath)

		Expect(r.Config.OutputPath).To(Equal(outputPath))
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
		fakeGrepper runnerfakes.FakeGrepInterface
		user        tracker.ProjectMembership
		story       tracker.Story
	)

	BeforeEach(func() {
		fakeProject = runnerfakes.FakeProjectInterface{}
		fakeWriter = runnerfakes.FakeWriterInterface{}
		fakeGrepper = runnerfakes.FakeGrepInterface{}

		c = Config{
			GitAuthorEmail: "a@example.com",
			OutputPath:     "./tmp/example.txt",
		}
	})

	JustBeforeEach(func() {
		r = Runner{
			Project: &fakeProject,
			Writer:  &fakeWriter,
			Grepper: &fakeGrepper,
			Config:  c,
		}

		r.Exec()
	})

	Context("when story ID is needed", func() {
		BeforeEach(func() {
			fakeGrepper.FileAlreadyHasStoryIDReturns(false)

			user = tracker.ProjectMembership{
				Person: tracker.Person{Initials: "DV"},
			}
			fakeProject.FindUserByEmailReturns(user, nil)

			story = tracker.Story{
				ID: 123,
			}
			fakeProject.FindCurrentStoryReturns(story, nil)
		})

		It("fetches the user", func() {
			Expect(fakeProject.FindUserByEmailCallCount()).To(Equal(1))
			Expect(fakeProject.FindUserByEmailArgsForCall(0)).To(Equal(c.GitAuthorEmail))
		})

		It("fetches the current story", func() {
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

	Describe("when story ID is already present", func() {
		BeforeEach(func() {
			fakeGrepper.FileAlreadyHasStoryIDReturns(true)
		})

		It("does not fetch the story ID", func() {
			Expect(fakeProject.FindUserByEmailCallCount()).To(Equal(0))
			Expect(fakeProject.FindCurrentStoryCallCount()).To(Equal(0))
			Expect(fakeWriter.WriteToFileCallCount()).To(Equal(0))
		})
	})

})
