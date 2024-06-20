package resourceprovider

import (
	"context"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	"github.com/rs/zerolog/log"
)

var (
	bigOne = uint256.NewInt(1)
)

type CpuWorker struct {
	cfg *WorkerConfig

	state atomic.Int32
	quit  chan chan struct{}
}

func NewCpuWorker(cfg *WorkerConfig) (Worker, error) {
	return &CpuWorker{
		cfg:  cfg,
		quit: make(chan chan struct{}, 1),
	}, nil
}

func (w *CpuWorker) Stop() {
	if w.state.Load() == 0 {
		return
	}

	//cancel previous task
	waitChan := make(chan struct{})
	select {
	case w.quit <- waitChan: // may already exit
		<-waitChan
	default:
	}
}

func (w *CpuWorker) FindSolution(ctx context.Context, task *Task) {
	w.state.Store(1)
	defer w.state.Store(0)

	nonce := task.From.Clone()
	startTime := time.Now()

	hashesCompleted := uint64(0)
	ticker := time.NewTicker(time.Second * hashUpdateSecs)
	defer ticker.Stop()

OUT:
	for {
		select {
		case <-ctx.Done():
			break OUT
		case respCh := <-w.quit:
			respCh <- struct{}{}
			return
		case <-ticker.C:
			w.cfg.updateHashes <- hashesCompleted
			hashesCompleted = 0
		default:
			// Non-blocking select to fall through
		}

		if nonce.Cmp(task.End) >= 0 {
			return
		}
		hashNumber, err := calculateHashNumber(task.Challenge, nonce.ToBig())
		if err != nil {
			log.Err(err).Msg("Calculate hash number")
			return
		}
		hashesCompleted += 1

		// Check if the hash is below the target difficulty
		if hashNumber.Cmp(task.Difficulty) == -1 {
			log.Info().Str("Elapsed Time", time.Since(startTime).String()).
				Str("challenge", new(big.Int).SetBytes(task.Challenge[:]).String()).
				Str("Nonce", nonce.String()).
				Str("HashNumber", hashNumber.String()).
				Msg("Success!")
			w.cfg.resultCh <- TaskResult{
				Id:    task.Id,
				Nonce: nonce.Clone(),
			}
		}

		nonce.Add(nonce, bigOne)
	}
}

func calculateHashNumber(challenge [32]byte, nonce *big.Int) (*uint256.Int, error) {
	data, err := formatMinerArgs(challenge, nonce)
	if err != nil {
		return nil, err
	}

	// Calculate Keccak-256 hash
	hashResult := crypto.Keccak256(data)

	return new(uint256.Int).SetBytes(hashResult), nil
}
