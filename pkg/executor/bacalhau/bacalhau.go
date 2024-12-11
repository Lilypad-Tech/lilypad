package bacalhau

import (
	"bufio"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	"github.com/lilypad-tech/lilypad/pkg/data"
	executorlib "github.com/lilypad-tech/lilypad/pkg/executor"
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
	ApiPort               string
	JobStatusPollInterval uint64
	ResultsDirectory      string
}

type BacalhauExecutor struct {
	Options     BacalhauExecutorOptions
	bacalhauEnv []string
	bacalhauClient BacalhauClient
}

func NewBacalhauExecutor(options BacalhauExecutorOptions) (*BacalhauExecutor, error) {

	apiHost := fmt.Sprintf("http://%s:%s", options.ApiHost, options.ApiPort)

	client, err := newBacalhauClient(apiHost)
	if err != nil {
		return nil, err
	}

	return &BacalhauExecutor{
		Options: options,
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
		return false, errors.New("Bacalhau version must be v1.5.1")
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
	fmt.Printf("Running job: %s\n", deal.ID)
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

	resultsPath := filepath.Join(executor.Options.ResultsDirectory, fmt.Sprintf("%s.tar.gz", executionId))

	gzipFile, err := os.Open(resultsPath)
	if err != nil {
		return "", "", fmt.Errorf("error opening gzip file: %s", err.Error())
	}
	defer gzipFile.Close()

	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		return "", "", fmt.Errorf("error creating gzip reader: %s", err.Error())
	}
	defer gzipReader.Close()

	tarPath := strings.TrimSuffix(resultsPath, ".gz")
	tarFile, err := os.Create(tarPath)
	if err != nil {
		return "", "", fmt.Errorf("error creating tar file: %s", err.Error())
	}
	defer tarFile.Close()

	_, err = io.Copy(tarFile, gzipReader)
	if err != nil {
		return "", "", fmt.Errorf("error writing tar file: %s", err.Error())
	}

	resultsPath = tarPath

	go func() {
		if err := os.Remove(resultsPath + ".gz"); err != nil {
			log.Printf("error removing gzip file: %s", err.Error())
		}
	}()

	cid, err = GenerateCID(resultsPath)
	if err != nil {
		return "", "", fmt.Errorf("error generating CID: %s", err.Error())
	}

	return cid, resultsPath, nil

}

func GenerateCID(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open file %s: %w", path, err)
	}
	defer file.Close()

	ds := dsync.MutexWrap(datastore.NewNullDatastore())
	bs := blockstore.NewBlockstore(ds)
	bs = blockstore.NewIdStore(bs)
	bsrv := blockservice.New(bs, offline.Exchange(bs))
	dsrv := merkledag.NewDAGService(bsrv)

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
