//go:build integration

package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"
)

type rateResult struct {
	path         string
	okCount      int
	limitedCount int
}

// This test suite sends 100 requests over approximately half a second.
// We assume the solver uses the default rate limiting settings with
// a request limit of 5 and window length of 10 seconds.
func TestRateLimiter(t *testing.T) {
	paths := []string{
		"/api/v1/resource_offers",
		"/api/v1/job_offers",
		"/api/v1/deals",
	}

	var wg sync.WaitGroup
	ch := make(chan rateResult, len(paths))

	// Send off callers to run concurrently
	for _, path := range paths {
		wg.Add(1)

		go func() {
			defer wg.Done()
			makeCalls(t, path, ch)
		}()
	}

	wg.Wait()
	close(ch)

	expectedOkCount := 5
	for result := range ch {
		if result.okCount > expectedOkCount {
			t.Errorf(
				"%s allowed %d requests and limited %d requests, but expected limiting after %d requests\n",
				result.path, result.okCount, result.limitedCount, expectedOkCount,
			)
		}
	}
}

func makeCalls(t *testing.T, path string, ch chan rateResult) {
	var okCount int
	var limitedCount int

	// Make 100 requests
	for range 100 {
		requestURL := fmt.Sprintf("http://localhost:%d%s", 8081, path)
		res, err := http.Get(requestURL)

		if err != nil {
			t.Errorf("Get request failed on %s: %s\n", path, err)
			os.Exit(1)
		}

		if res.StatusCode == 200 {
			okCount++
		} else if res.StatusCode == 429 {
			limitedCount++
		} else {
			t.Errorf("Expected a 200 or 429 status code, but received a %d\n", res.StatusCode)
		}

		// Wait before making next call
		time.Sleep(5 * time.Millisecond)
	}

	ch <- rateResult{
		path:         path,
		okCount:      okCount,
		limitedCount: limitedCount,
	}
}
