package options

import (
	"fmt"
	"net/url"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/solver/stats"
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
	if options.Enabled {
		if options.URL == "" {
			return fmt.Errorf("STATS_URL must be set when using the stats API")
		}
		if _, err := url.ParseRequestURI(options.URL); err != nil {
			return fmt.Errorf("Invalid STATS_URL format: %v", err)
		}
	}
	return nil
}
