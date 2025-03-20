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
	"sync"
	"syscall"
	"text/template"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/rs/zerolog/log"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/module/shortcuts"
	"github.com/lilypad-tech/lilypad/pkg/system"
)

const REPO_DIR = "repos"

type RefCache struct {
	cache map[string]map[string]string
	mu    sync.RWMutex
}

func NewRefCache() *RefCache {
	return &RefCache{
		cache: make(map[string]map[string]string),
	}
}

func (rc *RefCache) Get(repoURL, reference string) string {
	rc.mu.RLock()
	defer rc.mu.RUnlock()

	if refs, ok := rc.cache[repoURL]; ok {
		return refs[reference]
	}
	return ""
}

func (rc *RefCache) Set(repoURL, reference, commitHash string) {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	if _, ok := rc.cache[repoURL]; !ok {
		rc.cache[repoURL] = make(map[string]string)
	}
	rc.cache[repoURL][reference] = commitHash
}

var globalRefCache = NewRefCache()

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

// tryLock attempts to create a lock file for the given repository path
// returns a cleanup function that should be called to release the lock
func tryLock(repoPath string) (func(), error) {
	return tryLockWithTimeout(repoPath, 30*time.Second)
}

// tryLockWithTimeout is like tryLock but with a configurable timeout
func tryLockWithTimeout(repoPath string, timeout time.Duration) (func(), error) {
	lockFile := filepath.Join(repoPath, ".lilypad.lock")

	// Try to acquire lock for the specified duration
	timeoutChan := time.After(timeout)
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-timeoutChan:
			return nil, fmt.Errorf("timeout waiting for repository lock")
		case <-ticker.C:
			// Try to create lock file
			f, err := os.OpenFile(lockFile, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
			if err == nil {
				// Write PID to lock file
				fmt.Fprintf(f, "%d", os.Getpid())
				f.Close()

				// Return cleanup function
				return func() {
					os.Remove(lockFile)
				}, nil
			}

			// If file exists, check if the process holding the lock is still alive
			if os.IsExist(err) {
				// Read PID from lock file
				pidBytes, err := os.ReadFile(lockFile)
				if err == nil {
					var pid int
					_, err = fmt.Sscanf(string(pidBytes), "%d", &pid)
					if err == nil {
						// Check if process exists
						process, err := os.FindProcess(pid)
						if err == nil {
							// On Unix, this always succeeds, so we need to send signal 0 to test
							err = process.Signal(syscall.Signal(0))
							if err != nil {
								// Process is not running, remove stale lock
								os.Remove(lockFile)
							}
						}
					}
				}
			}
		}
	}
}

func CloneModule(module data.ModuleConfig) (*git.Repository, error) {
	maxRetries := 3
	for retry := 0; retry < maxRetries; retry++ {
		commitHash := globalRefCache.Get(module.Repo, module.Hash)

		if commitHash == "" {
			var err error
			commitHash, err = resolveToCommitHash(module.Repo, module.Hash)
			if err != nil {
				log.Debug().
					Str("repo", module.Repo).
					Str("reference", module.Hash).
					Err(err).
					Msg("Failed to resolve reference")
				continue
			}

			globalRefCache.Set(module.Repo, module.Hash, commitHash)

			log.Debug().
				Str("repo", module.Repo).
				Str("reference", module.Hash).
				Str("commitHash", commitHash).
				Msg("Resolved reference to commit hash")
		}

		repoPath, err := getHashBasedRepoPath(module.Repo, commitHash)
		if err != nil {
			return nil, fmt.Errorf("failed to get hash-based repo path: %w", err)
		}

		repoDir, err := system.EnsureDataDir(repoPath)
		if err != nil {
			return nil, fmt.Errorf("failed to create repo directory: %w", err)
		}

		gitDirPath := filepath.Join(repoDir, ".git")
		fileInfo, err := os.Stat(gitDirPath)
		repoExists := (err == nil && fileInfo.IsDir())

		if repoExists {
			log.Debug().
				Str("path", repoDir).
				Str("commitHash", commitHash).
				Msg("Using cached repository")

			// For existing repos, just open and verify the hash
			repo, err := git.PlainOpen(repoDir)
			if err != nil {
				log.Error().
					Str("repoDir", repoDir).
					Err(err).
					Msg("Cached repo corrupt - will re-clone")
				_ = os.RemoveAll(repoDir)
				continue
			}

			// Verify we have the correct hash
			head, err := repo.Head()
			if err != nil {
				log.Error().
					Str("repoDir", repoDir).
					Err(err).
					Msg("Cached repo corrupt - will re-clone")
				_ = os.RemoveAll(repoDir)
				continue
			}

			if head.Hash().String() != commitHash {
				log.Error().
					Str("repoDir", repoDir).
					Str("expectedHash", commitHash).
					Str("actualHash", head.Hash().String()).
					Msg("Cached repo has wrong hash - will re-clone")
				_ = os.RemoveAll(repoDir)
				continue
			}

			return repo, nil
		}

		// Only acquire lock when we need to clone
		unlock, err := tryLock(repoDir)
		if err != nil {
			return nil, fmt.Errorf("failed to acquire repository lock: %w", err)
		}

		log.Debug().
			Str("path", repoDir).
			Str("remote", module.Repo).
			Str("commitHash", commitHash).
			Msg("Cloning repository with specific hash")

		// Double check if another process created the repo while we were waiting for the lock
		if fileInfo, err := os.Stat(gitDirPath); err == nil && fileInfo.IsDir() {
			unlock()
			continue // Retry from the start to use the existing repo
		}

		if removeErr := os.RemoveAll(repoDir); removeErr != nil {
			unlock()
			return nil, fmt.Errorf("failed to remove corrupt repo: %w", removeErr)
		}

		// Clone the repository
		repo, err := git.PlainClone(repoDir, false, &git.CloneOptions{
			URL:      module.Repo,
			Progress: os.Stdout,
		})
		if err != nil {
			unlock()
			log.Warn().
				Str("repo", module.Repo).
				Int("retry", retry+1).
				Msg("Re-cloning repository")
			continue
		}

		// Resolve and checkout the specific commit
		h, err := repo.ResolveRevision(plumbing.Revision(commitHash))
		if err != nil {
			log.Debug().
				Str("commitHash", commitHash).
				Msg("Hash not found in initial clone, fetching all refs")

			fetchOpts := &git.FetchOptions{
				Force:    true,
				Tags:     git.AllTags,
				Progress: os.Stdout,
			}

			if fetchErr := repo.FetchContext(context.Background(), fetchOpts); fetchErr != nil &&
				fetchErr != git.NoErrAlreadyUpToDate {
				unlock()
				return nil, fmt.Errorf("failed to fetch refs: %w", fetchErr)
			}

			h, err = repo.ResolveRevision(plumbing.Revision(commitHash))
			if err != nil {
				unlock()
				return nil, fmt.Errorf("hash %s not found even after fetch: %w", commitHash, err)
			}
		}

		w, err := repo.Worktree()
		if err != nil {
			unlock()
			return nil, fmt.Errorf("failed to get worktree: %w", err)
		}

		err = w.Checkout(&git.CheckoutOptions{
			Hash: *h,
		})
		if err != nil {
			unlock()
			return nil, fmt.Errorf("failed to checkout hash %s: %w", commitHash, err)
		}

		// Verify the checkout worked
		head, err := repo.Head()
		if err != nil {
			unlock()
			return nil, fmt.Errorf("failed to get HEAD after checkout: %w", err)
		}

		if head.Hash().String() != h.String() {
			unlock()
			return nil, fmt.Errorf("checkout verification failed: HEAD is %s, expected %s",
				head.Hash().String(), h.String())
		}

		// Successfully cloned and checked out
		unlock()
		return repo, nil
	}

	return nil, fmt.Errorf("failed to clone repo after %d attempts", maxRetries)
}

// PrepareModule get a module cloned and checked out then return the text content of the template
//   - process shortcuts
//   - check if we have the repo cloned
//   - checkout the correct hash
//   - check and read the file
func PrepareModule(module data.ModuleConfig) (string, error) {
	originalRef := module.Hash

	module, err := ProcessModule(module)
	if err != nil {
		return "", err
	}

	if originalRef != module.Hash {
		log.Debug().
			Str("originalRef", originalRef).
			Str("resolvedHash", module.Hash).
			Msg("Resolved reference to hash")
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
		Str("hash", module.Hash).
		Str("repo", module.Repo).
		Msg("Checking out repository at specific hash")

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
		Str("path", module.Path).
		Int("size", len(fileContents)).
		Msg("Read module template file")
	return string(fileContents), nil
}

func subst(format string, jsonEncodedInputs ...string) string {
	jsonDecodedInputs := make([]interface{}, 0, len(jsonEncodedInputs))

	for _, input := range jsonEncodedInputs {
		var s string

		if err := json.Unmarshal([]byte(input), &s); err != nil {
			log.Debug().AnErr("subst: json unmarshall", err).Msgf("input:%s", input)
			panic("subst: invalid input")
		}

		jsonDecodedInputs = append(jsonDecodedInputs, s)
	}
	log.Debug().Interface("jsonDecodedInputs", jsonDecodedInputs).Msg("Template substitution inputs")

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
	tmpl := template.New(templateName).Funcs(template.FuncMap{
		"subst": subst,
	})
	tmpl, err = tmpl.Parse(moduleText)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	newInputs := make(map[string]string)
	// For now, for each input, json encode it so that it's safe to put into the template
	for k, v := range inputs {
		bs, err := json.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal string %q: %w", v, err)
		}
		newInputs[k] = string(bs)
	}

	var template bytes.Buffer
	if err := tmpl.Execute(&template, newInputs); err != nil {
		return nil, fmt.Errorf(
			"error executing template: %w (inputs=%+v)",
			err,
			newInputs,
		)
	}

	var moduleData data.Module
	bs := template.Bytes()
	if err := json.Unmarshal(bs, &moduleData); err != nil {
		return nil, fmt.Errorf(
			"error unmarshalling resulting JSON: %w, content: %s",
			err,
			bs,
		)
	}

	return &moduleData, nil
}

func resolveToCommitHash(repoURL, reference string) (string, error) {
	tmpDir, err := os.MkdirTemp("", "git-resolver-*")
	if err != nil {
		return "", fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	log.Debug().
		Str("repo", repoURL).
		Str("reference", reference).
		Msg("Resolving git reference to commit hash")

	repo, err := git.PlainClone(tmpDir, true, &git.CloneOptions{ // true = bare repo
		URL:          repoURL,
		Progress:     os.Stdout,
		NoCheckout:   true,
		Depth:        1,
		SingleBranch: false,
		Tags:         git.AllTags,
	})
	if err != nil {
		return "", fmt.Errorf("failed to clone for reference resolution: %w", err)
	}

	hash, err := repo.ResolveRevision(plumbing.Revision(reference))
	if err != nil {
		log.Debug().
			Str("reference", reference).
			Msg("Reference not found in minimal clone, fetching all refs")

		fetchOpts := &git.FetchOptions{
			Force:    true,
			Tags:     git.AllTags,
			Progress: os.Stdout,
			Depth:    1,
		}

		if fetchErr := repo.FetchContext(context.Background(), fetchOpts); fetchErr != nil &&
			fetchErr != git.NoErrAlreadyUpToDate {
			return "", fmt.Errorf("failed to fetch refs: %w", fetchErr)
		}

		hash, err = repo.ResolveRevision(plumbing.Revision(reference))
		if err != nil {
			return "", fmt.Errorf("reference %s not found: %w", reference, err)
		}
	}

	return hash.String(), nil
}

func getHashBasedRepoPath(repoURL, hash string) (string, error) {
	if repoURL == "" {
		return "", fmt.Errorf("empty repository URL")
	}
	if hash == "" {
		return "", fmt.Errorf("empty commit hash")
	}

	parsedURL, err := url.Parse(repoURL)
	if err != nil {
		return "", fmt.Errorf("url parsing failed with %v", err)
	}

	pathParts := strings.Split(strings.Trim(parsedURL.Path, "/"), "/")
	if len(pathParts) < 2 {
		return "", fmt.Errorf("invalid git URL")
	}

	host := parsedURL.Host
	owner := pathParts[0]
	repo := pathParts[1]

	return filepath.Join(REPO_DIR, host, owner, repo, hash), nil
}
