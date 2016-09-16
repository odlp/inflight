package runner

import (
	"fmt"
	"os"
	"strconv"

	"github.com/odlp/inflight/project"
)

type Config struct {
	OutputPath       string
	GitAuthorEmail   string
	TrackerAPIToken  string
	TrackerProjectID int
}

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
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s, err := r.Project.FindCurrentStory(u)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outputText := fmt.Sprintf("[#%d]\n", s.ID)
	r.Writer.WriteToFile(r.Config.OutputPath, outputText)
}

const EnvGitAuthorEmail = "GIT_AUTHOR_EMAIL"
const EnvTrackerAPIToken = "TRACKER_API_TOKEN"
const EnvTrackerProjectID = "TRACKER_PROJECT_ID"

func configWithOutputPath(outputPath string) Config {
	projectID, _ := strconv.Atoi(os.Getenv(EnvTrackerProjectID))
	return Config{
		OutputPath:       outputPath,
		GitAuthorEmail:   os.Getenv(EnvGitAuthorEmail),
		TrackerAPIToken:  os.Getenv(EnvTrackerAPIToken),
		TrackerProjectID: projectID,
	}
}
