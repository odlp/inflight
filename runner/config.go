package runner

import (
	"os"
	"path/filepath"
	"strconv"
)

const EnvGitAuthorEmail = "GIT_AUTHOR_EMAIL"
const EnvTrackerAPIToken = "TRACKER_API_TOKEN"
const EnvTrackerProjectID = "TRACKER_PROJECT_ID"

type Config struct {
	OutputPath       string
	CachePath        string
	GitAuthorEmail   string
	TrackerAPIToken  string
	TrackerProjectID int
}

func configWithOutputPath(outputPath string) Config {
	projectID, _ := strconv.Atoi(os.Getenv(EnvTrackerProjectID))
	return Config{
		OutputPath:       outputPath,
		CachePath:        cachePathFromOutputPath(outputPath),
		GitAuthorEmail:   os.Getenv(EnvGitAuthorEmail),
		TrackerAPIToken:  os.Getenv(EnvTrackerAPIToken),
		TrackerProjectID: projectID,
	}
}

func cachePathFromOutputPath(outputPath string) string {
	outputAbs, _ := filepath.Abs(outputPath)
	baseDir := filepath.Dir(outputAbs)
	return filepath.Join(baseDir, ".inflight-cache")
}
