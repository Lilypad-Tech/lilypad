package bacalhau

import (
	"bytes"
	"errors"
	"fmt"
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
	id, err := executor.getJobID(deal, module)
	if err != nil {
		return nil, err
	}

	for {
		jobInfo, err := executor.bacalhauClient.GetJob(id)
		if err != nil {
			return nil, fmt.Errorf("error getting job %s: %s", id, err.Error())
		}

		if jobInfo.Executions == nil {
			return nil, fmt.Errorf("no executions retrieved for job %s", id)
		}

		jobExecutions := jobInfo.Executions.Items

		if len(jobExecutions) > 0 {
			// Check that the job completed successfully
			if jobInfo.Job.State.StateType != models.JobStateTypeCompleted {
				return nil, fmt.Errorf("job %s did not complete yet: %s", id, jobInfo.Job.State.StateType.String())
			}
			break
		}

		time.Sleep(5 * time.Second)
	}

	system.EnsureDataDir(RESULTS_DIR)
	resultsDir := system.GetDataDir(filepath.Join(RESULTS_DIR, deal.ID))

	cid, resultsDir, err := executor.prepareResults(id)
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

func (executor *BacalhauExecutor) prepareResults(jobId string) (cid string, resultsDir string, err error) {

	tmpDir, err := os.MkdirTemp("", "bacalhau-results-*")
	if err != nil {
		return "", "", fmt.Errorf("error creating temp dir: %s", err.Error())
	}

	fmt.Printf("extracting results for job %s\n to tmpdir: %s\n", jobId, tmpDir)

	tarPath := filepath.Join("/tmp/lp-bacalhau", fmt.Sprintf("%s.tar.gz", jobId))

	err = executorlib.ExtractTarGz(tarPath, tmpDir)
	if err != nil {
		return "", "", fmt.Errorf("error extracting tar.gz: %s", err.Error())
	}

	fmt.Printf("generating CID for job %s\n", jobId)

	cid, err = GenerateCID(tmpDir)
	if err != nil {
		return "", "", fmt.Errorf("error generating CID: %s", err.Error())
	}

	return cid, tmpDir, nil

}

func GenerateCID(path string) (string, error) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error reading file: %s", err.Error())
	}
	fileReader := bytes.NewReader(fileData)

	ds := dsync.MutexWrap(datastore.NewNullDatastore())
	bs := blockstore.NewBlockstore(ds)
	bs = blockstore.NewIdStore(bs)

	bsrv := blockservice.New(bs, offline.Exchange(bs))
	dsrv := merkledag.NewDAGService(bsrv)

	ufsImportParams := uih.DagBuilderParams{
		Maxlinks:  uih.DefaultLinksPerBlock, // Default max of 174 links per block
		RawLeaves: true,                     // Leave the actual file bytes untouched instead of wrapping them in a dag-pb protobuf wrapper
		CidBuilder: cid.V1Builder{ // Use CIDv1 for all links
			Codec:    uint64(multicodec.Raw),
			MhType:   uint64(multicodec.Sha2_256), // Use SHA2-256 as the hash function
			MhLength: -1,                          // Use the default hash length for the given hash function (in this case 256 bits)
		},
		Dagserv: dsrv,
		NoCopy:  false,
	}
	ufsBuilder, err := ufsImportParams.New(chunker.NewSizeSplitter(fileReader, chunker.DefaultBlockSize)) // Split the file up into fixed sized 256KiB chunks
	if err != nil {
		return cid.Undef.String(), fmt.Errorf("error creating builder: %s", err.Error())
	}
	nd, err := balanced.Layout(ufsBuilder) // Arrange the graph with a balanced layout
	if err != nil {
		return cid.Undef.String(), fmt.Errorf("error creating builder: %s", err.Error())
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
