package project_test

import (
	. "github.com/odlp/inflight/project"
	"github.com/odlp/inflight/project/projectfakes"
	"github.com/odlp/inflight/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserCache.TryFindUser", func() {
	Context("cache hit", func() {
		It("returns a user with the corresponding initials", func() {
			c := cacheWithPath("./fixtures/cache.yaml")
			foundUser := c.TryFindUser("author@example.com")

			Expect(foundUser).ToNot(BeNil())
			Expect(foundUser.Person.Initials).To(Equal("AE"))
		})

		Context("cache miss", func() {
			It("returns nil", func() {
				c := cacheWithPath("./fixtures/cache-no-users.yaml")
				foundUser := c.TryFindUser("author@example.com")

				Expect(foundUser).To(BeNil())
			})
		})

		Context("cache does not exist", func() {
			It("returns nil", func() {
				c := cacheWithPath("/tmp/not-a-file")
				foundUser := c.TryFindUser("author@example.com")

				Expect(foundUser).To(BeNil())
			})
		})

		Context("cache is not yaml", func() {
			It("returns nil", func() {
				c := cacheWithPath("./fixtures/cache-wrong-format.json")
				foundUser := c.TryFindUser("author@example.com")

				Expect(foundUser).To(BeNil())
			})
		})
	})
})

var _ = Describe("UserCache.CacheFoundUser", func() {
	var fakeFileSystem projectfakes.FakeFileSystemInterface

	BeforeEach(func() {
		fakeFileSystem = projectfakes.FakeFileSystemInterface{}
	})

	Context("cache file exists", func() {
		It("inserts the user to cache YAML and writes back", func() {
			fakeFileSystem.ReadFromFileReturns([]byte("users"), nil)
			cachePath := "/tmp/some-file"

			c := UserCache{
				CachePath:  cachePath,
				FileSystem: &fakeFileSystem,
			}

			c.CacheFoundUser("someuser@example.com", "SU")

			Expect(fakeFileSystem.ReadFromFileCallCount()).To(Equal(1))
			Expect(fakeFileSystem.WriteToFileCallCount()).To(Equal(1))

			writtenCachePath, writtenYAML := fakeFileSystem.WriteToFileArgsForCall(0)
			Expect(writtenCachePath).To(Equal(cachePath))
			Expect(writtenYAML).To(ContainSubstring("someuser@example.com: SU"))
		})
	})
})

func cacheWithPath(path string) UserCache {
	return UserCache{
		CachePath:  path,
		FileSystem: util.FileSystem{},
	}
}
