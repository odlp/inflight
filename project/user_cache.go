package project

import (
	"github.com/odlp/go-tracker"
	"gopkg.in/yaml.v2"
)

type UserCache struct {
	CachePath  string
	FileSystem FileSystemInterface
}

//go:generate counterfeiter . FileSystemInterface
type FileSystemInterface interface {
	WriteToFile(filePath string, text string)
	ReadFromFile(filePath string) ([]byte, error)
}

type userCacheFile struct {
	Users map[string]string `yaml:"users"`
}

func (c UserCache) TryFindUser(email string) *tracker.ProjectMembership {
	cacheData, err := c.FileSystem.ReadFromFile(c.CachePath)
	if err != nil {
		return nil
	}

	cache := userCacheFile{}

	yaml.Unmarshal(cacheData, &cache)
	initials := cache.Users[email]

	if initials == "" {
		return nil
	}

	return &tracker.ProjectMembership{
		Person: tracker.Person{
			Initials: initials,
		},
	}
}

func (c UserCache) CacheFoundUser(email, initials string) {
	cacheData, _ := c.FileSystem.ReadFromFile(c.CachePath)

	cache := userCacheFile{}
	yaml.Unmarshal(cacheData, &cache)

	if cache.Users == nil {
		cache.Users = map[string]string{}
	}

	cache.Users[email] = initials

	cacheOut, _ := yaml.Marshal(cache)

	c.FileSystem.WriteToFile(c.CachePath, string(cacheOut))
}
