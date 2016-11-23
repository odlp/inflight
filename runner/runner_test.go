package runner_test

import (
	"os"

	"github.com/odlp/go-tracker"
	"github.com/odlp/inflight/project"
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

	It("passes the cache path to the project cache", func() {
		outputPath := "./tmp/foo.txt"
		r := NewRunner(outputPath)

		p := r.Project.(project.Project)
		cache := p.Cache.(project.UserCache)
		Expect(cache.CachePath).To(HaveSuffix("tmp/.inflight-cache"))
	})
})

var _ = Describe("Exec", func() {
	var (
		r              Runner
		c              Config
		fakeProject    runnerfakes.FakeProjectInterface
		fakeFileSystem runnerfakes.FakeFileSystemInterface
		fakeGrepper    runnerfakes.FakeGrepInterface
		user           tracker.ProjectMembership
		story          tracker.Story
	)

	BeforeEach(func() {
		fakeProject = runnerfakes.FakeProjectInterface{}
		fakeFileSystem = runnerfakes.FakeFileSystemInterface{}
		fakeGrepper = runnerfakes.FakeGrepInterface{}

		c = Config{
			GitAuthorEmail: "a@example.com",
			OutputPath:     "./tmp/example.txt",
		}
	})

	JustBeforeEach(func() {
		r = Runner{
			Project:    &fakeProject,
			FileSystem: &fakeFileSystem,
			Grepper:    &fakeGrepper,
			Config:     c,
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
			Expect(fakeFileSystem.WriteToFileCallCount()).To(Equal(1))

			capturedOutputPath, capturedOutputText := fakeFileSystem.WriteToFileArgsForCall(0)

			Expect(capturedOutputPath).To(Equal(c.OutputPath))
			Expect(capturedOutputText).To(Equal("\n[#123]\n\n"))
		})
	})

	Describe("when story ID is already present", func() {
		BeforeEach(func() {
			fakeGrepper.FileAlreadyHasStoryIDReturns(true)
		})

		It("does not fetch the story ID", func() {
			Expect(fakeProject.FindUserByEmailCallCount()).To(Equal(0))
			Expect(fakeProject.FindCurrentStoryCallCount()).To(Equal(0))
			Expect(fakeFileSystem.WriteToFileCallCount()).To(Equal(0))
		})
	})

})
