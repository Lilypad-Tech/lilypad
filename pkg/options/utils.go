package options

import (
	"os"
	"strconv"
	"strings"
)

func GetDefaultServeOptionString(envName string, defaultValue string) string {
	envValue := os.Getenv(envName)
	if envValue != "" {
		return envValue
	}
	return defaultValue
}

func GetDefaultServeOptionUint64(envName string, defaultValue uint64) uint64 {
	envValue := os.Getenv(envName)
	if envValue != "" {
		// convert envValue to int
		i, err := strconv.Atoi(envValue)
		if err == nil {
			return uint64(i)
		}
		return 0
	}
	return defaultValue
}

func GetDefaultServeOptionStringArray(envName string, defaultValue []string) []string {
	envValue := os.Getenv(envName)
	if envValue != "" {
		return strings.Split(envValue, ",")
	}
	return defaultValue
}

func GetDefaultServeOptionInt(envName string, defaultValue int) int {
	envValue := os.Getenv(envName)
	if envValue != "" {
		i, err := strconv.Atoi(envValue)
		if err == nil {
			return i
		}
	}
	return defaultValue
}
