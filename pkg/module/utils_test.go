//go:build unit

package module

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/data"
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

func TestTryLockWithTimeout(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Test 1: Successfully acquire a lock
	t.Run("Acquire Lock", func(t *testing.T) {
		unlock, err := tryLockWithTimeout(tempDir, 1*time.Second)
		assert.NoError(t, err, "Should successfully acquire lock")
		assert.NotNil(t, unlock, "Should return unlock function")

		// Verify lock file exists
		lockFile := filepath.Join(tempDir, ".lilypad.lock")
		_, err = os.Stat(lockFile)
		assert.NoError(t, err, "Lock file should exist")

		// Verify lock file contains expected content
		content, err := os.ReadFile(lockFile)
		assert.NoError(t, err, "Should be able to read lock file")

		var pid int
		var hostname, timestamp string
		_, err = fmt.Sscanf(string(content), "%d %s %s", &pid, &hostname, &timestamp)
		assert.NoError(t, err, "Should be able to parse lock file content")
		assert.Equal(t, os.Getpid(), pid, "PID should match current process")

		// Clean up
		unlock()

		// Verify lock file is removed
		_, err = os.Stat(lockFile)
		assert.Error(t, err, "Lock file should be removed")
		assert.True(t, os.IsNotExist(err), "Lock file should not exist")
	})

	// Test 2: Cannot acquire lock when already locked
	t.Run("Lock Contention", func(t *testing.T) {
		unlock1, err := tryLockWithTimeout(tempDir, 1*time.Second)
		assert.NoError(t, err, "Should successfully acquire first lock")
		defer unlock1()

		// Try to acquire second lock - should fail quickly
		unlock2, err := tryLockWithTimeout(tempDir, 500*time.Millisecond) // Use shorter timeout
		assert.Error(t, err, "Should fail to acquire second lock")
		assert.Nil(t, unlock2, "Should not return unlock function")
		assert.Contains(t, err.Error(), "timeout", "Error should mention timeout")
	})

	// Test 3: Can acquire lock after stale lock is removed
	t.Run("Stale Lock", func(t *testing.T) {
		lockFile := filepath.Join(tempDir, ".lilypad.lock")

		// Create a stale lock file with non-existent PID
		content := fmt.Sprintf("%d %s %s", 999999999, "test-host", time.Now().Format(time.RFC3339))
		err := os.WriteFile(lockFile, []byte(content), 0600)
		assert.NoError(t, err, "Should create stale lock file")

		// Should be able to acquire lock despite stale file
		unlock, err := tryLockWithTimeout(tempDir, 1*time.Second)
		assert.NoError(t, err, "Should acquire lock after stale lock")
		assert.NotNil(t, unlock, "Should return unlock function")

		// Clean up
		unlock()
	})

	// Test 4: Can acquire lock after old lock is removed (timeout based)
	t.Run("Old Lock", func(t *testing.T) {
		lockFile := filepath.Join(tempDir, ".lilypad.lock")

		// Create an old lock file with valid PID but old timestamp
		oldTime := time.Now().Add(-11 * time.Minute).Format(time.RFC3339)
		content := fmt.Sprintf("%d %s %s", os.Getpid(), "test-host", oldTime)
		err := os.WriteFile(lockFile, []byte(content), 0600)
		assert.NoError(t, err, "Should create old lock file")

		// Should be able to acquire lock despite old file
		unlock, err := tryLockWithTimeout(tempDir, 1*time.Second)
		assert.NoError(t, err, "Should acquire lock after old lock")
		assert.NotNil(t, unlock, "Should return unlock function")

		// Clean up
		unlock()
	})

	// Test 5: Concurrent locking attempts should work correctly
	t.Run("Concurrent Locking", func(t *testing.T) {
		// Create multiple directories to avoid test interference
		concurrentDir := filepath.Join(t.TempDir(), "concurrent")
		err := os.MkdirAll(concurrentDir, 0755)
		assert.NoError(t, err, "Should create concurrent test directory")

		// Number of concurrent goroutines trying to acquire the lock
		numGoroutines := 5

		// This will track which goroutine got the lock
		gotLock := make(chan int, numGoroutines)
		var wg sync.WaitGroup

		// Launch goroutines to try to get the lock
		for i := range numGoroutines {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				unlock, err := tryLockWithTimeout(concurrentDir, 2*time.Second)
				if err == nil && unlock != nil {
					// Got the lock!
					gotLock <- id
					// Hold the lock briefly
					time.Sleep(100 * time.Millisecond)
					unlock()
				}
			}(i)
		}

		// Collect the results
		lockHolders := make([]int, 0, numGoroutines)
		go func() {
			wg.Wait()
			close(gotLock)
		}()

		for id := range gotLock {
			lockHolders = append(lockHolders, id)
		}

		assert.True(t, len(lockHolders) > 0, "At least one goroutine should get the lock")
		// In an ideal world, all 5 would get it sequentially, but we can't guarantee the timing
		// so we just check that someone did
	})

	// Test 6: Can acquire lock when timestamp format is invalid
	t.Run("Invalid Timestamp", func(t *testing.T) {
		lockFile := filepath.Join(tempDir, ".lilypad.lock")

		// Create a lock file with invalid timestamp format
		invalidTimestamp := "not-a-timestamp"
		content := fmt.Sprintf("%d %s %s", os.Getpid(), "test-host", invalidTimestamp)
		err := os.WriteFile(lockFile, []byte(content), 0600)
		assert.NoError(t, err, "Should create lock file with invalid timestamp")

		// Verify the malformed lock file exists
		_, err = os.Stat(lockFile)
		assert.NoError(t, err, "Lock file should exist")

		// Should be able to acquire lock despite invalid timestamp format
		unlock, err := tryLockWithTimeout(tempDir, 1*time.Second)
		assert.NoError(t, err, "Should acquire lock after detecting invalid timestamp")
		assert.NotNil(t, unlock, "Should return unlock function")

		// Clean up
		unlock()
	})
}
