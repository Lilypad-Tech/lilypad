//go:build unit

package module

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/lilypad-tech/lilypad/pkg/data"
)

func TestPrepareModule(t *testing.T) {
	text, err := PrepareModule(data.ModuleConfig{
		Name: "cowsay:v0.0.2",
	})

	assert.NoError(t, err, "Should not return an error")
	assert.Contains(t, text, "cowsay", "Should contain the message")
	fmt.Printf("%s\n", text)
}

func TestLoadModule(t *testing.T) {
	module, err := LoadModule(data.ModuleConfig{
		Name: "cowsay:v0.0.2",
	}, map[string]string{
		"Message": "Hello, world!",
	})

	assert.NoError(t, err, "Should not return an error")
	assert.Equal(t, "grycap/cowsay@sha256:fad516b39e3a587f33ce3dbbb1e646073ef35e0b696bcf9afb1a3e399ce2ab0b", module.Job.Spec.Docker.Image)
	assert.Equal(t, "Hello, world!", module.Job.Spec.Docker.Entrypoint[1])
}

// TestSubst: [subst] can correctly substitute json encoded values into the template string.
func TestSubst(t *testing.T) {
	format := "Hello, %s!"
	inputs := []string{"hiro"}
	expectedOutput := "Hello, hiro!"

	jsonEncodedInputs := make([]string, 0, len(inputs))

	for _, input := range inputs {
		inputJ, err := json.Marshal(input)
		if err != nil {
			t.Fatalf("json marshall failed %v", err)
		}
		jsonEncodedInputs = append(jsonEncodedInputs, string(inputJ))
	}
	t.Logf("jsonEncodedInputs -%v %d", jsonEncodedInputs, len(jsonEncodedInputs))

	actualOutput := subst(format, jsonEncodedInputs...)

	if actualOutput != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, actualOutput)
	}
}

func TestTryLock(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Test 1: Successfully acquire a lock
	t.Run("Acquire Lock", func(t *testing.T) {
		unlock, err := tryLockWithTimeout(tempDir, 1*time.Second)
		assert.NoError(t, err, "Should successfully acquire lock")
		assert.NotNil(t, unlock, "Should return unlock function")

		// Verify lock file exists
		_, err = os.Stat(filepath.Join(tempDir, ".lilypad.lock"))
		assert.NoError(t, err, "Lock file should exist")

		// Clean up
		unlock()

		// Verify lock file is removed
		_, err = os.Stat(filepath.Join(tempDir, ".lilypad.lock"))
		assert.Error(t, err, "Lock file should be removed")
	})

	// Test 2: Cannot acquire lock when already locked
	t.Run("Lock Contention", func(t *testing.T) {
		unlock1, err := tryLockWithTimeout(tempDir, 1*time.Second)
		assert.NoError(t, err, "Should successfully acquire first lock")
		defer unlock1()

		// Try to acquire second lock - should fail quickly
		unlock2, err := tryLockWithTimeout(tempDir, 1*time.Second)
		assert.Error(t, err, "Should fail to acquire second lock")
		assert.Nil(t, unlock2, "Should not return unlock function")
	})

	// Test 3: Can acquire lock after stale lock is removed
	t.Run("Stale Lock", func(t *testing.T) {
		lockFile := filepath.Join(tempDir, ".lilypad.lock")

		// Create a stale lock file with non-existent PID
		err := os.WriteFile(lockFile, []byte("999999999"), 0600)
		assert.NoError(t, err, "Should create stale lock file")

		// Should be able to acquire lock despite stale file
		unlock, err := tryLockWithTimeout(tempDir, 1*time.Second)
		assert.NoError(t, err, "Should acquire lock after stale lock")
		assert.NotNil(t, unlock, "Should return unlock function")

		// Clean up
		unlock()
	})
}
