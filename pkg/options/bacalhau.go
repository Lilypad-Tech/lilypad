package options

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lilypad-tech/lilypad/pkg/executor/bacalhau"
	"github.com/spf13/cobra"
)

func GetDefaultBacalhauOptions() bacalhau.BacalhauExecutorOptions {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("error getting home directory: %s", err.Error()))
	}

	resultsPath := filepath.Join(home, ".bacalhau", "compute", "results", "local-publisher")

	return bacalhau.BacalhauExecutorOptions{
		ApiHost:               GetDefaultServeOptionString("BACALHAU_API_HOST", "localhost"),
		ApiPort:               GetDefaultServeOptionString("BACALHAU_API_PORT", "1234"),
		JobStatusPollInterval: GetDefaultServeOptionUint64("JOB_STATUS_POLL_INTERVAL", 5),
		ResultsDirectory:      GetDefaultServeOptionString("BACALHAU_RESULTS_DIRECTORY", resultsPath),
	}
}

func AddBacalhauCliFlags(cmd *cobra.Command, bacalhauOptions *bacalhau.BacalhauExecutorOptions) {
	cmd.PersistentFlags().StringVar(
		&bacalhauOptions.ApiHost, "bacalhau-api-host", bacalhauOptions.ApiHost,
		`The api hostname for the bacalhau cluster to run jobs`,
	)

	cmd.PersistentFlags().StringVar(
		&bacalhauOptions.ApiPort, "bacalhau-api-port", bacalhauOptions.ApiPort,
		`The api port for the bacalhau cluster to run jobs`,
	)
}

func CheckBacalhauOptions(options bacalhau.BacalhauExecutorOptions) error {
	if options.ApiHost == "" {
		return fmt.Errorf("No bacalhau service specified - please use BACALHAU_API_HOST or --bacalhau-api-host")
	}
	return nil
}
