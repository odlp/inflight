package runner

import (
	"fmt"
	"os"

	"github.com/odlp/inflight/project"
)

type Runner struct {
	Project project.Project
	Writer  Writer
	Config  Config
}

func NewRunner(outputPath string) Runner {
	c := configWithOutputPath(outputPath)
	p := project.NewProject(c.TrackerAPIToken, c.TrackerProjectID)
	return Runner{
		Project: p,
		Writer:  Writer{},
		Config:  c,
	}
}

func (r Runner) Exec() {
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
