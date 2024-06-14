package resourceprovider

import (
	"context"
	"fmt"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/holiman/uint256"
	"github.com/lilypad-tech/lilypad/pkg/web3"
	"github.com/rs/zerolog/log"
)

var (
	bigOne = uint256.NewInt(1)
)

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

type Worker struct {
	id    int
	state atomic.Int32

	updateHashes chan uint64

	resultCh chan TaskResult
	quit     chan chan struct{}
}

func NewWorker(id int, updateHashes chan uint64, resultCh chan TaskResult) *Worker {
	return &Worker{
		id:           id,
		updateHashes: updateHashes,
		resultCh:     resultCh,
		quit:         make(chan chan struct{}, 1),
	}
}
func (w *Worker) Stop() {
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

func (w *Worker) Solve(ctx context.Context, task *Task) {
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
			w.updateHashes <- hashesCompleted
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
			w.resultCh <- TaskResult{
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

func TriggerNewPowRound(ctx context.Context, web3SDK *web3.Web3SDK) (common.Hash, error) {
	tx, err := web3SDK.Contracts.Pow.TriggerNewPowRound(web3SDK.TransactOpts)
	if err != nil {
		return common.Hash{}, err
	}

	receipt, err := web3SDK.WaitTx(ctx, tx)
	if err != nil {
		return common.Hash{}, err
	}

	if receipt.Status != 1 {
		return tx.Hash(), fmt.Errorf("trigger new pow round")
	}
	return tx.Hash(), nil
}
