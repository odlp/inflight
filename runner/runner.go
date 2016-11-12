package runner

import (
	"fmt"
	"os"

	"github.com/odlp/go-tracker"
	"github.com/odlp/inflight/project"
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

//go:generate counterfeiter . WriterInterface
type WriterInterface interface {
	WriteToFile(filePath string, text string)
}

type Runner struct {
	Project ProjectInterface
	Writer  WriterInterface
	Grepper GrepInterface
	Config  Config
}

func NewRunner(outputPath string) Runner {
	c := configWithOutputPath(outputPath)
	p := project.NewProject(c.TrackerAPIToken, c.TrackerProjectID)
	return Runner{
		Project: p,
		Writer:  Writer{},
		Grepper: Grepper{},
		Config:  c,
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

	outputText := fmt.Sprintf("[#%d]\n", s.ID)
	r.Writer.WriteToFile(r.Config.OutputPath, outputText)
}

func gracefulExitIfError(err error) {
	if err != nil {
		fmt.Printf("Inflight: %s\n", err.Error())
		os.Exit(0)
	}
}
