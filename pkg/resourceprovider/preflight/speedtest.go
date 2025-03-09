package preflight

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

const (
	// MinimumBandwidthMBps is the minimum required bandwidth in MB/s
	MinimumBandwidthMBps = 25

	// SpeedTestURL is the URL used for testing download speed
	// We use a reasonably sized file from a reliable CDN
	SpeedTestURL = "https://proof.ovh.net/files/100Mb.dat"

	// SpeedTestTimeout is the maximum time allowed for the speed test
	SpeedTestTimeout = 30 * time.Second

	// SpeedTestMinBytes is the minimum amount of data to download for reliable measurement
	SpeedTestMinBytes = 10 * 1024 * 1024 // 10 MB
)

type speedTestConfig struct {
	required       bool
	minBandwidthMB int
	testURL        string
}

// checkNetworkSpeed verifies if the network has sufficient bandwidth
func (p *preflightChecker) checkNetworkSpeed(ctx context.Context, config *speedTestConfig) checkResult {
	if !config.required {
		log.Info().Msg("Network speed check not required, skipping")
		return checkResult{
			passed:  true,
			message: "Network speed check skipped (not required)",
		}
	}

	log.Info().Msgf("🌐 Starting network speed test (minimum: %d MB/s)", config.minBandwidthMB)

	// Create a context with timeout
	testCtx, cancel := context.WithTimeout(ctx, SpeedTestTimeout)
	defer cancel()

	// Create HTTP request
	req, err := http.NewRequestWithContext(testCtx, http.MethodGet, config.testURL, nil)
	if err != nil {
		return checkResult{
			passed:  false,
			error:   fmt.Errorf("failed to create request: %w", err),
			message: "Failed to create network speed test request",
		}
	}

	// Start timing
	startTime := time.Now()

	// Execute the request
	client := &http.Client{
		Timeout: SpeedTestTimeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return checkResult{
			passed:  false,
			error:   fmt.Errorf("failed to execute request: %w", err),
			message: "Network connectivity issue during speed test",
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return checkResult{
			passed:  false,
			error:   fmt.Errorf("unexpected HTTP status code: %d", resp.StatusCode),
			message: fmt.Sprintf("Speed test server returned error code: %d", resp.StatusCode),
		}
	}

	// Read the response body
	var totalBytes int64
	bytesBuffer := make([]byte, 32*1024) // 32KB buffer

	for {
		if testCtx.Err() != nil {
			return checkResult{
				passed:  false,
				error:   fmt.Errorf("speed test timed out: %w", testCtx.Err()),
				message: "Network speed test timed out",
			}
		}

		n, err := resp.Body.Read(bytesBuffer)
		totalBytes += int64(n)

		// If we've read sufficient data for a reliable test, we can stop
		if totalBytes >= SpeedTestMinBytes {
			break
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			return checkResult{
				passed:  false,
				error:   fmt.Errorf("error reading response: %w", err),
				message: "Error during network speed test",
			}
		}
	}

	// Calculate elapsed time and speed
	elapsedTime := time.Since(startTime).Seconds()
	speedMBps := float64(totalBytes) / 1024 / 1024 / elapsedTime

	log.Info().
		Float64("speed_mbps", speedMBps).
		Float64("elapsed_seconds", elapsedTime).
		Int64("bytes_transferred", totalBytes).
		Msgf("🌐 Network speed test results")

	if speedMBps < float64(config.minBandwidthMB) {
		return checkResult{
			passed:  false,
			error:   fmt.Errorf("network speed below minimum requirement: %.2f MB/s < %d MB/s", speedMBps, config.minBandwidthMB),
			message: fmt.Sprintf("Network speed insufficient: %.2f MB/s (minimum: %d MB/s)", speedMBps, config.minBandwidthMB),
		}
	}

	return checkResult{
		passed:  true,
		message: fmt.Sprintf("Network speed sufficient: %.2f MB/s", speedMBps),
	}
}
