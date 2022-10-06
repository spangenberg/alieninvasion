package version

import "fmt"

var (
	BuildDate  string
	GitBranch  string
	GitCommit  string
	GitState   string
	GitSummary string
	Version    string
)

func String() string {
	if GitState == "clean" {
		return fmt.Sprintf("%s\nbuild date %s", Version, BuildDate)
	}
	return fmt.Sprintf("%s (%s)\nbuild date %s", Version, GitSummary, BuildDate)
}
