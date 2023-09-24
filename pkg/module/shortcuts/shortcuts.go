package shortcuts

import (
	"fmt"

	"github.com/bacalhau-project/lilypad/pkg/data"
)

// TODO: allow shortcode github.com/lukemarsden/lilypad-sdxl:v0.0.1 (tag), enforce sha1 for tags on the server side (like a pin file)
// github.com/lp-mods/lilypad-sdxl:v0.0.1

// parse something with no slashes in it as github.com/bacalhau-project/lilypad-default-module-<shortcode>

func getModuleVersions(name string) (string, map[string]data.ModuleConfig, error) {
	switch name {
	case "cowsay":
		return getCowsayVersions()
	}
	return "", nil, fmt.Errorf("no module found for %s", name)
}

func GetModule(name string, version string) (data.ModuleConfig, error) {
	if name == "" {
		return data.ModuleConfig{}, fmt.Errorf("module name is empty")
	}

	// // If there are no slashes in the module name, default to
	// // github.com/bacalhau-project/lilypad-module-<name>
	// if !strings.Contains(name, "/") {
	// 	name = fmt.Sprintf("github.com/bacalhau-project/lilypad-module-%s", name)
	// }

	latestVersion, versions, err := getModuleVersions(name)
	if err != nil {
		return data.ModuleConfig{}, err
	}
	if version == "" {
		version = latestVersion
	}
	module, ok := versions[version]
	if !ok {
		return data.ModuleConfig{}, fmt.Errorf("no module version found for %s %s", name, version)
	}
	return module, nil
}
