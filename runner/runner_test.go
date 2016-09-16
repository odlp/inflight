package runner_test

import (
	"os"

	. "github.com/odlp/inflight/runner"

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
