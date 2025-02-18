//go:build integration && solver

package solver_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hashicorp/go-retryablehttp"
	httputil "github.com/lilypad-tech/lilypad/pkg/http"
)

type anuraRateResult struct {
	path              string
	okCount           int
	limitedCount      int
	unauthorizedCount int
}

type anuraTestCase struct {
	name           string
	privateKey     string
	expectedOK     int
	expectedLimit  int
	expectedUnauth int
}

// This test suite sends 10 requests to Anura endpoints. We test both valid and invalid
// signatures. Invalid signatures should be rate limited after 5 attempts.
func TestAnuraRateLimiter(t *testing.T) {
	paths := []string{
		"/api/v1/anura/job_offers",
	}

	validKey := "b3994e7660abe5f65f729bb64163c6cd6b7d0b1a8c67881a7346e3e8c7f026f5"
	invalidKey := "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"

	t.Run("valid anura signature passes", func(t *testing.T) {
		tc := anuraTestCase{
			name:           "valid signature should pass",
			privateKey:     validKey,
			expectedOK:     10, // All 10 requests should succeed
			expectedLimit:  0,  // No rate limiting
			expectedUnauth: 0,  // No unauthorized responses
		}
		runAnuraTest(t, paths, tc)
	})

	// Wait a bit between tests to let rate limiter reset
	time.Sleep(1 * time.Second)

	t.Run("invalid anura signature gets rate limited", func(t *testing.T) {
		tc := anuraTestCase{
			name:           "invalid signature should be rate limited",
			privateKey:     invalidKey,
			expectedOK:     0, // No requests should succeed
			expectedLimit:  5, // 5 requests should be rate limited
			expectedUnauth: 5, // First 5 should be unauthorized
		}
		runAnuraTest(t, paths, tc)
	})
}

func runAnuraTest(t *testing.T, paths []string, tc anuraTestCase) {
	for _, path := range paths {
		result := makeAnuraCalls(t, path, tc)
		
		if result.okCount != tc.expectedOK {
			t.Errorf("%s: Expected %d successful requests, got %d",
				result.path, tc.expectedOK, result.okCount)
		}
		if result.limitedCount != tc.expectedLimit {
			t.Errorf("%s: Expected %d rate limited requests, got %d",
				result.path, tc.expectedLimit, result.limitedCount)
		}
		if result.unauthorizedCount != tc.expectedUnauth {
			t.Errorf("%s: Expected %d unauthorized requests, got %d",
				result.path, tc.expectedUnauth, result.unauthorizedCount)
		}
	}
}

func makeAnuraCalls(t *testing.T, path string, tc anuraTestCase) anuraRateResult {
	var okCount, limitedCount, unauthorizedCount int

	// Configure client with no retries and silent logger
	client := retryablehttp.NewClient()
	client.RetryMax = 0
	client.Logger = nil // Disable logging
	client.CheckRetry = func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		return false, nil // Never retry
	}

	for i := 0; i < 10; i++ {
		req, err := retryablehttp.NewRequest("GET", fmt.Sprintf("http://localhost:%d%s", 8081, path), nil)
		if err != nil {
			t.Errorf("Failed to create request: %s\n", err)
			return anuraRateResult{}
		}

		privateKey, err := crypto.HexToECDSA(tc.privateKey)
		if err != nil {
			t.Errorf("Failed to parse private key: %s\n", err)
			return anuraRateResult{}
		}

		err = httputil.AddAnuraHeaders(req, privateKey, crypto.PubkeyToAddress(privateKey.PublicKey).String())
		if err != nil {
			t.Errorf("Failed to add Anura headers: %s\n", err)
			return anuraRateResult{}
		}

		res, err := client.Do(req)
		if err != nil {
			t.Errorf("Request failed on %s: %s\n", path, err)
			return anuraRateResult{}
		}

		switch res.StatusCode {
		case http.StatusOK:
			okCount++
		case http.StatusTooManyRequests:
			limitedCount++
		case http.StatusUnauthorized:
			unauthorizedCount++
		default:
			t.Errorf("Expected 200, 401, or 429 status code, but received %d\n", res.StatusCode)
		}

		time.Sleep(100 * time.Millisecond)
	}

	return anuraRateResult{
		path:              path,
		okCount:           okCount,
		limitedCount:      limitedCount,
		unauthorizedCount: unauthorizedCount,
	}
}
