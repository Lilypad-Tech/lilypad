package resourceprovider

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"
)

var (
	bigOne = big.NewInt(1)
)

type task struct {
	challenge  [32]byte
	difficulty *big.Int
}

type Worker struct {
	nodeId        string
	walletAddress common.Address
	quit          chan chan struct{}
}

func NewWorker(nodeId string, walletAddress common.Address) *Worker {
	return &Worker{
		nodeId:        nodeId,
		walletAddress: walletAddress,
		quit:          make(chan chan struct{}),
	}
}
func (w Worker) Reset() {
	waitChan := make(chan struct{})
	w.quit <- waitChan
	<-waitChan
}

func (w Worker) Solve(ctx context.Context, task *task) (*big.Int, error) {
	nonce := big.NewInt(0)
	startTime := time.Now()

OUT:
	for {
		select {
		case <-ctx.Done():
			break OUT
		case respCh := <-w.quit:
			respCh <- struct{}{}
			break OUT
		default:
			// Non-blocking select to fall through
		}

		data, err := formatMinerArgs(task.challenge, nonce)
		if err != nil {
			return nil, err
		}

		// Calculate Keccak-256 hash
		hashResult := crypto.Keccak256(data)

		// Check if the hash is below the target difficulty
		hashNumber := new(big.Int).SetBytes(hashResult)
		//	fmt.Println("number ", hashNumber.Text(2))
		//	fmt.Println("difficulty", difficulty.Text(2))
		if hashNumber.Cmp(task.difficulty) == -1 {
			log.Info().Msgf("Success! Nonce: %d, Elapsed Time: %d, Hash: %v", time.Since(startTime), nonce, hashResult)
			return nonce, nil
		}

		nonce.Add(nonce, bigOne)
	}

	return nil, nil
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
