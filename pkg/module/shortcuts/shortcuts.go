package shortcuts

import (
	"fmt"

	"github.com/bacalhau-project/lilypad/pkg/data"
)

func getModuleVersions(name string) (string, map[string]data.Module, error) {
	switch name {
	case "cowsay":
		return getCowsayVersions()
	}
	return "", nil, fmt.Errorf("no module found for %s", name)
}

func GetModule(name string, version string) (data.Module, error) {
	if name == "" {
		return data.Module{}, fmt.Errorf("module name is empty")
	}
	latestVersion, versions, err := getModuleVersions(name)
	if err != nil {
		return data.Module{}, err
	}
	if version == "" {
		version = latestVersion
	}
	module, ok := versions[version]
	if !ok {
		return data.Module{}, fmt.Errorf("no module version found for %s %s", name, version)
	}
	return module, nil
}
