package shortcuts

import (
	"github.com/bacalhau-project/lilypad/pkg/data"
)

func getCowsayVersions() (string, map[string]data.ModuleConfig, error) {
	latestVersion := "v0.0.1"
	versions := map[string]data.ModuleConfig{
		"v0.0.1": {
			Repo: "https://github.com/bacalhau-project/lilypad.git",
			Hash: "d3b3cdd14874ec9bcbfdcd543afe2aa632d7427f",
			Path: "module_templates/cowsay.json",
		},
	}
	return latestVersion, versions, nil
}
