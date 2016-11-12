package runner_test

import (
	"path/filepath"

	. "github.com/odlp/inflight/runner"

	. "github.com/onsi/ginkgo/extensions/table"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grepper", func() {
	DescribeTable("file already has story ID",
		func(fileName string, expected bool) {
			g := Grepper{}
			path, _ := filepath.Abs("../runner/fixtures/" + fileName)
			Expect(g.FileAlreadyHasStoryID(path)).To(Equal(expected))
		},
		Entry("no story ID", "no-story-id.txt", false),
		Entry("story ID present", "story-id.txt", true),
		Entry("story ID with 'finishes'", "story-id-with-keyword.txt", true),
	)
})
