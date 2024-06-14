package resourceprovider

import (
	"context"
	"math/big"
	"time"

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

type SubmitWork func(nonce *big.Int)

type CPUMiner struct {
	submit SubmitWork

	runningWorkers []*Worker

	numWorkers int

	task chan Task

	updateHashes chan uint64
}

func NewCpuMiner(nodeId string, numWorkers int, task chan Task, submit SubmitWork) *CPUMiner {
	return &CPUMiner{
		numWorkers:   numWorkers,
		task:         task,
		updateHashes: make(chan uint64),
		submit:       submit,
	}
}

func (m *CPUMiner) Start(ctx context.Context) {
	go m.miningWorkerController(ctx)
	go m.speedMonitor(ctx)
}

// speedMonitor handles tracking the number of hashes per second the mining
// process is performing.  It must be run as a goroutine.
func (m *CPUMiner) speedMonitor(ctx context.Context) {
	log.Debug().Msg("CPU miner speed monitor started")
	var hashesPerSec float64
	var totalHashes uint64
	ticker := time.NewTicker(time.Second * hpsUpdateSecs)
	defer ticker.Stop()

out:
	for {
		select {
		// Periodic updates from the workers with how many hashes they
		// have performed.
		case numHashes := <-m.updateHashes:
			totalHashes += numHashes

		// Time to update the hashes per second.
		case <-ticker.C:
			curHashesPerSec := float64(totalHashes) / hpsUpdateSecs
			if hashesPerSec == 0 {
				hashesPerSec = curHashesPerSec
			}
			hashesPerSec = (hashesPerSec + curHashesPerSec) / 2
			if totalHashes != 0 && hashesPerSec != 0 {
				log.Info().Msgf("Hash speed: %6.0f kilohashes/s",
					hashesPerSec/1000)
			}
			totalHashes = 0
		case <-ctx.Done():
			break out
		}
	}

	log.Debug().Msgf("CPU miner speed monitor done")
}

func (m *CPUMiner) miningWorkerController(ctx context.Context) {
	resultCh := make(chan TaskResult)
	launchWorkers := func(numWorkers int) {
		for i := 0; i < numWorkers; i++ {
			w := NewWorker(i, m.updateHashes, resultCh)
			m.runningWorkers = append(m.runningWorkers, w)
		}
	}

	maxUint256 := new(uint256.Int).Sub(uint256.NewInt(0), uint256.NewInt(1))
	noncePerWorker := new(uint256.Int).Div(maxUint256, uint256.NewInt(uint64(m.numWorkers)))

	// Launch the current number of workers by default.
	launchWorkers(m.numWorkers)

	stopWrokers := func() {
		for _, w := range m.runningWorkers {
			w.Stop()
		}
	}

	spawNewWork := func(allTask *Task) {
		for i, w := range m.runningWorkers {
			w.Stop()
			from := new(uint256.Int).Mul(noncePerWorker, uint256.NewInt(uint64(i)))
			end := new(uint256.Int).Mul(noncePerWorker, uint256.NewInt(uint64(i+1)))
			go w.Solve(ctx, &Task{
				Id:         allTask.Id,
				Challenge:  allTask.Challenge,
				Difficulty: allTask.Difficulty,
				From:       from,
				End:        end,
			})
		}
	}

	cache, _ := lru.New[uuid.UUID, *uint256.Int](2048) //prevent submint multiple times, consider power multiple cpu, use a little bigger value
out:
	for {
		select {
		case <-ctx.Done():
			stopWrokers()
			break out
		case result := <-resultCh:
			_, ok := cache.Get(result.Id)
			if ok {
				log.Warn().Msg("This work has been submit before")
				continue
			}
			m.submit(result.Nonce.ToBig())
			stopWrokers()
			cache.Add(result.Id, new(uint256.Int))
		case allTask := <-m.task:
			stopWrokers()
			spawNewWork(&allTask)
		}
	}
}
