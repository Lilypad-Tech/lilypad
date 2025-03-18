//go:build unit

package system

import (
	"testing"
)

func TestIsSupported(t *testing.T) {
	originalVersion := Version

	tests := []struct {
		name                string
		embeddedVersion     string
		minimumVersion      string
		versionToCheck      string
		expectedMinVersion  string
		expectedIsSupported bool
	}{
		{
			name:                "No embedded version nor minimum version returns true",
			embeddedVersion:     "",
			minimumVersion:      "",
			versionToCheck:      "v0.1.0",
			expectedMinVersion:  "",
			expectedIsSupported: true,
		},
		{
			name:                "No embedded version but has minimum version - newer version supported",
			embeddedVersion:     "",
			minimumVersion:      "v1.2.3",
			versionToCheck:      "v1.3.0",
			expectedMinVersion:  "v1.2.3",
			expectedIsSupported: true,
		},
		{
			name:                "No embedded version but has minimum version - older version not supported",
			embeddedVersion:     "",
			minimumVersion:      "v1.2.3",
			versionToCheck:      "v1.1.0",
			expectedMinVersion:  "v1.2.3",
			expectedIsSupported: false,
		},
		{
			name:                "No minimum version uses embedded major version - newer version supported",
			embeddedVersion:     "v2.3.4",
			minimumVersion:      "",
			versionToCheck:      "v2.0.0",
			expectedMinVersion:  "v2",
			expectedIsSupported: true,
		},
		{
			name:                "No minimum version uses embedded major version - older version not supported",
			embeddedVersion:     "v2.3.4",
			minimumVersion:      "",
			versionToCheck:      "v1.9.9",
			expectedMinVersion:  "v2",
			expectedIsSupported: false,
		},
		{
			name:                "Minimum version used when specified - newer version supported",
			embeddedVersion:     "v3.4.5",
			minimumVersion:      "v2.0.0",
			versionToCheck:      "v2.1.0",
			expectedMinVersion:  "v2.0.0",
			expectedIsSupported: true,
		},
		{
			name:                "Minimum version used when specified - exact match supported",
			embeddedVersion:     "v3.4.5",
			minimumVersion:      "v2.0.0",
			versionToCheck:      "v2.0.0",
			expectedMinVersion:  "v2.0.0",
			expectedIsSupported: true,
		},
		{
			name:                "Minimum version used when specified - older version not supported",
			embeddedVersion:     "v3.4.5",
			minimumVersion:      "v2.0.0",
			versionToCheck:      "v1.9.9",
			expectedMinVersion:  "v2.0.0",
			expectedIsSupported: false,
		},
		{
			name:                "Invalid version to check returns false",
			embeddedVersion:     "v3.4.5",
			minimumVersion:      "v2.0.0",
			versionToCheck:      "invalid",
			expectedMinVersion:  "v2.0.0",
			expectedIsSupported: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			Version = tc.embeddedVersion

			vc, err := NewVersionConfig(tc.minimumVersion)
			if err != nil {
				t.Fatalf("Failed to create VersionConfig: %v", err)
			}

			minVersion, isSupported := vc.IsSupported(tc.versionToCheck)

			if minVersion != tc.expectedMinVersion {
				t.Errorf("Expected min version %q, got %q", tc.expectedMinVersion, minVersion)
			}

			if isSupported != tc.expectedIsSupported {
				t.Errorf("Expected isSupported to be %v, got %v", tc.expectedIsSupported, isSupported)
			}
		})
	}

	Version = originalVersion
}

func TestNewVersionConfig_Validation(t *testing.T) {
	originalVersion := Version

	tests := []struct {
		name            string
		embeddedVersion string
		minimumVersion  string
		expectError     bool
	}{
		{
			name:            "Valid versions",
			embeddedVersion: "v1.2.3",
			minimumVersion:  "v1.0.0",
			expectError:     false,
		},
		{
			name:            "Invalid embedded version",
			embeddedVersion: "1.2.3", // missing v prefix
			minimumVersion:  "v1.0.0",
			expectError:     true,
		},
		{
			name:            "Invalid minimum version",
			embeddedVersion: "v1.2.3",
			minimumVersion:  "1.0.0", // missing v prefix
			expectError:     true,
		},
		{
			name:            "Minimum version greater than embedded version",
			embeddedVersion: "v1.2.3",
			minimumVersion:  "v2.0.0",
			expectError:     true,
		},
		{
			name:            "Empty embedded version is valid",
			embeddedVersion: "",
			minimumVersion:  "v1.0.0",
			expectError:     false,
		},
		{
			name:            "Empty minimum version is valid",
			embeddedVersion: "v1.2.3",
			minimumVersion:  "",
			expectError:     false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			Version = tc.embeddedVersion

			_, err := NewVersionConfig(tc.minimumVersion)
			if (err != nil) != tc.expectError {
				t.Errorf("Expected error: %v, got error: %v", tc.expectError, err)
			}
		})
	}

	Version = originalVersion
}
