package system

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"golang.org/x/mod/semver"
)

var Version string
var CommitSHA string

type VersionConfig struct {
	version        string
	minimumVersion string
}

func NewVersionConfig(minimumVersion string) (*VersionConfig, error) {
	if Version != "" && !semver.IsValid(Version) {
		return nil, fmt.Errorf("invalid embedded version: %s", Version)
	}

	if minimumVersion != "" && !semver.IsValid(minimumVersion) {
		return nil, fmt.Errorf("invalid minimum version: %s", minimumVersion)
	}

	if Version != "" &&
		minimumVersion != "" &&
		semver.Compare(minimumVersion, Version) > 0 {
		return nil, fmt.Errorf("minimum version %s is greater than embedded version %s", minimumVersion, Version)
	}

	return &VersionConfig{
		version:        Version,
		minimumVersion: minimumVersion,
	}, nil
}

func (vc *VersionConfig) IsSupported(version string) (string, bool) {
	var minVersion string
	if vc.minimumVersion != "" {
		minVersion = vc.minimumVersion
	} else if vc.version != "" {
		// Use embedded major version when minimum version is not set.
		// This option follows semantic versioning for breaking changes, so that the
		// minimum version can be omitted when using semantic versioning.
		minVersion = semver.Major(vc.version)
	} else {
		// Return supported when neither an embedded version nor minimum version is set.
		// This condition supports local development.
		return "", true
	}

	if !semver.IsValid(version) {
		log.Warn().Msgf("invalid version %s", version)
		return minVersion, false
	}

	result := semver.Compare(version, minVersion)
	return minVersion, result >= 0 // True if version >= minVersion
}
