package bacalhau

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/lilypad-tech/lilypad/pkg/data"
	executorlib "github.com/lilypad-tech/lilypad/pkg/executor"
	"github.com/lilypad-tech/lilypad/pkg/ipfs"

	"github.com/bacalhau-project/bacalhau/pkg/models"
	"github.com/ipfs/boxo/blockservice"
	blockstore "github.com/ipfs/boxo/blockstore"
	chunker "github.com/ipfs/boxo/chunker"
	offline "github.com/ipfs/boxo/exchange/offline"
	"github.com/ipfs/boxo/ipld/merkledag"
	"github.com/ipfs/boxo/ipld/unixfs/importer/balanced"
	uih "github.com/ipfs/boxo/ipld/unixfs/importer/helpers"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	dsync "github.com/ipfs/go-datastore/sync"
	"github.com/lilypad-tech/lilypad/pkg/system"
	multicodec "github.com/multiformats/go-multicodec"
)

const RESULTS_DIR = "bacalhau-results"

type BacalhauExecutorOptions struct {
	ApiHost string
	ApiPort string
}

type BacalhauExecutor struct {
	Options     BacalhauExecutorOptions
	bacalhauEnv []string

	bacalhauClient BacalhauClient
	ipfsClient     *ipfs.Client
}

func NewBacalhauExecutor(options BacalhauExecutorOptions, ipfsClient *ipfs.Client) (*BacalhauExecutor, error) {

	// apiHost := ""
	// if options.ApiHost != "DO_NOT_SET" {
	// 	apiHost = fmt.Sprintf("BACALHAU_API_HOST=%s", options.ApiHost)
	// }
	// fmt.Printf("OPTIONS: %s\n", options)
	// bacalhauEnv := []string{
	// 	apiHost,
	// 	fmt.Sprintf("HOME=%s", os.Getenv("HOME")),
	// }

	// log.Debug().Msgf("bacalhauEnv: %s", bacalhauEnv)
	apiHost := fmt.Sprintf("http://%s:%s", "localhost", "20000")

	client, err := NewBacalhauClient(apiHost)
	if err != nil {
		return nil, err
	}

	return &BacalhauExecutor{
		Options: options,

		bacalhauClient: *client,
		ipfsClient:     ipfsClient,
	}, nil
}

func (executor *BacalhauExecutor) Id() (string, error) {
	id, err := executor.bacalhauClient.GetID()
	if err != nil {
		return "", fmt.Errorf("error getting bacalhau ID %s", err.Error())
	}
	return id, nil
}

// Checks that Bacalhau is installed, correctly configured, and available
func (executor *BacalhauExecutor) IsAvailable() (bool, error) {
	alive, err := executor.bacalhauClient.Alive()
	if !alive || err != nil {
		return false, fmt.Errorf("Bacalhau is not currently available. Please ensure that Bacalhau is running, then try again. %w", err)
	}

	// Check that we have the right version of Bacalhau
	version, err := executor.bacalhauClient.GetVersion()
	if err != nil {
		return false, fmt.Errorf("error getting bacalhau version %s", err.Error())
	}
	// TODO: we may want to relax this
	if version != "v1.5.1" {
		return false, errors.New("Bacalhau version must be v1.5.1")
	}

	return true, nil
}

func (executor *BacalhauExecutor) GetMachineSpecs() ([]data.MachineSpec, error) {
	return executor.bacalhauClient.GetMachineSpecs()
}

func (executor *BacalhauExecutor) RunJob(
	deal data.DealContainer,
	module data.Module,
) (*executorlib.ExecutorResults, error) {
	fmt.Printf("Running job: %s\n", deal.ID)
	jobId, err := executor.getJobID(deal, module)
	if err != nil {
		return nil, err
	}

	var jobExecutions []*models.Execution
	for {
		jobInfo, err := executor.bacalhauClient.GetJob(jobId)
		if err != nil {
			return nil, fmt.Errorf("error getting job %s: %s", jobId, err.Error())
		}

		if jobInfo.Executions == nil {
			return nil, fmt.Errorf("no executions retrieved for job %s", jobId)
		}

		jobExecutions = jobInfo.Executions.Items

		if len(jobExecutions) > 0 {
			// Check that the job completed successfully
			if jobInfo.Job.State.StateType != models.JobStateTypeCompleted {
				return nil, fmt.Errorf("job %s did not complete yet: %s", jobId, jobInfo.Job.State.StateType.String())
			}
			break
		}

		time.Sleep(5 * time.Second)
	}

	system.EnsureDataDir(RESULTS_DIR)
	resultsDir := system.GetDataDir(filepath.Join(RESULTS_DIR, deal.ID))

	cid, resultsDir, err := executor.prepareResults(jobId, jobExecutions[0].ID)
	if err != nil {
		return nil, fmt.Errorf("error preparing results: %s", err.Error())
	}

	results := &executorlib.ExecutorResults{
		ResultsDir:       resultsDir,
		ResultsCID:       cid,
		InstructionCount: 1,
	}

	return results, nil
}

func (executor *BacalhauExecutor) prepareResults(jobId string, executionId string) (cid string, resultsDir string, err error) {

	home, err := os.UserHomeDir()
	if err != nil {
		return "", "", fmt.Errorf("error getting home directory: %s", err.Error())
	}

	resultsPath := filepath.Join(home, ".bacalhau", "compute", "results", "local-publisher", fmt.Sprintf("%s.tar.gz", executionId))

	fmt.Printf("generating CID for job %s\n with execution %s\n", jobId, executionId)
	cid, err = GenerateCID(resultsPath)
	if err != nil {
		return "", "", fmt.Errorf("error generating CID: %s", err.Error())
	}

	fmt.Printf("CID for job %s with execution %s: %s\n", jobId, executionId, cid)

	return cid, resultsPath, nil

}

func GenerateCID(path string) (string, error) {
	// Open file for streaming
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open file %s: %w", path, err)
	}
	defer file.Close()

	// Create services with cleanup
	ds := dsync.MutexWrap(datastore.NewNullDatastore())
	bs := blockstore.NewBlockstore(ds)
	bs = blockstore.NewIdStore(bs)
	bsrv := blockservice.New(bs, offline.Exchange(bs))
	dsrv := merkledag.NewDAGService(bsrv)

	// Ensure cleanup
	defer func() {
		if err := bsrv.Close(); err != nil {
			log.Printf("failed to close blockservice: %v", err)
		}
	}()

	ufsImportParams := uih.DagBuilderParams{
		Maxlinks:  uih.DefaultLinksPerBlock,
		RawLeaves: true,
		CidBuilder: cid.V1Builder{
			Codec:    uint64(multicodec.Raw),
			MhType:   uint64(multicodec.Sha2_256),
			MhLength: -1,
		},
		Dagserv: dsrv,
		NoCopy:  false,
	}

	// Use buffered reader for better performance
	bufReader := bufio.NewReader(file)

	ufsBuilder, err := ufsImportParams.New(chunker.NewSizeSplitter(bufReader, chunker.DefaultBlockSize))
	if err != nil {
		return "", fmt.Errorf("failed to create UnixFS builder: %w", err)
	}

	nd, err := balanced.Layout(ufsBuilder)
	if err != nil {
		return "", fmt.Errorf("failed to create DAG layout: %w", err)
	}

	return nd.Cid().String(), nil
}

// run the bacalhau job and return the job ID
func (executor *BacalhauExecutor) getJobID(
	deal data.DealContainer,
	module data.Module,
) (string, error) {
	putJobResponse, err := executor.bacalhauClient.PostJob(module.Job)
	if err != nil {
		return "", fmt.Errorf("error creating job %s -> %s", deal.ID, err.Error())
	}

	id := putJobResponse.JobID

	return id, nil
}

// Compile-time interface check:
var _ executorlib.Executor = (*BacalhauExecutor)(nil)
