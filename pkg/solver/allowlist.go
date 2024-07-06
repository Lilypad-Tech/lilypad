package solver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// AllowlistItem represents an allowlist item
type AllowlistItem struct {
	Module string `json:"module"`
}

var allowlistCheckerEnabled bool

func EnableAllowlistChecker() {
	allowlistCheckerEnabled = true
}

func DisableAllowlistChecker() {
	allowlistCheckerEnabled = false
}

func loadAllowlist(filepath string) ([]AllowlistItem, error) {
	if !allowlistCheckerEnabled {
		return nil, fmt.Errorf("allowlist checker is disabled")
	}

	var bytes []byte
	var err error

	if strings.HasPrefix(filepath, "http://") || strings.HasPrefix(filepath, "https://") {
		resp, err := http.Get(filepath)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to fetch allowlist from URL: %s", resp.Status)
		}

		bytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err = ioutil.ReadFile(filepath)
		if err != nil {
			return nil, err
		}
	}

	var allowlist []AllowlistItem
	if err := json.Unmarshal(bytes, &allowlist); err != nil {
		return nil, err
	}

	return allowlist, nil
}

func AllowlistFormatting(allowlist []AllowlistItem) string {
	if !allowlistCheckerEnabled {
		return "allowlist checker is disabled"
	}

	var formattedAllowlist []string
	for _, item := range allowlist {
		formattedAllowlist = append(formattedAllowlist, item.Module)
	}
	return strings.Join(formattedAllowlist, ", ")
}

func fetchAllowlistFromGitHub(repo, path string) ([]AllowlistItem, error) {
	if !allowlistCheckerEnabled {
		return nil, fmt.Errorf("allowlist checker is disabled")
	}

	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/main/%s", repo, path)
	return loadAllowlist(url)
}

func saveAllowlistToFile(repo, path, localPath string) error {
	if !allowlistCheckerEnabled {
		return fmt.Errorf("allowlist checker is disabled")
	}

	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/main/%s", repo, path)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch allowlist: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch allowlist from URL: %s", resp.Status)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	err = ioutil.WriteFile(localPath, bytes, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func allowlistPull() {
	if !allowlistCheckerEnabled {
		fmt.Println("allowlist checker is disabled")
		return
	}

	allowlist, err := fetchAllowlistFromGitHub("lilypad-tech/module-allowlist", "allowlist.json")
	if err != nil {
		fmt.Printf("Error fetching allowlist: %v\n", err)
		return
	}

	formatted := AllowlistFormatting(allowlist)
	fmt.Println("Allowlist:", formatted)
}
