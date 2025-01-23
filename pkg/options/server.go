package options

import (
	"fmt"

	"github.com/lilypad-tech/lilypad/pkg/http"
	"github.com/spf13/cobra"
)

func GetDefaultServerOptions() http.ServerOptions {
	return http.ServerOptions{
		URL:           GetDefaultServeOptionString("SERVER_URL", ""),
		Host:          GetDefaultServeOptionString("SERVER_HOST", "0.0.0.0"),
		Port:          GetDefaultServeOptionInt("SERVER_PORT", 8080), //nolint:gomnd
		AccessControl: GetDefaultAccessControlOptions(),
		RateLimiter:   GetDefaultRateLimiterOptions(),
	}
}

func GetDefaultAccessControlOptions() http.AccessControlOptions {
	return http.AccessControlOptions{
		ValidationTokenSecret:     GetDefaultServeOptionString("SERVER_VALIDATION_TOKEN_SECRET", ""),
		ValidationTokenExpiration: GetDefaultServeOptionInt("SERVER_VALIDATION_TOKEN_EXPIRATION", 604800), // one week
		ValidationTokenKid:        GetDefaultServeOptionString("SERVER_VALIDATION_TOKEN_KID", ""),
	}
}

func GetDefaultRateLimiterOptions() http.RateLimiterOptions {
	return http.RateLimiterOptions{
		RequestLimit: GetDefaultServeOptionInt("SERVER_RATE_REQUEST_LIMIT", 5),
		WindowLength: GetDefaultServeOptionInt("SERVER_RATE_WINDOW_LENGTH", 10),
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
}

func CheckServerOptions(options http.ServerOptions) error {
	if options.URL == "" {
		return fmt.Errorf("SERVER_URL is required")
	}
	if options.AccessControl.ValidationTokenSecret == "" {
		return fmt.Errorf("SERVER_VALIDATION_TOKEN_SECRET is required")
	}
	if options.AccessControl.ValidationTokenKid == "" {
		return fmt.Errorf("SERVER_VALIDATION_TOKEN_KID is required")
	}
	return nil
}
