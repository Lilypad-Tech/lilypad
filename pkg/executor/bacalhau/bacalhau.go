package bacalhau

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"

	config "github.com/ipfs/kubo/config"
	core "github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/repo/fsrepo"
	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/data/bacalhau"
	executorlib "github.com/lilypad-tech/lilypad/pkg/executor"

	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/rs/zerolog/log"

	coreapi "github.com/ipfs/kubo/core/coreapi"
	coreiface "github.com/ipfs/kubo/core/coreiface"
	iface "github.com/ipfs/kubo/core/coreiface"
	"github.com/ipfs/kubo/core/node/libp2p"

	"github.com/ipfs/kubo/plugin/loader"
)

const RESULTS_DIR = "bacalhau-results"

type BacalhauExecutorOptions struct {
	ApiHost string
	ApiPort string
}

type BacalhauExecutor struct {
	Options        BacalhauExecutorOptions
	bacalhauEnv    []string
	bacalhauClient BacalhauClient
}

var (
	node *core.IpfsNode
	n    iface.CoreAPI
)

func setupPlugins(externalPluginsPath string) error {
	// Load any external plugins if available on externalPluginsPath
	plugins, err := loader.NewPluginLoader(filepath.Join(externalPluginsPath, "plugins"))
	if err != nil {
		return fmt.Errorf("error loading plugins: %s", err)
	}

	// Load preloaded and external plugins
	if err := plugins.Initialize(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	if err := plugins.Inject(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	return nil
}

var loadPluginsOnce sync.Once

// CreateRepo Create a Temporary Repo
func CreateRepo() (string, error) {
	var onceErr error
	loadPluginsOnce.Do(func() {
		onceErr = setupPlugins("")
	})
	if onceErr != nil {
		return "", onceErr
	}

	repoPath, err := os.MkdirTemp("", "ipfs-shell")
	if err != nil {
		return "", fmt.Errorf("failed to get temp dir: %s", err)
	}

	// Create a config with default options and a 2048 bit key
	cfg, err := config.Init(io.Discard, 2048)

	if err != nil {
		return "", err
	}

	// Create the repo with the config
	err = fsrepo.Init(repoPath, cfg)
	if err != nil {
		return "", fmt.Errorf("failed to init ephemeral node: %s", err)
	}

	return repoPath, nil
}

// createNode Creates an IPFS node and returns its coreAPI
func createNode(ctx context.Context, repoPath string) (*core.IpfsNode, error) {
	// Open the repo
	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		return nil, err
	}

	// Construct the node
	nodeOptions := &core.BuildCfg{
		Online:  true,
		Routing: libp2p.DHTOption, // This option sets the node to be a full DHT node (both fetching and storing DHT Records)
		// Routing: libp2p.DHTClientOption, // This option sets the node to be a client DHT node (only fetching records)
		Repo: repo,
	}

	return core.NewNode(ctx, nodeOptions)
}

type APIHandler struct {
	api coreiface.CoreAPI
}

func NewAPIHandler(api coreiface.CoreAPI) *APIHandler {
	return &APIHandler{api: api}
}

func (h *APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/v0/id":
		h.handleID(w, r)
	case "/api/v0/add":
		h.handleAdd(w, r)
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

type IDResponse struct {
	ID           string   `json:"ID"`
	PublicKey    string   `json:"PublicKey"`
	Addresses    []string `json:"Addresses"`
	AgentVersion string   `json:"AgentVersion"`
	Protocols    []string `json:"Protocols"`
}

func (h *APIHandler) handleID(w http.ResponseWriter, r *http.Request) {
	// id, err := h.api.Key().Self(n.)
	//id := node.Identity
	resp := IDResponse{
		ID:           node.Identity.String(),
		PublicKey:    "CAESIIzbqSqqlEOcZBlOLXS94+h2efs3blK47fsgMd9H6Okj", //node.Identity.PublicKey.String(),
		Addresses:    []string{"/ip4/127.0.0.1/tcp/4001/p2p/12D3KooWKJDWrpxAD2Aa4FgdPVXVPVAHRjHBahHaiLfLjiao6EyQ"},
		AgentVersion: "kubo/0.28.0/",
		Protocols: []string{"/ipfs/bitswap",

			"/ipfs/bitswap/1.0.0",
			"/ipfs/bitswap/1.1.0",
			"/ipfs/bitswap/1.2.0",
			"/ipfs/id/1.0.0",
			"/ipfs/id/push/1.0.0",
			"/ipfs/lan/kad/1.0.0",
			"/ipfs/ping/1.0.0",
			"/libp2p/autonat/1.0.0",
			"/libp2p/circuit/relay/0.2.0/stop",
			"/libp2p/dcutr",
			"/x/"},
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func (h *APIHandler) handleAdd(w http.ResponseWriter, r *http.Request) {
	// Implement your logic to handle the `/api/v0/add` endpoint
	// ...
}
func NewBacalhauExecutor(options BacalhauExecutorOptions) (*BacalhauExecutor, error) {
	client, err := newClient(options)
	if err != nil {
		fmt.Println("Error:", err)
	}
	bacalhauEnv := []string{
		fmt.Sprintf("BACALHAU_API_HOST=%s", options.ApiHost),
		fmt.Sprintf("HOME=%s", os.Getenv("HOME")),
	}
	s, e := CreateRepo()
	if e != nil {
		fmt.Println("Error:", e)
	}
	// var err error
	node, err = createNode(context.Background(), s)
	if err != nil {
		// Handle error
	}
	// Obtain the CoreAPI instance
	api, err := coreapi.NewCoreAPI(node)
	if err != nil {
		// Handle error
	}
	// Create an APIHandler instance
	handler := NewAPIHandler(api)

	// Start the HTTP server on port 5001
	fmt.Println("Starting server on :5001")
	go http.ListenAndServe(":5001", handler)

	fmt.Println("ID " + node.Identity.String())
	if err != nil {
		fmt.Println("Error:", err)

	}
	n, err = coreapi.NewCoreAPI(node)

	if err != nil {
		fmt.Println("Error initializing IPFS node:", err)
		// return "", err
	}

	log.Debug().Msgf("bacalhauEnv: %s", bacalhauEnv)
	return &BacalhauExecutor{
		Options:        options,
		bacalhauEnv:    bacalhauEnv,
		bacalhauClient: *client,
	}, nil
}
func (executor *BacalhauExecutor) Id() (string, error) {
	nodeIdCmd := exec.Command(
		"bacalhau",
		"id",
	)
	nodeIdCmd.Env = executor.bacalhauEnv

	runOutputRaw, err := nodeIdCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error calling get id results %s, %s", err.Error(), runOutputRaw)
	}

	splitOutputs := strings.Split(strings.Trim(string(runOutputRaw), " \t\n"), "\n")
	runOutput := splitOutputs[len(splitOutputs)-1]

	var idResult struct {
		ID       string
		ClientID string
	}
	err = json.Unmarshal([]byte(runOutput), &idResult)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling job JSON %s %s", err.Error(), runOutputRaw)
	}

	return idResult.ID, nil
}

// Checks that Bacalhau is installed, correctly configured, and available
func (executor *BacalhauExecutor) IsAvailable() (bool, error) {
	isAlive, err := executor.bacalhauClient.alive()
	if !isAlive || err != nil {
		return false, fmt.Errorf("Bacalhau is not currently available. Please ensure that Bacalhau is running, then try again. %w", err)
	}

	// Check that we have the right version of Bacalhau
	version, err := executor.bacalhauClient.getVersion()
	if err != nil {
		return false, err
	}
	// TODO: we may want to relax this
	if version.GitVersion != "v1.3.2" {
		return false, errors.New("Bacalhau version must be v1.3.2")
	}

	return true, nil
}

func (executor *BacalhauExecutor) GetMachineSpecs() ([]data.MachineSpec, error) {
	var specs []data.MachineSpec
	result, err := executor.bacalhauClient.getNodes()
	if err != nil {
		return specs, err
	}

	for _, node := range result.Nodes {
		spec := data.MachineSpec{
			CPU:  int(node.Info.ComputeNodeInfo.MaxCapacity.CPU) * 1000, // convert float to "milli-CPU"
			RAM:  int(node.Info.ComputeNodeInfo.MaxCapacity.Memory),
			GPU:  int(node.Info.ComputeNodeInfo.MaxCapacity.GPU),
			Disk: int(node.Info.ComputeNodeInfo.MaxCapacity.Disk),
		}
		for _, gpu := range node.Info.ComputeNodeInfo.MaxCapacity.GPUs {
			spec.GPUs = append(spec.GPUs, data.GPUSpec{
				Name:   gpu.Name,
				Vendor: string(gpu.Vendor),
				VRAM:   int(gpu.Memory),
			})
		}
		specs = append(specs, spec)
	}
	return specs, nil
}
func (executor *BacalhauExecutor) RunJob(
	deal data.DealContainer,
	module data.Module,
) (*executorlib.ExecutorResults, error) {
	id, err := executor.getJobID(deal, module)
	if err != nil {
		return nil, err
	}

	resultsDir, err := executor.copyJobResults(deal.ID, id)
	if err != nil {
		return nil, err
	}

	jobState, err := executor.getJobState(deal.ID, id)
	if err != nil {
		return nil, err
	}

	if len(jobState.State.Executions) <= 0 {
		return nil, fmt.Errorf("no executions found for job %s", id)
	}

	if jobState.State.State != bacalhau.JobStateCompleted {
		return nil, fmt.Errorf("job %s did not complete successfully: %s", id, jobState.State.State.String())
	}

	// TODO: we should think about WASM and instruction count here
	results := &executorlib.ExecutorResults{
		ResultsDir:       resultsDir,
		ResultsCID:       jobState.State.Executions[0].PublishedResult.CID,
		InstructionCount: 1,
	}

	return results, nil
}

// run the bacalhau job and return the job ID
func (executor *BacalhauExecutor) getJobID(
	deal data.DealContainer,
	module data.Module,
) (string, error) {
	// get a JSON string of the job
	jsonBytes, err := json.Marshal(module.Job)
	if err != nil {
		return "", fmt.Errorf("error getting job JSON for deal %s -> %s", deal.ID, err.Error())
	}
	bacalhauJobSpecDir, err := system.EnsureDataDir(filepath.Join("bacalhau-job-specs", deal.ID))
	if err != nil {
		return "", fmt.Errorf("error creating a local folder for job specs %s -> %s", deal.ID, err.Error())
	}
	jobPath := filepath.Join(bacalhauJobSpecDir, "job.json")
	err = system.WriteFile(jobPath, jsonBytes)
	if err != nil {
		return "", fmt.Errorf("error writing job JSON %s -> %s", deal.ID, err.Error())
	}
	// jobPath = "/home/arsen/lilypad/job.json"
	jobBytes, err := ioutil.ReadFile(jobPath)
	if err != nil {
		return "", fmt.Errorf("error reading job file %s -> %s", deal.ID, err.Error())
	}
	fmt.Println("jobPath ", string(jobBytes))

	runCmd := exec.Command(
		"bacalhau",
		"create",
		"--id-only",
		"--wait",
		jobPath,
	)
	runCmd.Env = executor.bacalhauEnv

	runOutput, err := runCmd.CombinedOutput()
	fmt.Println("runOutput " + string(runOutput))
	if err != nil {
		return "", fmt.Errorf("error running command %s -> %s, %s", deal.ID, err.Error(), runOutput)
	}

	id := strings.TrimSpace(string(runOutput))
	fmt.Printf("Got bacalhau job ID: %s\n", id)

	return id, nil
}

// func downloadDirectory(node *core.IpfsNode, cid path.Path, p string) error {
func downloadDirectory(n iface.CoreAPI, cid path.Path, p string) error {

	// Get directory contents
	unixfs := n.Unixfs()
	dirEntries, err := unixfs.Ls(context.Background(), cid)
	if err != nil {
		return err
	}
	for entry := range dirEntries {
		// Construct path for the current entry
		entryPath := p + "/" + entry.Name
		if entry.Type == iface.TDirectory {
			// If it's a directory, create corresponding local directory and download its contents
			if err := os.MkdirAll(entryPath, 0755); err != nil {
				return err
			}
			if err := downloadDirectory(n, path.FromCid(entry.Cid), entryPath); err != nil {
				return err
			}
		} else {
			// If it's a file, download it
			file, err := unixfs.Get(context.Background(), path.FromCid(entry.Cid))
			if err != nil {
				return err
			}
			defer file.Close()

			// Read file content
			// content, err := ioutil.ReadAll(file)
			f, ok := file.(files.File)
			if !ok {
				// handle error, file is not a regular file but could be a directory
			}
			content, err := ioutil.ReadAll(f)
			if err != nil {
				return err
			}

			// Write content to local file
			if err := ioutil.WriteFile(entryPath, content, 0644); err != nil {
				return err
			}
		}
	}

	return nil
}
func (executor *BacalhauExecutor) copyJobResults(dealID string, jobID string) (string, error) {
	resultsDir, err := system.EnsureDataDir(filepath.Join(RESULTS_DIR, dealID))
	if err != nil {
		return "", fmt.Errorf("error creating a local folder of results %s -> %s", dealID, err.Error())
	}

	//example bacalhau list --id-filter 2f5b845a-6ad3-41f1-969c-18a4209082e7 --output json
	info := exec.Command("bacalhau", "list", "--id-filter", jobID, "--output", "json")
	info.Env = executor.bacalhauEnv
	output, err := info.Output()
	fmt.Println("bacalhau list output ", string(output))

	var jobData []JobData
	err = json.Unmarshal([]byte(output), &jobData)
	if err != nil {
		fmt.Println("Error:", err)

	}
	if len(jobData) == 0 {
		return "", fmt.Errorf("no job data found for job %s check that bacalhau is running", jobID)

	}

	// Example CID of the directory you want to download
	cidObject, err := cid.Decode(jobData[0].State.Executions[0].PublishedResults.CID)
	cid := path.FromCid(cidObject)
	if err != nil {
		fmt.Println("Error parsing CID:", err)

	}

	// Download the directory recursively
	err = downloadDirectory(n, cid, resultsDir)
	if err != nil {
		fmt.Println("Error downloading directory:", err)

	}
	fmt.Println("Directory downloaded successfully.")
	return resultsDir, nil

}

func (executor *BacalhauExecutor) getJobState(dealID string, jobID string) (*bacalhau.JobWithInfo, error) {
	describeCmd := exec.Command(
		"bacalhau",
		"describe",
		"--json",
		jobID,
	)
	describeCmd.Env = executor.bacalhauEnv

	output, err := describeCmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error calling describe command results %s -> %s", dealID, err.Error())
	}

	var job bacalhau.JobWithInfo
	err = json.Unmarshal(output, &job)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling job JSON %s -> %s", dealID, err.Error())
	}

	return &job, nil
}

// Compile-time interface check:
var _ executorlib.Executor = (*BacalhauExecutor)(nil)

type JobMetadata struct {
	ID        string    `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	ClientID  string    `json:"ClientID"`
	Requester struct {
		RequesterNodeID    string `json:"RequesterNodeID"`
		RequesterPublicKey string `json:"RequesterPublicKey"`
	} `json:"Requester"`
}

type JobSpec struct {
	Engine        string `json:"Engine"`
	Verifier      string `json:"Verifier"`
	Publisher     string `json:"Publisher"`
	PublisherSpec struct {
		Type string `json:"Type"`
	} `json:"PublisherSpec"`
	Docker struct {
		Image      string   `json:"Image"`
		Entrypoint []string `json:"Entrypoint"`
	} `json:"Docker"`
}

type JobResources struct {
	GPU string `json:"GPU"`
}

type JobNetwork struct {
	Type string `json:"Type"`
}

type JobDeal struct {
	Concurrency int `json:"Concurrency"`
}

type Job struct {
	APIVersion string       `json:"APIVersion"`
	Metadata   JobMetadata  `json:"Metadata"`
	Spec       JobSpec      `json:"Spec"`
	Resources  JobResources `json:"Resources"`
	Network    JobNetwork   `json:"Network"`
	Timeout    int          `json:"Timeout"`
	Deal       JobDeal      `json:"Deal"`
}

type JobExecution struct {
	JobID              string `json:"JobID"`
	NodeId             string `json:"NodeId"`
	ComputeReference   string `json:"ComputeReference"`
	State              string `json:"State"`
	AcceptedAskForBid  bool   `json:"AcceptedAskForBid"`
	VerificationResult struct {
		Complete bool `json:"Complete"`
		Result   bool `json:"Result"`
	} `json:"VerificationResult"`
	PublishedResults struct {
		StorageSource string `json:"StorageSource"`
		Name          string `json:"Name"`
		CID           string `json:"CID"`
	} `json:"PublishedResults"`
	RunOutput struct {
		Stdout          string `json:"stdout"`
		Stdouttruncated bool   `json:"stdouttruncated"`
		Stderr          string `json:"stderr"`
		Stderrtruncated bool   `json:"stderrtruncated"`
		ExitCode        int    `json:"exitCode"`
		RunnerError     string `json:"runnerError"`
	} `json:"RunOutput"`
	Version    int       `json:"Version"`
	CreateTime time.Time `json:"CreateTime"`
	UpdateTime time.Time `json:"UpdateTime"`
}

type JobState struct {
	JobID      string         `json:"JobID"`
	Executions []JobExecution `json:"Executions"`
	State      string         `json:"State"`
	Version    int            `json:"Version"`
	CreateTime time.Time      `json:"CreateTime"`
	UpdateTime time.Time      `json:"UpdateTime"`
	TimeoutAt  time.Time      `json:"TimeoutAt"`
}

type JobData struct {
	Job   Job      `json:"Job"`
	State JobState `json:"State"`
}
type Log struct {
	id        string
	timestamp string
	details   string
}
type Event struct {
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	Details   string `json:"details"`
}
type Update struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}
type Data struct {
	Dealid   string `json:"dealid"`
	State    string `json:"state"`
	Metadata string `json:"metadata"`
}
type JobUpdate struct {
	ID         string `json:"id"`
	ModuleID   string `json:"module_id"`
	JobID      string `json:"job_id"`
	Status     string `json:"status"`
	TimeUpdate string `json:"time_update"`
	Details    string `json:"details"`
	TimeStart  string `json:"time_start"`
	// Message     string `json:"message"`
}
type DbJob struct {
	id          string
	module_id   string
	job_id      string
	status      string
	time_update string
	details     string
	time_start  string
}
