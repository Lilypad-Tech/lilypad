package shortcuts

import (
	"fmt"
	"strings"

	"github.com/bacalhau-project/lilypad/pkg/data"
)

// allow shortcode github.com/lukemarsden/lilypad-sdxl:v0.0.1 (tag),
// TODO: enforce sha1 for tags on the server side (like a pin file)

// parse something with no slashes in it as
// github.com/bacalhau-project/lilypad-module-<shortcode>

func GetModule(name string) (data.ModuleConfig, error) {
	// parse name per following valid formats
	// github.com/user/repo:tag --> Repo: https://github.com/user/repo; Hash = tag
	// bar:tag --> Repo = https://github.com/bacalhau-project/lilypad-module-<bar>, Hash = tag
	if name == "" {
		return data.ModuleConfig{}, fmt.Errorf("module name is empty")
	}
	parts := strings.Split(name, ":")
	if len(parts) != 2 {
		return data.ModuleConfig{}, fmt.Errorf("invalid module name format: %s", name)
	}
	repo, hash := parts[0], parts[1]
	if strings.Contains(name, "/") {
		repo = fmt.Sprintf("https://%s", parts[0])
	} else {
		repo = fmt.Sprintf("https://github.com/bacalhau-project/lilypad-module-%s", parts[0])
	}

	// TODO: docs for authoring a module
	module := data.ModuleConfig{
		Name: "",
		Repo: repo,
		Hash: hash,
		Path: "/lilypad_module.json.tmpl",
	}

	return module, nil
}
