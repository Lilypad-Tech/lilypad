package module

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/module/shortcuts"
	"github.com/bacalhau-project/lilypad/pkg/system"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/rs/zerolog/log"
)

const REPO_DIR = "repos"

func getRepoLocalPath(repoURL string) (string, error) {
	parsedURL, err := url.Parse(repoURL)
	if err != nil {
		return "", err
	}

	pathParts := strings.Split(strings.Trim(parsedURL.Path, "/"), "/")
	if len(pathParts) < 2 {
		return "", fmt.Errorf("Invalid git URL")
	}

	return filepath.Join(REPO_DIR, pathParts[0], pathParts[1]), nil
}

func CheckModuleOptions(options data.ModuleConfig) error {
	if options.Repo == "" {
		return fmt.Errorf("MODULE_REPO is required")
	}
	if options.Hash == "" {
		return fmt.Errorf("MODULE_HASH is required")
	}
	if options.Path == "" {
		return fmt.Errorf("MODULE_PATH is required")
	}
	return nil
}

// given a module - check if it's a shortcut and if yes
// expand the shortcut into the other module props
func ProcessModule(module data.ModuleConfig) (data.ModuleConfig, error) {
	// we have been given a shortcut
	// let's try to resolve this shortcut into a full module definition
	if module.Name != "" {
		module, err := shortcuts.GetModule(module.Name, module.Version)
		if err != nil {
			return module, err
		}
		return module, nil
	}
	err := CheckModuleOptions(module)
	if err != nil {
		return module, err
	}
	return module, nil
}

// clone the given repo and return the full local path
// to the repo
func CloneModule(module data.ModuleConfig) (*git.Repository, error) {
	repoPath, err := getRepoLocalPath(module.Repo)
	if err != nil {
		return nil, err
	}
	repoDir, err := system.DataDir(repoPath)
	if err != nil {
		return nil, err
	}
	fileInfo, err := os.Stat(filepath.Join(repoDir, ".git"))
	if err == nil && fileInfo.IsDir() {
		log.Debug().
			Str("repo dir", repoDir).
			Msgf("")
		return git.PlainOpen(repoDir)
	} else if os.IsNotExist(err) {
		log.Debug().
			Str("repo dir", repoDir).
			Str("repo remote", module.Repo).
			Msgf("")
		return git.PlainClone(repoDir, false, &git.CloneOptions{
			URL: module.Repo,
		})
	} else {
		return nil, err
	}
}

// get a module cloned and checked out then return the text content of the template
//   - process shortcuts
//   - check if we have the repo cloned
//   - checkout the correct hash
//   - check and read the file
func PrepareModule(module data.ModuleConfig) (string, error) {
	module, err := ProcessModule(module)
	if err != nil {
		return "", err
	}
	repo, err := CloneModule(module)
	if err != nil {
		return "", err
	}
	worktree, err := repo.Worktree()
	if err != nil {
		return "", err
	}
	repoDir := worktree.Filesystem.Root()
	log.Debug().
		Str("checkout hash", module.Hash).
		Msgf(module.Repo)
	err = worktree.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(module.Hash),
	})
	if err != nil {
		return "", err
	}
	templatePath := filepath.Join(repoDir, module.Path)
	_, err = os.Stat(templatePath)
	if err != nil {
		return "", err
	}
	fileContents, err := os.ReadFile(templatePath)
	if err != nil {
		return "", err
	}
	log.Debug().
		Str("read file", module.Path).
		Msgf(string(fileContents))
	return string(fileContents), nil
}

// - prepare the module - now we have the text of the template
// - inject the given values using template syntax
// - JSON parse and check we don't have errors
func LoadModule(module data.ModuleConfig, inputs map[string]string) (*data.Module, error) {
	moduleText, err := PrepareModule(module)
	if err != nil {
		return nil, err
	}

	templateName := fmt.Sprintf("%s-%s-%s", module.Repo, module.Path, module.Hash)
	tmpl, err := template.New(templateName).Parse(moduleText)
	if err != nil {
		return nil, err
	}
	var template bytes.Buffer
	if err := tmpl.Execute(&template, inputs); err != nil {
		return nil, err
	}

	var moduleData data.Module
	if err := json.Unmarshal([]byte(template.String()), &moduleData); err != nil {
		return nil, err
	}

	return &moduleData, nil
}
