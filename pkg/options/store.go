package options

import (
	"fmt"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/solver/store"
	"github.com/spf13/cobra"
)

func GetDefaultStoreOptions() store.StoreOptions {
	return store.StoreOptions{
		Type:         GetDefaultServeOptionString("STORE_TYPE", "database"),
		ConnStr:      GetDefaultServeOptionString("STORE_CONN_STR", ""),
		GormLogLevel: GetDefaultServeOptionString("STORE_GORM_LOG_LEVEL", "silent"),
	}
}

func AddStoreCliFlags(cmd *cobra.Command, storeOptions *store.StoreOptions) {
	cmd.PersistentFlags().StringVar(
		&storeOptions.Type, "store-type", storeOptions.Type,
		`The store type used by the solver, one of "database" or "memory" (STORE_TYPE).`,
	)
	cmd.PersistentFlags().StringVar(
		&storeOptions.ConnStr, "store-conn-str", storeOptions.ConnStr,
		`The database store connection string (STORE_CONN_STR).`,
	)
	cmd.PersistentFlags().StringVar(
		&storeOptions.GormLogLevel, "store-gorm-log-level", storeOptions.GormLogLevel,
		`The database store gorm log level, one of "silent", "info", "error", "warn" (STORE_GORM_LOG_LEVEL).`,
	)
}

func CheckStoreOptions(options store.StoreOptions) error {
	if options.Type != "database" && options.Type != "memory" {
		return fmt.Errorf("STORE_TYPE must be \"database\" or \"memory\"")
	}
	if options.Type == "database" && options.ConnStr == "" {
		return fmt.Errorf("STORE_CONN_STR is required when using the database store")
	}
	if options.GormLogLevel != "silent" &&
		options.GormLogLevel != "info" &&
		options.GormLogLevel != "error" &&
		options.GormLogLevel != "warn" {
		return fmt.Errorf("STORE_GORM_LOG_LEVEL must be \"silent\", \"info\", \"error\", or \"warn\"")
	}

	return nil
}
