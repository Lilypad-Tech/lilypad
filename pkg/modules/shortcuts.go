package modules

import (
	"fmt"

	"github.com/bacalhau-project/lilypad/pkg/data"
)

func GetCowsayVersions() (string, map[string]data.Module, error) {
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

func GetModuleVersions(name string) (string, map[string]data.Module, error) {
	switch name {
	case "cowsay":
		return GetCowsayVersions()
	}
	return "", nil, fmt.Errorf("no module found for %s", name)
}

func GetModule(name string, version string) (data.Module, error) {
	if name == "" {
		return data.Module{}, fmt.Errorf("module name is empty")
	}
	latestVersion, versions, err := GetModuleVersions(name)
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
