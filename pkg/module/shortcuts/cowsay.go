package shortcuts

import (
	"github.com/bacalhau-project/lilypad/pkg/data"
)

func getCowsayVersions() (string, map[string]data.Module, error) {
	latestVersion := "v0.0.1"
	versions := map[string]data.Module{
		"v0.0.1": {
			Repo: "https://github.com/bacalhau-project/lilypad.git",
			Hash: "abe5eb1033e2c3cadf723960e501a0cbe29d3116",
			Path: "modules/cowsay.json",
		},
	}
	return latestVersion, versions, nil
}
