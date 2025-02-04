package options

import (
	"fmt"

	"github.com/lilypad-tech/lilypad/pkg/solver/stats"
	"github.com/spf13/cobra"
)

func GetDefaultStatsOptions() stats.StatsOptions {
	return stats.StatsOptions{
		Enabled: GetDefaultServeOptionBool("ENABLE_STATS", false),
		URL:     GetDefaultServeOptionString("STATS_URL", ""),
	}
}

func AddStatsCliFlags(cmd *cobra.Command, statsOptions *stats.StatsOptions) {
	cmd.PersistentFlags().BoolVar(
		&statsOptions.Enabled, "enable-stats", statsOptions.Enabled,
		`Enable stats (ENABLE_STATS)`,
	)
	cmd.PersistentFlags().StringVar(
		&statsOptions.URL, "stats-url", statsOptions.URL,
		`The stats endpoint to connect to (STATS_URL).`,
	)
}

func CheckStatsOptions(options stats.StatsOptions) error {
	if options.Enabled && options.URL == "" {
		return fmt.Errorf("STATS_URL must be set when using the stats API")
	}
	return nil
}
