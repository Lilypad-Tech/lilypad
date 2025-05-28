package options

import (
	"github.com/Lilypad-Tech/lilypad/v2/pkg/solver"
	"github.com/spf13/cobra"
)

func GetDefaultTimeoutOptions() solver.SolverTimeoutOptions {
	return solver.SolverTimeoutOptions{
		MatchSeconds:     GetDefaultServeOptionUint64("TIMEOUT_MATCH_SECONDS", 30),
		ExecutionSeconds: GetDefaultServeOptionUint64("TIMEOUT_EXECUTION_SECONDS", 900), // 15 minutes
		DownloadSeconds:  GetDefaultServeOptionUint64("TIMEOUT_DOWNLOAD_SECONDS", 3600), // One hour
		TotalSeconds:     GetDefaultServeOptionUint64("TIMEOUT_TOTAL_SECONDS", 86400),   // One day
	}
}

func AddTimeoutCliFlags(cmd *cobra.Command, timeoutConfig *solver.SolverTimeoutOptions) {
	cmd.PersistentFlags().Uint64Var(
		&timeoutConfig.MatchSeconds, "timeout-match-seconds", timeoutConfig.MatchSeconds,
		`The number of seconds to match before timing out (TIMEOUT_MATCH_SECONDS)`,
	)
	cmd.PersistentFlags().Uint64Var(
		&timeoutConfig.ExecutionSeconds, "timeout-execution-seconds", timeoutConfig.ExecutionSeconds,
		`The number of seconds for a resource provider to execute before timing out (TIMEOUT_EXECUTION_SECONDS)`,
	)
	cmd.PersistentFlags().Uint64Var(
		&timeoutConfig.DownloadSeconds, "timeout-download-seconds", timeoutConfig.DownloadSeconds,
		`The number of seconds for a job creator to download job outputs before timing out (TIMEOUT_DOWNLOAD_SECONDS)`,
	)
	cmd.PersistentFlags().Uint64Var(
		&timeoutConfig.TotalSeconds, "timeout-total-seconds", timeoutConfig.TotalSeconds,
		`The total number of seconds before a job times out (TIMEOUT_TOTAL_SECONDS)`,
	)
}
