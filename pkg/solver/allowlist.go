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

func loadAllowlist(filepath string) ([]AllowlistItem, error) {
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
	var formattedAllowlist []string
	for _, item := range allowlist {
		formattedAllowlist = append(formattedAllowlist, item.Module)
	}
	return strings.Join(formattedAllowlist, ", ")
}

func fetchAllowlistFromGitHub(repo, path string) ([]AllowlistItem, error) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/main/%s", repo, path)
	return loadAllowlist(url)
}

func saveAllowlistToFile(repo, path, localPath string) error {
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
	allowlist, err := fetchAllowlistFromGitHub("lilypad-tech/module-allowlist", "allowlist.json")
	if err != nil {
		fmt.Printf("Error fetching allowlist: %v\n", err)
		return
	}

	formatted := AllowlistFormatting(allowlist)
	fmt.Println("Allowlist:", formatted)
}

func main() {
	err := saveAllowlistToFile("lilypad-tech/module-allowlist", "allowlist.json", "local_allowlist.json")
	if err != nil {
		fmt.Printf("Error saving allowlist: %v\n", err)
		return
	}

	fmt.Println("Allowlist successfully downloaded to local_allowlist.json")

	allowlistPull()
}
