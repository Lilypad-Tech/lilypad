package module

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"text/template"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/rs/zerolog/log"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/data"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/module/shortcuts"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/system"
)

const REPO_DIR = "repos"

// Helper function to get integer value from environment variable or return default
func getDefaultOptionInt(envName string, defaultValue int) int {
	envValue := os.Getenv(envName)
	if envValue != "" {
		i, err := strconv.Atoi(envValue)
		if err == nil {
			return i
		}
	}
	return defaultValue
}

// ModuleTimeoutConfig holds all timeout configuration values
type ModuleTimeoutConfig struct {
	// Lock related timeouts
	DefaultLockTimeout  time.Duration // Default timeout for acquiring lock
	ExtendedLockTimeout time.Duration // Extended timeout for lock acquisition during clone
	LockTickInterval    time.Duration // Interval between lock acquisition attempts
	StaleLockThreshold  time.Duration // Time after which a lock is considered stale

	// Git operation related timeouts and retries
	MaxRetries             int           // Maximum number of retries for git operations
	RetryDelayMultiplier   time.Duration // Base delay multiplier for retries
	NetworkRetryMultiplier time.Duration // Additional delay for network-related errors
}

// GetDefaultModuleTimeoutOptions returns default options from environment variables or hardcoded defaults
func GetDefaultModuleTimeoutOptions() ModuleTimeoutConfig {
	return ModuleTimeoutConfig{
		DefaultLockTimeout:     time.Duration(getDefaultOptionInt("MODULE_DEFAULT_LOCK_TIMEOUT", 30)) * time.Second,
		ExtendedLockTimeout:    time.Duration(getDefaultOptionInt("MODULE_EXTENDED_LOCK_TIMEOUT", 60)) * time.Second,
		LockTickInterval:       time.Duration(getDefaultOptionInt("MODULE_LOCK_TICK_INTERVAL", 100)) * time.Millisecond,
		StaleLockThreshold:     time.Duration(getDefaultOptionInt("MODULE_STALE_LOCK_THRESHOLD", 600)) * time.Second, // 10 minutes
		MaxRetries:             getDefaultOptionInt("MODULE_MAX_RETRIES", 5),
		RetryDelayMultiplier:   time.Duration(getDefaultOptionInt("MODULE_RETRY_DELAY_MULTIPLIER", 2)) * time.Second,
		NetworkRetryMultiplier: time.Duration(getDefaultOptionInt("MODULE_NETWORK_RETRY_MULTIPLIER", 10)) * time.Second,
	}
}

// DefaultTimeoutConfig provides default values for all timeouts
var DefaultTimeoutConfig = GetDefaultModuleTimeoutOptions()

// Current timeout configuration, can be modified at runtime
var TimeoutConfig = DefaultTimeoutConfig

// ConfigureTimeouts allows users to modify the timeout configuration
// Partial configurations are supported - only the provided fields will be updated
func ConfigureTimeouts(config ModuleTimeoutConfig) {
	// Only update non-zero values
	if config.DefaultLockTimeout > 0 {
		TimeoutConfig.DefaultLockTimeout = config.DefaultLockTimeout
	}

	if config.ExtendedLockTimeout > 0 {
		TimeoutConfig.ExtendedLockTimeout = config.ExtendedLockTimeout
	}

	if config.LockTickInterval > 0 {
		TimeoutConfig.LockTickInterval = config.LockTickInterval
	}

	if config.StaleLockThreshold > 0 {
		TimeoutConfig.StaleLockThreshold = config.StaleLockThreshold
	}

	if config.MaxRetries > 0 {
		TimeoutConfig.MaxRetries = config.MaxRetries
	}

	if config.RetryDelayMultiplier > 0 {
		TimeoutConfig.RetryDelayMultiplier = config.RetryDelayMultiplier
	}

	if config.NetworkRetryMultiplier > 0 {
		TimeoutConfig.NetworkRetryMultiplier = config.NetworkRetryMultiplier
	}

	log.Debug().
		Interface("timeoutConfig", TimeoutConfig).
		Msg("Updated module timeout configuration")
}

// ResetTimeouts resets all timeouts to default values
func ResetTimeouts() {
	TimeoutConfig = DefaultTimeoutConfig
	log.Debug().Msg("Reset module timeout configuration to defaults")
}

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

func tryLockWithTimeout(repoPath string, timeout time.Duration) (func(), error) {
	// Ensure the directory exists before trying to create a lock file
	if err := os.MkdirAll(repoPath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory for lock: %w", err)
	}

	lockFile := filepath.Join(repoPath, ".lilypad.lock")

	log.Debug().Str("path", lockFile).Msg("Attempting to acquire lock")

	// Try to acquire lock for the specified duration
	timeoutChan := time.After(timeout)
	ticker := time.NewTicker(TimeoutConfig.LockTickInterval)
	defer ticker.Stop()

	for {
		select {
		case <-timeoutChan:
			log.Warn().Str("path", lockFile).Msg("Timeout waiting for repository lock")
			return nil, fmt.Errorf("timeout waiting for repository lock")
		case <-ticker.C:
			// Try to create lock file
			f, err := os.OpenFile(lockFile, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
			if err == nil {
				// Write PID and hostname to lock file for better debugging
				hostname, _ := os.Hostname()
				fmt.Fprintf(f, "%d %s %s", os.Getpid(), hostname, time.Now().Format(time.RFC3339))
				f.Close()

				log.Debug().Str("path", lockFile).Msg("Lock acquired")

				// Return cleanup function
				return func() {
					log.Debug().Str("path", lockFile).Msg("Releasing lock")
					os.Remove(lockFile)
				}, nil
			}

			// If file exists, check if the process holding the lock is still alive
			if os.IsExist(err) {
				// Read lock file content
				content, readErr := os.ReadFile(lockFile)
				if readErr != nil {
					log.Warn().Str("path", lockFile).Err(readErr).Msg("Could not read lock file")
					continue
				}

				var pid int
				var hostname, timestamp string

				_, scanErr := fmt.Sscanf(string(content), "%d %s %s", &pid, &hostname, &timestamp)
				if scanErr != nil {
					log.Warn().
						Str("path", lockFile).
						Str("content", string(content)).
						Err(scanErr).
						Msg("Could not parse lock file, removing stale lock")
					os.Remove(lockFile)
					continue
				}

				// Check if the lock is too old (> 10 minutes)
				lockTime, parseErr := time.Parse(time.RFC3339, timestamp)
				if parseErr != nil {
					log.Warn().
						Str("path", lockFile).
						Str("timestamp", timestamp).
						Err(parseErr).
						Msg("Could not parse timestamp in lock file, removing stale lock")
					os.Remove(lockFile)
					continue
				}

				if time.Since(lockTime) > TimeoutConfig.StaleLockThreshold {
					log.Warn().
						Str("path", lockFile).
						Str("created", timestamp).
						Msg("Lock file is too old, removing")
					os.Remove(lockFile)
					continue
				}

				// Check if process exists
				process, err := os.FindProcess(pid)
				if err != nil {
					// On some systems, FindProcess never fails
					log.Debug().Int("pid", pid).Msg("Could not find process, removing lock")
					os.Remove(lockFile)
					continue
				}

				// On Unix, this always succeeds, so we need to send signal 0 to test
				err = process.Signal(syscall.Signal(0))
				if err != nil {
					log.Debug().
						Int("pid", pid).
						Str("hostname", hostname).
						Msg("Process not running, removing stale lock")
					os.Remove(lockFile)
				} else if hostname != "" {
					currentHostname, _ := os.Hostname()
					if hostname != currentHostname {
						log.Debug().
							Str("lockHostname", hostname).
							Str("currentHostname", currentHostname).
							Msg("Lock from different host")
					}
				}
			} else {
				log.Debug().Err(err).Msg("Unexpected error trying to create lock file")
			}
		}
	}
}

func CloneModule(module data.ModuleConfig) (*git.Repository, error) {
	maxRetries := TimeoutConfig.MaxRetries
	var lastErr error

	for retry := range maxRetries {
		// If this is a retry, add a small delay to avoid hammering the system
		if retry > 0 {
			delay := TimeoutConfig.RetryDelayMultiplier * time.Duration(retry)
			log.Debug().
				Str("repo", module.Repo).
				Int("retry", retry).
				Dur("delay", delay).
				Msg("Waiting before retry attempt")
			time.Sleep(delay)
		}

		commitHash := globalRefCache.Get(module.Repo, module.Hash)

		if commitHash == "" {
			var err error
			commitHash, err = resolveToCommitHash(module.Repo, module.Hash)
			if err != nil {
				lastErr = err
				log.Warn().
					Str("repo", module.Repo).
					Str("reference", module.Hash).
					Err(err).
					Int("retry", retry+1).
					Msg("Failed to resolve reference, will retry")

				// For network errors, wait longer before retrying
				if strings.Contains(err.Error(), "EOF") ||
					strings.Contains(err.Error(), "dial") ||
					strings.Contains(err.Error(), "connection") ||
					strings.Contains(err.Error(), "timeout") {
					time.Sleep(TimeoutConfig.NetworkRetryMultiplier * time.Duration(retry+1))
				}
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

		// Check if repo already exists and is valid before trying to lock
		gitDirPath := filepath.Join(repoDir, ".git")
		fileInfo, err := os.Stat(gitDirPath)
		if err == nil && fileInfo.IsDir() {
			// Try to open the repo without locking
			repo, err := git.PlainOpen(repoDir)
			if err == nil {
				// Verify we have the correct hash
				head, err := repo.Head()
				if err == nil && head.Hash().String() == commitHash {
					log.Debug().
						Str("path", repoDir).
						Str("commitHash", commitHash).
						Msg("Using cached repository")
					return repo, nil
				} else {
					log.Warn().
						Str("path", repoDir).
						Str("expectedHash", commitHash).
						Str("actualHash", func() string {
							if head != nil && err == nil {
								return head.Hash().String()
							}
							return "unknown"
						}()).
						Msg("Repository exists but has wrong hash, will rebuild")
				}
			} else {
				log.Warn().
					Str("path", repoDir).
					Err(err).
					Msg("Error opening existing repository, will rebuild")
			}
		}

		// Get an exclusive lock for rebuilding this repo
		// Use a longer timeout for the lock acquisition
		unlock, err := tryLockWithTimeout(repoDir, TimeoutConfig.ExtendedLockTimeout)
		if err != nil {
			lastErr = err
			log.Warn().
				Str("path", repoDir).
				Err(err).
				Msg("Failed to acquire lock, will retry")
			continue
		}

		// We'll use this to ensure lock is always released
		var repoResult *git.Repository
		var resultErr error

		func() {
			defer unlock() // Ensure lock is released even if we panic

			// After we've acquired the lock, check again if the repo exists and is valid
			// as another process might have created it while we were waiting
			fileInfo, err := os.Stat(gitDirPath)
			if err == nil && fileInfo.IsDir() {
				repo, err := git.PlainOpen(repoDir)
				if err == nil {
					head, err := repo.Head()
					if err == nil && head.Hash().String() == commitHash {
						log.Debug().
							Str("path", repoDir).
							Str("commitHash", commitHash).
							Msg("Repository created while waiting for lock")
						repoResult = repo
						return
					}
				}
				// If repo exists but is invalid, we'll need to clean it
				log.Warn().
					Str("path", repoDir).
					Msg("Removing corrupt or invalid repository")
			}

			// Clean directory before attempting to clone
			// First, ensure the parent directory exists
			parentDir := filepath.Dir(repoDir)
			if err := os.MkdirAll(parentDir, 0755); err != nil {
				resultErr = fmt.Errorf("failed to create parent directory: %w", err)
				return
			}

			// Create a temporary directory for the clone in the same parent directory
			tmpDir, err := os.MkdirTemp(parentDir, "git-clone-*")
			if err != nil {
				resultErr = fmt.Errorf("failed to create temporary directory: %w", err)
				return
			}

			// Ensure cleanup of the temp directory on any exit path
			defer func() {
				if tmpDir != "" {
					_ = os.RemoveAll(tmpDir)
				}
			}()

			log.Debug().
				Str("tmpDir", tmpDir).
				Str("remote", module.Repo).
				Str("commitHash", commitHash).
				Msg("Cloning repository to temporary directory")

			// Clone the repository into the temporary directory
			repo, err := git.PlainClone(tmpDir, false, &git.CloneOptions{
				URL:      module.Repo,
				Progress: os.Stdout,
			})
			if err != nil {
				resultErr = fmt.Errorf("failed to clone repository: %w", err)
				log.Warn().
					Str("repo", module.Repo).
					Int("retry", retry+1).
					Err(err).
					Msg("Error cloning repository")
				return
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
					resultErr = fmt.Errorf("failed to fetch refs: %w", fetchErr)
					return
				}

				h, err = repo.ResolveRevision(plumbing.Revision(commitHash))
				if err != nil {
					resultErr = fmt.Errorf("hash %s not found even after fetch: %w", commitHash, err)
					return
				}
			}

			w, err := repo.Worktree()
			if err != nil {
				resultErr = fmt.Errorf("failed to get worktree: %w", err)
				return
			}

			err = w.Checkout(&git.CheckoutOptions{
				Hash: *h,
			})
			if err != nil {
				resultErr = fmt.Errorf("failed to checkout hash %s: %w", commitHash, err)
				return
			}

			// Verify the checkout worked
			head, err := repo.Head()
			if err != nil {
				resultErr = fmt.Errorf("failed to get HEAD after checkout: %w", err)
				return
			}

			if head.Hash().String() != h.String() {
				resultErr = fmt.Errorf("checkout verification failed: HEAD is %s, expected %s",
					head.Hash().String(), h.String())
				return
			}

			// Successfully cloned and checked out in the temporary directory
			// Now we can safely remove the target directory if it exists
			if err := os.RemoveAll(repoDir); err != nil {
				resultErr = fmt.Errorf("failed to remove existing directory: %w", err)
				return
			}

			// Ensure the target parent directory exists
			if err := os.MkdirAll(filepath.Dir(repoDir), 0755); err != nil {
				resultErr = fmt.Errorf("failed to create parent directory: %w", err)
				return
			}

			// Rename the temporary directory to the target directory
			// This is an atomic operation on most filesystems, which helps prevent partial clones
			if err := os.Rename(tmpDir, repoDir); err != nil {
				resultErr = fmt.Errorf("failed to move repository from temp location: %w", err)
				return
			}

			// Successfully moved, don't delete the temp directory
			tmpDir = ""

			// Open the repository at its final location
			finalRepo, err := git.PlainOpen(repoDir)
			if err != nil {
				resultErr = fmt.Errorf("failed to open repository after move: %w", err)
				return
			}

			log.Info().
				Str("path", repoDir).
				Str("commitHash", commitHash).
				Msg("Repository successfully cloned and checked out")
			repoResult = finalRepo
		}()

		// If we got a repository, return it
		if repoResult != nil {
			return repoResult, nil
		}

		// Track the last error
		if resultErr != nil {
			lastErr = resultErr
		}

		// If we hit a fatal error that won't be resolved by retrying, return it immediately
		if resultErr != nil && (strings.Contains(resultErr.Error(), "authentication") ||
			strings.Contains(resultErr.Error(), "not found")) {
			return nil, resultErr
		}

		log.Warn().
			Str("repo", module.Repo).
			Int("retry", retry+1).
			Err(resultErr).
			Msg("Repository operation failed, will retry")
	}

	if lastErr != nil {
		return nil, fmt.Errorf("failed to clone repo after %d attempts: %w", maxRetries, lastErr)
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
		Str("tmpDir", tmpDir).
		Msg("Resolving git reference to commit hash")

	// Clone with retries
	var repo *git.Repository
	var cloneErr error

	for retry := range TimeoutConfig.MaxRetries {
		if retry > 0 {
			log.Debug().
				Str("repo", repoURL).
				Int("retry", retry).
				Msg("Retrying reference resolution")
			time.Sleep(TimeoutConfig.RetryDelayMultiplier * time.Duration(retry))
		}

		repo, cloneErr = git.PlainClone(tmpDir, true, &git.CloneOptions{ // true = bare repo
			URL:          repoURL,
			Progress:     os.Stdout,
			NoCheckout:   true,
			SingleBranch: false,
			Tags:         git.AllTags,
		})

		if cloneErr == nil {
			break
		}

		// If it's an auth error or not found, don't retry
		if strings.Contains(cloneErr.Error(), "authentication") ||
			strings.Contains(cloneErr.Error(), "not found") {
			return "", fmt.Errorf("failed to clone for reference resolution: %w", cloneErr)
		}
	}

	if cloneErr != nil {
		return "", fmt.Errorf("failed to clone for reference resolution after retries: %w", cloneErr)
	}

	hash, err := repo.ResolveRevision(plumbing.Revision(reference))
	if err != nil {
		log.Debug().
			Str("reference", reference).
			Err(err).
			Msg("Reference not found in minimal clone, fetching all refs")

		// Try to fetch with retries
		var fetchErr error
		for retry := range TimeoutConfig.MaxRetries {
			if retry > 0 {
				time.Sleep(TimeoutConfig.RetryDelayMultiplier * time.Duration(retry))
			}

			fetchOpts := &git.FetchOptions{
				Force:    true,
				Tags:     git.AllTags,
				Progress: os.Stdout,
			}

			fetchErr = repo.FetchContext(context.Background(), fetchOpts)
			if fetchErr == nil || fetchErr == git.NoErrAlreadyUpToDate {
				break
			}
		}

		if fetchErr != nil && fetchErr != git.NoErrAlreadyUpToDate {
			return "", fmt.Errorf("failed to fetch refs: %w", fetchErr)
		}

		hash, err = repo.ResolveRevision(plumbing.Revision(reference))
		if err != nil {
			return "", fmt.Errorf("reference %s not found: %w", reference, err)
		}
	}

	hashStr := hash.String()
	log.Debug().
		Str("repo", repoURL).
		Str("reference", reference).
		Str("hash", hashStr).
		Msg("Successfully resolved reference to commit hash")

	return hashStr, nil
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

// File inputs

func HasInputFiles(inputFiles data.InputFiles) bool {
	return len(inputFiles.Required) > 0 || len(inputFiles.Optional) > 0
}

// Check that required files exist and only explicitly defined
// files are in the target directory
func ValidateInputFiles(path string, inputFiles data.InputFiles) error {
	allowedFiles := make(map[string]bool)
	for _, file := range inputFiles.Required {
		allowedFiles[file] = true
	}
	for _, file := range inputFiles.Optional {
		allowedFiles[file] = true
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("error reading inputs path %s: %w", path, err)
	}

	foundRequired := make(map[string]bool)
	for _, entry := range entries {
		fileName := entry.Name()

		if entry.IsDir() {
			return fmt.Errorf("subdirectory %s found, but subdirectories are not allowed in inputs directory", fileName)
		}

		if !allowedFiles[fileName] {
			return fmt.Errorf("file %s found but not listed in required or optional files", fileName)
		}

		if slices.Contains(inputFiles.Required, fileName) {
			foundRequired[fileName] = true
		}
	}

	for _, requiredFile := range inputFiles.Required {
		if !foundRequired[requiredFile] {
			return fmt.Errorf("required file %s not found in inputs directory", requiredFile)
		}
	}

	return nil
}
