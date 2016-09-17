package runner

import (
	"os"
	"strconv"
)

const EnvGitAuthorEmail = "GIT_AUTHOR_EMAIL"
const EnvTrackerAPIToken = "TRACKER_API_TOKEN"
const EnvTrackerProjectID = "TRACKER_PROJECT_ID"

type Config struct {
	OutputPath       string
	GitAuthorEmail   string
	TrackerAPIToken  string
	TrackerProjectID int
}

func configWithOutputPath(outputPath string) Config {
	projectID, _ := strconv.Atoi(os.Getenv(EnvTrackerProjectID))
	return Config{
		OutputPath:       outputPath,
		GitAuthorEmail:   os.Getenv(EnvGitAuthorEmail),
		TrackerAPIToken:  os.Getenv(EnvTrackerAPIToken),
		TrackerProjectID: projectID,
	}
}
