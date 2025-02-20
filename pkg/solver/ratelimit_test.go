//go:build integration && solver

package solver_test

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"testing"
	"time"
)

type rateResult struct {
	path         string
	okCount      int
	limitedCount int
}

type rateTestCase struct {
	name          string
	headers       map[string]string
	expectedOK    int
	expectedLimit int
}

// This test suite sends 200 requests to three different paths. We send the
// requests in rate limited groups. The rate limited group should allow 5/100
// requests through.
//
// We assume the solver uses the default rate limiting settings with
// a request limit of 5 and window length of 10 seconds.
func TestRateLimiter(t *testing.T) {
	paths := []string{
		"/api/v1/resource_offers",
		"/api/v1/job_offers",
		"/api/v1/deals",
	}

	// The solver should rate limit when forwarded
	// headers are set to 1.2.3.4.
	forwardedHeaders := []map[string]string{
		{"True-Client-IP": "1.2.3.4"},
		{"X-Real-IP": "1.2.3.4"},
		{"X-Forwarded-For": "1.2.3.4"},
	}

	t.Run("requests are rate limited", func(t *testing.T) {
		// Select a random header on each test run. Over time we test them all.
		headers := forwardedHeaders[rand.Intn(len(forwardedHeaders))]
		tc := rateTestCase{
			name:          fmt.Sprintf("rate limited with headers %v", headers),
			headers:       headers,
			expectedOK:    5,
			expectedLimit: 95,
		}
		runRateLimitTest(t, paths, tc)
	})
}

func runRateLimitTest(t *testing.T, paths []string, tc rateTestCase) {
	var wg sync.WaitGroup
	ch := make(chan rateResult, len(paths))

	for _, path := range paths {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			makeCalls(t, path, ch, tc)
		}(path)
	}

	wg.Wait()
	close(ch)

	for result := range ch {
		if result.okCount != tc.expectedOK {
			t.Errorf("%s: Expected %d successful requests, got %d",
				result.path, tc.expectedOK, result.okCount)
		}
		if result.limitedCount != tc.expectedLimit {
			t.Errorf("%s: Expected %d rate limited requests, got %d",
				result.path, tc.expectedLimit, result.limitedCount)
		}
	}
}

func makeCalls(t *testing.T, path string, ch chan rateResult, tc rateTestCase) {
	var okCount int
	var limitedCount int
	client := &http.Client{}

	for i := 0; i < 100; i++ {
		req, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost:%d%s", 8081, path), nil)

		// Set test case headers
		for key, value := range tc.headers {
			req.Header.Set(key, value)
		}

		res, err := client.Do(req)
		if err != nil {
			t.Errorf("Request failed on %s: %s\n", path, err)
			return
		}

		if res.StatusCode == 200 {
			okCount++
		} else if res.StatusCode == 429 {
			limitedCount++
		} else {
			t.Errorf("Expected a 200 or 429 status code, but received a %d\n", res.StatusCode)
		}

		time.Sleep(5 * time.Millisecond)
	}

	ch <- rateResult{
		path:         path,
		okCount:      okCount,
		limitedCount: limitedCount,
	}
}
