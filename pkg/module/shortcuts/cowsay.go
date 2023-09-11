package shortcuts

import (
	"github.com/bacalhau-project/lilypad/pkg/data"
)

func getCowsayVersions() (string, map[string]data.ModuleConfig, error) {
	latestVersion := "v0.0.1"
	versions := map[string]data.ModuleConfig{
		"v0.0.1": {
			Repo: "https://github.com/bacalhau-project/lilypad.git",
			Hash: "2e7aae5a5ef4956fe0ee2fde7633e7966f977afc",
			Path: "module_templates/cowsay.json",
		},
	}
	return latestVersion, versions, nil
}
