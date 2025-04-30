package options

import (
	"fmt"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/http"
	"github.com/spf13/cobra"
)

func GetDefaultServerOptions() http.ServerOptions {
	return http.ServerOptions{
		URL:           GetDefaultServeOptionString("SERVER_URL", ""),
		Host:          GetDefaultServeOptionString("SERVER_HOST", "0.0.0.0"),
		Port:          GetDefaultServeOptionInt("SERVER_PORT", 8080), //nolint:gomnd
		AccessControl: GetDefaultAccessControlOptions(),
		RateLimiter:   GetDefaultRateLimiterOptions(),
		Storage:       GetDefaultStorageOptions(),
	}
}

func GetDefaultAccessControlOptions() http.AccessControlOptions {
	return http.AccessControlOptions{
		// When false, any resource provider may post a resource offer.
		EnableResourceProviderAllowlist: GetDefaultServeOptionBool("SERVER_ENABLE_RESOURCE_PROVIDER_ALLOWLIST", false),
		ValidationTokenSecret:           GetDefaultServeOptionString("SERVER_VALIDATION_TOKEN_SECRET", ""),
		ValidationTokenExpiration:       GetDefaultServeOptionInt("SERVER_VALIDATION_TOKEN_EXPIRATION", 604800), // one week
		ValidationTokenKid:              GetDefaultServeOptionString("SERVER_VALIDATION_TOKEN_KID", ""),
		AnuraAddresses:                  GetDefaultServeOptionStringArray("SERVER_ANURA_ADDRESSES", []string{}),
		OfferTimestampDiffSeconds:       GetDefaultServeOptionInt("SERVER_OFFER_TIMESTAMP_DIFF_SECONDS", 30),
		EnableVersionCheck:              GetDefaultServeOptionBool("SERVER_ENABLE_VERSION_CHECK", true),
		MinimumVersion:                  GetDefaultServeOptionString("SERVER_MINIMUM_VERSION", ""),
	}
}

func GetDefaultRateLimiterOptions() http.RateLimiterOptions {
	return http.RateLimiterOptions{
		RequestLimit: GetDefaultServeOptionInt("SERVER_RATE_REQUEST_LIMIT", 5),
		WindowLength: GetDefaultServeOptionInt("SERVER_RATE_WINDOW_LENGTH", 10),
	}
}

func GetDefaultStorageOptions() http.StorageOptions {
	return http.StorageOptions{
		MaximumFileInputsMemoryMB: GetDefaultServeOptionInt("SERVER_MAX_FILE_INPUTS_MEMORY_MB", 20),
		MaximumFileInputsSizeMB:   GetDefaultServeOptionInt("SERVER_MAX_FILE_INPUTS_SIZE_MB", 10),
	}
}

func AddServerCliFlags(cmd *cobra.Command, serverOptions *http.ServerOptions) {
	cmd.PersistentFlags().StringVar(
		&serverOptions.URL, "server-url", serverOptions.URL,
		`The URL the api server is listening on (SERVER_URL).`,
	)
	cmd.PersistentFlags().StringVar(
		&serverOptions.Host, "server-host", serverOptions.Host,
		`The host to bind the api server to (SERVER_HOST).`,
	)
	cmd.PersistentFlags().IntVar(
		&serverOptions.Port, "server-port", serverOptions.Port,
		`The port to bind the api server to (SERVER_PORT).`,
	)
	cmd.PersistentFlags().BoolVar(
		&serverOptions.AccessControl.EnableResourceProviderAllowlist, "server-enable-resource-provider-allowlist",
		serverOptions.AccessControl.EnableResourceProviderAllowlist,
		`Enable resource provider allowlist (SERVER_ENABLE_RESOURCE_PROVIDER_ALLOWLIST).`,
	)
	cmd.PersistentFlags().StringVar(
		&serverOptions.AccessControl.ValidationTokenSecret, "server-validation-token-secret",
		serverOptions.AccessControl.ValidationTokenSecret,
		`Secret for generating validation service JWTs (SERVER_VALIDATION_TOKEN_SECRET).`,
	)
	cmd.PersistentFlags().IntVar(
		&serverOptions.AccessControl.ValidationTokenExpiration, "server-validation-token-expiration",
		serverOptions.AccessControl.ValidationTokenExpiration,
		`Validation service JWT expiration in seconds (SERVER_VALIDATION_TOKEN_EXPIRATION).`,
	)
	cmd.PersistentFlags().StringVar(
		&serverOptions.AccessControl.ValidationTokenKid, "server-validation-token-kid",
		serverOptions.AccessControl.ValidationTokenKid,
		`Key ID header for validation service JWTs (SERVER_VALIDATION_TOKEN_KID).`,
	)
	cmd.PersistentFlags().IntVar(
		&serverOptions.RateLimiter.RequestLimit, "server-rate-request-limit", serverOptions.RateLimiter.RequestLimit,
		`The max requests over the rate window length (SERVER_RATE_REQUEST_LIMIT).`,
	)
	cmd.PersistentFlags().IntVar(
		&serverOptions.RateLimiter.WindowLength, "server-rate-window-length", serverOptions.RateLimiter.WindowLength,
		`The time window over which to limit in seconds (SERVER_RATE_WINDOW_LENGTH).`,
	)
	cmd.PersistentFlags().StringArrayVar(
		&serverOptions.AccessControl.AnuraAddresses, "server-anura-addresses", serverOptions.AccessControl.AnuraAddresses,
		`The Anura wallet addresses that are allowed to access anura endpoints (SERVER_ANURA_ADDRESSES).`,
	)
	cmd.PersistentFlags().IntVar(
		&serverOptions.AccessControl.OfferTimestampDiffSeconds, "server-offer-timestamp-diff-seconds", serverOptions.AccessControl.OfferTimestampDiffSeconds,
		`The diff before or after now when a job or resource offer must be received (SERVER_OFFER_TIMESTAMP_DIFF_SECONDS).`,
	)
	cmd.PersistentFlags().BoolVar(
		&serverOptions.AccessControl.EnableVersionCheck, "server-enable-version-check", serverOptions.AccessControl.EnableVersionCheck,
		`Enable version check (SERVER_ENABLE_VERSION_CHECK).`,
	)
	cmd.PersistentFlags().StringVar(
		&serverOptions.AccessControl.MinimumVersion, "server-minimum-version", serverOptions.AccessControl.MinimumVersion,
		`The minimum client version (SERVER_MINIMUM_VERSION).`,
	)
	cmd.PersistentFlags().IntVar(
		&serverOptions.Storage.MaximumFileInputsMemoryMB, "server-max-file-inputs-memory-mb", serverOptions.Storage.MaximumFileInputsMemoryMB,
		`The maxiumum memory used when accepting file inputs to a job (SERVER_MAX_FILE_INPUTS_MEMORY_MB).`,
	)
	cmd.PersistentFlags().IntVar(
		&serverOptions.Storage.MaximumFileInputsSizeMB, "server-max-file-inputs-size-mb", serverOptions.Storage.MaximumFileInputsSizeMB,
		`The maxiumum file inputs upload size to a job (SERVER_MAX_FILE_INPUTS_SIZE_MB).`,
	)
}

func CheckServerOptions(options http.ServerOptions, storeType string) error {
	if options.URL == "" {
		return fmt.Errorf("SERVER_URL is required")
	}
	if options.AccessControl.EnableResourceProviderAllowlist && storeType == "memory" {
		return fmt.Errorf("Enabling the resource provider allowlist requires the database store. Set STORE_TYPE to \"database\".")
	}
	if options.AccessControl.ValidationTokenSecret == "" {
		return fmt.Errorf("SERVER_VALIDATION_TOKEN_SECRET is required")
	}
	if options.AccessControl.ValidationTokenKid == "" {
		return fmt.Errorf("SERVER_VALIDATION_TOKEN_KID is required")
	}
	if options.AccessControl.OfferTimestampDiffSeconds < 1 {
		return fmt.Errorf("SERVER_OFFER_TIMESTAMP_DIFF_SECONDS must be greater than zero")
	}
	return nil
}
