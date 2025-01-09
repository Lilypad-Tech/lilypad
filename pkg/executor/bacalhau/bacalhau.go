package bacalhau

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/bacalhau-project/bacalhau/pkg/models"
	"github.com/ipfs/boxo/files"
	"github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/core/coreapi"
	"github.com/ipfs/kubo/core/coreiface/options"
	"github.com/lilypad-tech/lilypad/pkg/data"
	executorlib "github.com/lilypad-tech/lilypad/pkg/executor"
	"github.com/lilypad-tech/lilypad/pkg/system"
)

const RESULTS_DIR = "bacalhau-results"

type BacalhauExecutorOptions struct {
	ApiHost               string
	ApiPort               string
	JobStatusPollInterval uint64
	ResultsDirectory      string
}

type BacalhauExecutor struct {
	Options        BacalhauExecutorOptions
	bacalhauEnv    []string
	bacalhauClient BacalhauClient
}

func NewBacalhauExecutor(options BacalhauExecutorOptions) (*BacalhauExecutor, error) {

	apiHost := fmt.Sprintf("http://%s:%s", options.ApiHost, options.ApiPort)

	client, err := newBacalhauClient(apiHost)
	if err != nil {
		return nil, err
	}

	return &BacalhauExecutor{
		Options:        options,
		bacalhauClient: *client,
	}, nil
}

func (executor *BacalhauExecutor) Id() (string, error) {
	id, err := executor.bacalhauClient.getID()
	if err != nil {
		return "", fmt.Errorf("error getting bacalhau ID %s", err.Error())
	}
	return id, nil
}

// Checks that Bacalhau is installed, correctly configured, and available
func (executor *BacalhauExecutor) IsAvailable() (bool, error) {
	alive, err := executor.bacalhauClient.alive()
	if !alive || err != nil {
		return false, fmt.Errorf("Bacalhau is not currently available. Please ensure that Bacalhau is running, then try again. %w", err)
	}

	version, err := executor.bacalhauClient.getVersion()
	if err != nil {
		return false, fmt.Errorf("error getting bacalhau version %s", err.Error())
	}

	if version < "v1.5.1" {
		return false, errors.New("Bacalhau version must be greater than v1.5.1")
	}

	return true, nil
}

func (executor *BacalhauExecutor) GetMachineSpecs() ([]data.MachineSpec, error) {
	nodes, err := executor.bacalhauClient.getNodes()
	var specs []data.MachineSpec
	for _, node := range nodes {
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
	if err != nil {
		return nil, err
	}
	return specs, nil
}

func (executor *BacalhauExecutor) RunJob(
	deal data.DealContainer,
	module data.Module,
) (*executorlib.ExecutorResults, error) {
	jobId, err := executor.getJobID(deal, module)
	if err != nil {
		return nil, err
	}

	var jobExecutions []*models.Execution
	for {
		jobInfo, err := executor.bacalhauClient.getJob(jobId)
		if err != nil {
			return nil, fmt.Errorf("error getting job %s: %s", jobId, err.Error())
		}

		if jobInfo.Executions == nil {
			return nil, fmt.Errorf("no executions retrieved for job %s", jobId)
		}

		jobExecutions = jobInfo.Executions.Items

		if len(jobExecutions) > 0 {
			if jobInfo.Job.State.StateType == models.JobStateTypeCompleted {
				break
			}
		}

		time.Sleep(time.Duration(executor.Options.JobStatusPollInterval) * time.Second)
	}

	resultsDir, err := system.EnsureDataDir(filepath.Join(RESULTS_DIR, deal.ID))
	if err != nil {
		return nil, fmt.Errorf("error creating results directory: %s", err.Error())
	}
	outputDir, err := executor.fetchResults(jobId, resultsDir)
	if err != nil {
		return nil, fmt.Errorf("error fetching results: %s", err.Error())
	}

	cid, err := executor.prepareResults(outputDir)
	if err != nil {
		return nil, fmt.Errorf("error preparing results: %s", err.Error())
	}

	results := &executorlib.ExecutorResults{
		ResultsDir:       outputDir,
		ResultsCID:       cid,
		InstructionCount: 1,
	}

	return results, nil
}

func (executor *BacalhauExecutor) fetchResults(jobId string, resultsDir string) (string, error) {
	resultsUrl, err := executor.bacalhauClient.getJobResult(jobId)
	if err != nil {
		return "", fmt.Errorf("error fetching results: %s", err.Error())
	}

	// We need to make sure we use BACALHAU_HOST to get the correct URL
	u, err := url.Parse(resultsUrl)
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %s", err.Error())
	}
	u.Host = fmt.Sprintf("%s:%s", executor.Options.ApiHost, u.Port())
	resultsUrl = u.String()

	// Create the file
	tarballPath := filepath.Join(resultsDir, fmt.Sprintf("%s.tar.gz", jobId))
	out, err := os.Create(tarballPath)
	if err != nil {
		return "", fmt.Errorf("error creating file: %s", err.Error())
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(resultsUrl)
	if err != nil {
		return "", fmt.Errorf("error making GET request: %s", err.Error())
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", fmt.Errorf("error writing to file: %s", err.Error())
	}

	// Extract the tar.gz file
	outputPath := filepath.Join(resultsDir, jobId)
	err = executorlib.ExtractTarGz(tarballPath, outputPath)
	if err != nil {
		return "", fmt.Errorf("error extracting tar.gz file: %s", err.Error())
	}
	return outputPath, nil
}

func (executor *BacalhauExecutor) prepareResults(resultsDir string) (string, error) {
	cid, err := generateCID(resultsDir)
	if err != nil {
		return "", fmt.Errorf("error generating CID: %s", err.Error())
	}

	return cid, nil
}

// generateCID generates a CID for a given path
// This mimics the behavior of `ipfs add -Qrn /path`
func generateCID(path string) (string, error) {
	ctx := context.Background()

	stat, err := os.Stat(path)
	if err != nil {
		return "", fmt.Errorf("error getting file info: %s", err.Error())
	}
	fsNode, err := files.NewSerialFile(path, false, stat)
	if err != nil {
		return "", fmt.Errorf("error creating serial file: %s", err.Error())
	}

	node, err := core.NewNode(ctx, &core.BuildCfg{
		Online: false,
	})
	api, err := coreapi.NewCoreAPI(node)
	if err != nil {
		return "", fmt.Errorf("error creating core API: %s", err.Error())
	}

	opts := []options.UnixfsAddOption{
		options.Unixfs.HashOnly(true),
	}
	root, err := api.Unixfs().Add(ctx, fsNode, opts...)
	if err != nil {
		return "", fmt.Errorf("error adding to UnixFS: %s", err.Error())
	}

	return root.RootCid().String(), nil
}

func (executor *BacalhauExecutor) getJobID(
	deal data.DealContainer,
	module data.Module,
) (string, error) {
	putJobResponse, err := executor.bacalhauClient.postJob(module.Job)
	if err != nil {
		return "", fmt.Errorf("error creating job %s -> %s", deal.ID, err.Error())
	}

	id := putJobResponse.JobID

	return id, nil
}

// Compile-time interface check:
var _ executorlib.Executor = (*BacalhauExecutor)(nil)
