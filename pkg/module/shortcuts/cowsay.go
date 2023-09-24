package shortcuts

import (
	"github.com/bacalhau-project/lilypad/pkg/data"
)

func getCowsayVersions() (string, map[string]data.ModuleConfig, error) {
	latestVersion := "v0.0.1"
	versions := map[string]data.ModuleConfig{
		"v0.0.1": {
			Repo: "https://github.com/bacalhau-project/lilypad.git",
			Hash: "7ffcdefbf30b27edfcc1375bf2b302a820d3f6c0",
			Path: "module_templates/cowsay.json",
		},
	}
	return latestVersion, versions, nil
}
