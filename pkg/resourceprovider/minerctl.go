package resourceprovider

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/google/uuid"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/holiman/uint256"
	"github.com/rs/zerolog/log"
)

const (
	// hpsUpdateSecs is the number of seconds to wait in between each
	// update to the hashes per second monitor.
	hpsUpdateSecs = 10

	// hashUpdateSec is the number of seconds each worker waits in between
	// notifying the speed monitor with how many hashes have been completed
	// while they are actively searching for a solution.  This is done to
	// reduce the amount of syncs between the workers that must be done to
	// keep track of the hashes per second.
	hashUpdateSecs = 15
)

type SubmitWork func(nonce *big.Int, hashrate float64)
type Worker interface {
	FindSolution(ctx context.Context, task *Task)
	Stop()
}

type WorkerConfig struct {
	id           int
	updateHashes chan uint64
	resultCh     chan TaskResult

	//cuda
	gridSize       int
	blockSize      int
	hashsPerThread int
}

type Task struct {
	Id         uuid.UUID
	Challenge  [32]byte
	Difficulty *uint256.Int
	From       *uint256.Int
	End        *uint256.Int
}

type TaskResult struct {
	Id    uuid.UUID
	Nonce *uint256.Int
}

type MinerController struct {
	submit SubmitWork

	runningWorkers []Worker

	powCfg ResourceProviderPowOptions

	task chan Task

	updateHashes chan uint64
	totalHash    int64
}

func NewMinerController(nodeId string, powCfg ResourceProviderPowOptions, task chan Task, submit SubmitWork) *MinerController {
	return &MinerController{
		powCfg:       powCfg,
		task:         task,
		updateHashes: make(chan uint64),
		submit:       submit,
	}
}

func (m *MinerController) Start(ctx context.Context) {
	go m.miningWorkerController(ctx)
	go m.speedMonitor(ctx)
}

// speedMonitor handles tracking the number of hashes per second the mining
// process is performing.  It must be run as a goroutine.
func (m *MinerController) speedMonitor(ctx context.Context) {
	log.Debug().Msg("CPU miner speed monitor started")
	var hashesPerSec float64
	var latestTotalHash uint64
	ticker := time.NewTicker(time.Second * hpsUpdateSecs)
	defer ticker.Stop()

out:
	for {
		select {
		// Periodic updates from the workers with how many hashes they
		// have performed.
		case numHashes := <-m.updateHashes:
			latestTotalHash += numHashes
			m.totalHash += int64(numHashes)

		// Time to update the hashes per second.
		case <-ticker.C:
			curHashesPerSec := float64(latestTotalHash) / hpsUpdateSecs
			if hashesPerSec == 0 {
				hashesPerSec = curHashesPerSec
			}
			hashesPerSec = (hashesPerSec + curHashesPerSec) / 2
			if latestTotalHash != 0 && hashesPerSec != 0 {
				log.Info().Msgf("Hash speed: %6.0f kilohashes/s",
					hashesPerSec/1000)
			}
			latestTotalHash = 0
		case <-ctx.Done():
			break out
		}
	}

	log.Debug().Msgf("CPU miner speed monitor done")
}

func (m *MinerController) miningWorkerController(ctx context.Context) {
	numworkers := m.powCfg.NumWorkers
	if numworkers == 0 {
		numworkers = DefaultWorkerNum()
	}

	log.Info().Int("workernum", numworkers).Msg("Worker ready for mining")

	resultCh := make(chan TaskResult, numworkers*2) //avoid lock worker if have much work to submit
	launchWorkers := func(powCfg ResourceProviderPowOptions) error {
		for i := 0; i < numworkers; i++ {
			wCfg := &WorkerConfig{
				id:           i,
				updateHashes: m.updateHashes,
				resultCh:     resultCh,

				gridSize:       powCfg.CudaGridSize,
				blockSize:      powCfg.CudaBlockSize,
				hashsPerThread: powCfg.CudaHashsPerThread,
			}

			w, err := MaybeCudaOrCpu(wCfg)
			if err != nil {
				return err
			}

			m.runningWorkers = append(m.runningWorkers, w)
		}
		return nil
	}

	// Todo this split u256 max value to multiple part, and send each part to different worker to find solution
	// But we don't need so much big range in practice, uint64 range is enough to find solution, this also benefit to optimise hardware
	maxUint256 := new(uint256.Int).Sub(uint256.NewInt(0), uint256.NewInt(1))
	noncePerWorker := new(uint256.Int).Div(maxUint256, uint256.NewInt(uint64(numworkers)))

	// Launch the current number of workers by default.
	err := launchWorkers(m.powCfg)
	if err != nil {
		log.Err(err).Msg("Cannt create worker")
	}

	stopWorkers := func() {
		var wg sync.WaitGroup
		for _, worker := range m.runningWorkers {
			wg.Add(1)
			go func(w Worker) {
				defer wg.Done()
				w.Stop()
			}(worker)
		}
		wg.Wait()
	}

	spawNewWork := func(allTask *Task) {
		for i, w := range m.runningWorkers {
			w.Stop()
			from := new(uint256.Int).Mul(noncePerWorker, uint256.NewInt(uint64(i)))
			end := new(uint256.Int).Mul(noncePerWorker, uint256.NewInt(uint64(i+1)))
			go w.FindSolution(ctx, &Task{
				Id:         allTask.Id,
				Challenge:  allTask.Challenge,
				Difficulty: allTask.Difficulty,
				From:       from,
				End:        end,
			})
		}
	}

	cache, _ := lru.New[uuid.UUID, *uint256.Int](2048) //prevent submint multiple times, consider power multiple cpu, use a little bigger value
	workStartTime := time.Now()
out:
	for {
		select {
		case <-ctx.Done():
			stopWorkers()
			break out
		case result := <-resultCh:
			_, ok := cache.Get(result.Id)
			if ok {
				log.Warn().Msg("This work has been submit before")
				continue
			}

			dur := (float64(time.Since(workStartTime).Milliseconds()) / 1000.0)
			stopWorkers()
			cache.Add(result.Id, new(uint256.Int))
			time.Sleep(time.Second * hpsUpdateSecs) //to ensure data was reported
			hashrate := float64(m.totalHash) / 1000 / 1000 / dur
			m.submit(result.Nonce.ToBig(), hashrate)
		case allTask := <-m.task:
			stopWorkers()
			m.totalHash = 0
			workStartTime = time.Now()
			spawNewWork(&allTask)
		}
	}
}

func formatMinerArgs(challenge [32]byte, nonce *big.Int) ([]byte, error) {
	//todo use nonce in replace instead of building from scratch for better performance
	// keccak256(abi.encodePacked(lastChallenge, msg.sender, nodeId));
	bytes32Ty, _ := abi.NewType("bytes32", "", nil)
	uint256Ty, _ := abi.NewType("uint256", "", nil)

	arguments := abi.Arguments{
		{
			Type: bytes32Ty,
		},
		{
			Type: uint256Ty,
		},
	}

	bytes, err := arguments.Pack(
		challenge,
		nonce,
	)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
