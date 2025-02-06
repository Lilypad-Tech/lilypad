//go:build integration && solver

package solver_test

import (
	"fmt"
	"net/http"
	"os"
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

	// Test non-exempt IP first (using a non-localhost address)
	t.Run("Non-exempt IP is rate limited", func(t *testing.T) {
		client := &http.Client{}
		var okCount, limitedCount int

		for i := 0; i < 10; i++ {
			req, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost:8081%s", paths[0]), nil)
			// Set X-Forwarded-For to simulate request from non-exempt IP
			req.Header.Set("X-Forwarded-For", "1.2.3.4")
			
			res, err := client.Do(req)
			if err != nil {
				t.Fatalf("Request failed: %v", err)
			}

			if res.StatusCode == 200 {
				okCount++
			} else if res.StatusCode == 429 {
				limitedCount++
			} else {
				t.Errorf("Unexpected status code: %d", res.StatusCode)
			}
			
			time.Sleep(5 * time.Millisecond)
		}

		if okCount > 5 {
			t.Errorf("Expected at most 5 successful requests, got %d", okCount)
		}
		if limitedCount == 0 {
			t.Error("Expected some requests to be rate limited")
		}
	})

	// Test exempt IP (localhost)
	t.Run("Exempt IP is not rate limited", func(t *testing.T) {
		client := &http.Client{}
		var okCount, limitedCount int

		for i := 0; i < 10; i++ {
			req, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost:8081%s", paths[0]), nil)
			
			res, err := client.Do(req)
			if err != nil {
				t.Fatalf("Request failed: %v", err)
			}

			if res.StatusCode == 200 {
				okCount++
			} else if res.StatusCode == 429 {
				limitedCount++
			} else {
				t.Errorf("Unexpected status code: %d", res.StatusCode)
			}
			
			time.Sleep(5 * time.Millisecond)
		}

		if limitedCount > 0 {
			t.Errorf("Expected no rate limiting for exempt IP, but got %d limited requests", limitedCount)
		}
		if okCount != 10 {
			t.Errorf("Expected all 10 requests to succeed, got %d successful requests", okCount)
		}
	})
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
