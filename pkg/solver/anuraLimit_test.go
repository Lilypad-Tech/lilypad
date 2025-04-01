//go:build integration && solver

package solver_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hashicorp/go-retryablehttp"
	httputil "github.com/lilypad-tech/lilypad/v2/pkg/http"
)

type anuraTestCase struct {
	name       string
	privateKey string
	expectedOK bool
}

// This test suite tests Anura authentication with valid and invalid signatures
func TestAnuraAuth(t *testing.T) {
	paths := []string{
		"/api/v1/anura/job_offers",
	}

	validKey := "b3994e7660abe5f65f729bb64163c6cd6b7d0b1a8c67881a7346e3e8c7f026f5"
	invalidKey := "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"

	t.Run("valid anura signature passes", func(t *testing.T) {
		tc := anuraTestCase{
			name:       "valid signature should pass",
			privateKey: validKey,
			expectedOK: true,
		}
		runAnuraTest(t, paths, tc)
	})

	t.Run("invalid anura signature fails", func(t *testing.T) {
		tc := anuraTestCase{
			name:       "invalid signature should fail",
			privateKey: invalidKey,
			expectedOK: false,
		}
		runAnuraTest(t, paths, tc)
	})
}

func runAnuraTest(t *testing.T, paths []string, tc anuraTestCase) {
	for _, path := range paths {
		client := retryablehttp.NewClient()
		client.RetryMax = 0
		client.Logger = nil
		client.CheckRetry = func(ctx context.Context, resp *http.Response, err error) (bool, error) {
			return false, nil
		}

		req, err := retryablehttp.NewRequest("GET", fmt.Sprintf("http://localhost:%d%s", 8081, path), nil)
		if err != nil {
			t.Errorf("Failed to create request: %s\n", err)
			return
		}

		privateKey, err := crypto.HexToECDSA(tc.privateKey)
		if err != nil {
			t.Errorf("Failed to parse private key: %s\n", err)
			return
		}

		err = httputil.AddAnuraHeaders(req, privateKey, crypto.PubkeyToAddress(privateKey.PublicKey).String())
		if err != nil {
			t.Errorf("Failed to add Anura headers: %s\n", err)
			return
		}

		res, err := client.Do(req)
		if err != nil {
			t.Errorf("Request failed on %s: %s\n", path, err)
			return
		}

		if tc.expectedOK && res.StatusCode != http.StatusOK {
			t.Errorf("%s: Expected status code 200, got %d", path, res.StatusCode)
		}
		if !tc.expectedOK && res.StatusCode != http.StatusUnauthorized {
			t.Errorf("%s: Expected status code 401, got %d", path, res.StatusCode)
		}
	}
}
