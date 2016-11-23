package runner

import (
	"fmt"
	"os"

	tracker "github.com/odlp/go-tracker"
	"github.com/odlp/inflight/project"
	"github.com/odlp/inflight/util"
)

//go:generate counterfeiter . ProjectInterface
type ProjectInterface interface {
	FindUserByEmail(email string) (tracker.ProjectMembership, error)
	FindCurrentStory(member tracker.ProjectMembership) (tracker.Story, error)
}

//go:generate counterfeiter . GrepInterface
type GrepInterface interface {
	FileAlreadyHasStoryID(filePath string) bool
}

//go:generate counterfeiter . FileSystemInterface
type FileSystemInterface interface {
	WriteToFile(filePath string, text string)
}

type Runner struct {
	Project    ProjectInterface
	FileSystem FileSystemInterface
	Grepper    GrepInterface
	Config     Config
}

func NewRunner(outputPath string) Runner {
	c := configWithOutputPath(outputPath)
	p := project.NewProject(c.TrackerAPIToken, c.TrackerProjectID, c.CachePath)
	return Runner{
		Project:    p,
		FileSystem: util.FileSystem{},
		Grepper:    Grepper{},
		Config:     c,
	}
}

func (r Runner) Exec() {
	if r.Grepper.FileAlreadyHasStoryID(r.Config.OutputPath) {
		return
	}

	u, err := r.Project.FindUserByEmail(r.Config.GitAuthorEmail)
	gracefulExitIfError(err)

	s, err := r.Project.FindCurrentStory(u)
	gracefulExitIfError(err)

	r.FileSystem.WriteToFile(r.Config.OutputPath, formatStoryID(s.ID))
}

func formatStoryID(storyID int) string {
	return fmt.Sprintf("\n[#%d]\n\n", storyID)
}

func gracefulExitIfError(err error) {
	if err != nil {
		fmt.Printf("Inflight: %s\n", err.Error())
		os.Exit(0)
	}
}
