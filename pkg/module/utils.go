package module

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/rs/zerolog/log"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/module/shortcuts"
	"github.com/bacalhau-project/lilypad/pkg/system"
)

const REPO_DIR = "repos"

func getRepoLocalPath(repoURL string) (string, error) {
	parsedURL, err := url.Parse(repoURL)
	if err != nil {
		return "", fmt.Errorf("url parsing failed with %v", err)
	}

	pathParts := strings.Split(strings.Trim(parsedURL.Path, "/"), "/")
	if len(pathParts) < 2 {
		return "", fmt.Errorf("invalid git URL")
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
		module, err := shortcuts.GetModule(module.Name)
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

// clone the given repo and return the full local path to the repo
// TODO: check if we have the repo already cloned
// handle fetching new changes perhaps the commit hash is not the latest
// at the moment we will do the slow thing and clone the repo every time
func CloneModule(module data.ModuleConfig) (repo *git.Repository, err error) {
	repoPath, err := getRepoLocalPath(module.Repo)
	if err != nil {
		return nil, err
	}
	repoDir, err := system.EnsureDataDir(repoPath)
	if err != nil {
		return nil, err
	}
	fileInfo, err := os.Stat(filepath.Join(repoDir, ".git"))

	repoCloned := err == nil && fileInfo.IsDir()

	if !repoCloned {
		log.Debug().
			Str("repo clone", repoDir).
			Str("repo remote", module.Repo).
			Msgf("")
		return git.PlainClone(repoDir, false, &git.CloneOptions{
			URL:      module.Repo,
			Progress: os.Stdout,
		})
	}

	log.Debug().
		Str("repo exists", repoDir).
		Str("repo remote", module.Repo).
		Msgf("")

	repo, err = git.PlainOpen(repoDir)
	// err := os.RemoveAll(repoDir)
	if err != nil {
		return nil, err
	}

	// git fetch origin: Resolves #https://github.com/bacalhau-project/lilypad/issues/13
	gitFetchOptions := &git.FetchOptions{
		Tags:     git.AllTags,
		Progress: os.Stdout,
	}
	gitFetchOptions.Validate() // sets default values like remote=origin
	log.Info().Str("updating cached git repo", repoDir).Msgf("")
	err = repo.FetchContext(context.Background(), gitFetchOptions)

	// Check if hash or tag specified exists
	h, err := repo.ResolveRevision(plumbing.Revision(module.Hash))
	if err != nil {
		return nil, err
	}
	// XXX SECURITY: on RP side, need to verify this hash is in the allowlist
	// explicitly to ensure determinism (and that we're running the code we
	// explicitly approved)
	_, err = repo.Storer.EncodedObject(plumbing.AnyObject, *h)
	if err != nil {
		// this means there is no hash in the repo
		// so let's clean it up and clone it again
		err = os.RemoveAll(repoDir)
		if err != nil {
			return nil, err
		}
	}
	return
}

// PrepareModule get a module cloned and checked out then return the text content of the template
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
	// TODO: match tags against hash specified in server side allowlist
	repoDir := worktree.Filesystem.Root()
	log.Debug().
		Str("checkout hash", module.Hash).
		Msgf(module.Repo)

	h, err := repo.ResolveRevision(plumbing.Revision(module.Hash))
	if err != nil {
		return "", err
	}
	// XXX SECURITY: on RP side, need to verify this hash is in the allowlist
	// explicitly to ensure determinism (and that we're running the code we
	// explicitly approved)
	err = worktree.Checkout(&git.CheckoutOptions{
		Hash: *h,
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

func subst(format string, jsonEncodedInputs ...string) string {

	jsonDecodedInputs := make([]any, 0, len(jsonEncodedInputs))

	for _, input := range jsonEncodedInputs {
		var s string

		if err := json.Unmarshal([]byte(input), &s); err != nil {
			log.Debug().AnErr("subst: json unmarshall", err).Msgf("input:%s", input)
			panic("subst: invalid input")
		}

		jsonDecodedInputs = append(jsonDecodedInputs, s)
	}
	log.Printf("jsonDecodedInputs:%v", jsonDecodedInputs)

	return fmt.Sprintf(format, jsonDecodedInputs...)
}

// - prepare the module - now we have the text of the template
// - inject the given values using template syntax
// - JSON parse and check we don't have errors
func LoadModule(module data.ModuleConfig, inputs map[string]string) (*data.Module, error) {
	moduleText, err := PrepareModule(module)
	if err != nil {
		return nil, err
	}
	// TODO: golang handlebars implementation, with shortcode for string encoding e.g. escape_string

	templateName := fmt.Sprintf("%s-%s-%s", module.Repo, module.Path, module.Hash)
	tmpl, err := template.New(templateName).Parse(moduleText)
	tmpl.Funcs(template.FuncMap{
		"subst": subst,
	})
	if err != nil {
		return nil, err
	}

	newInputs := make(map[string]string)
	// For now, for each input, json encode it so that it's safe to put into the template
	for k, v := range inputs {
		bs, err := json.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal string %q", v)
		}
		newInputs[k] = string(bs)
	}

	var template bytes.Buffer
	if err := tmpl.Execute(&template, newInputs); err != nil {
		return nil, fmt.Errorf(
			"error executing template: %s (tmpl=%s, inputs=%+v)",
			err,
			moduleText,
			newInputs,
		)
	}

	var moduleData data.Module
	bs := template.Bytes()
	if err := json.Unmarshal(bs, &moduleData); err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling resulting json: %s, %s",
			err,
			bs,
		)
	}

	return &moduleData, nil
}
