package shortcuts

import (
	"github.com/bacalhau-project/lilypad/pkg/data"
)

func getCowsayVersions() (string, map[string]data.ModuleConfig, error) {
	latestVersion := "v0.0.1"
	versions := map[string]data.ModuleConfig{
		"v0.0.1": {
			Repo: "https://github.com/bacalhau-project/lilypad.git",
			Hash: "0817bd741283761aa898eaa3c50775a31f279e2b",
			Path: "module_templates/cowsay.json",
		},
	}
	return latestVersion, versions, nil
}
