package options

import (
	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/module"
	"github.com/spf13/cobra"
)

func GetDefaultModuleOptions() data.ModuleConfig {
	return data.ModuleConfig{
		// the shortcut name
		Name: GetDefaultServeOptionString("MODULE_NAME", ""),
		// the repo we can clone from
		Repo: GetDefaultServeOptionString("MODULE_REPO", ""),
		// the hash to checkout the repo
		Hash: GetDefaultServeOptionString("MODULE_HASH", ""),
		// the path to the go template file
		Path: GetDefaultServeOptionString("MODULE_PATH", ""),
	}
}

func AddModuleCliFlags(cmd *cobra.Command, moduleConfig *data.ModuleConfig) {
	cmd.PersistentFlags().StringVar(
		&moduleConfig.Name, "module-name", moduleConfig.Name,
		`The name of the shortcut module (MODULE_NAME)`,
	)
	cmd.PersistentFlags().StringVar(
		&moduleConfig.Repo, "module-repo", moduleConfig.Repo,
		`The (http) git repo we can close (MODULE_REPO)`,
	)
	cmd.PersistentFlags().StringVar(
		&moduleConfig.Hash, "module-hash", moduleConfig.Hash,
		`The hash of the repo we can checkout (MODULE_HASH)`,
	)
	cmd.PersistentFlags().StringVar(
		&moduleConfig.Path, "module-path", moduleConfig.Path,
		`The path in the repo to the go template (MODULE_PATH)`,
	)
}

// see if we have a shortcut and fill in the other values if we do
func ProcessModuleOptions(options data.ModuleConfig) (data.ModuleConfig, error) {
	return module.ProcessModule(options)
}

func CheckModuleOptions(options data.ModuleConfig) error {
	return module.CheckModuleOptions(options)
}
