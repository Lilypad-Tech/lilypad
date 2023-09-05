package solver

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/server"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/token"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type SolverOptions struct {
	Web3   web3.Web3Options
	Server server.ServerOptions
}

type Solver struct {
	Web3 *web3.ContractSDK
}

func NewSolver(
	options SolverOptions,
) (*Solver, error) {
	web3, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return nil, err
	}
	solver := &Solver{
		Web3: web3,
	}
	return solver, nil
}

func (solver *Solver) Start(ctx context.Context) error {
	go func() {
		for {
			tx, err := solver.Web3.Contracts.Token.Transfer(
				solver.Web3.Auth,
				common.HexToAddress("0x2546BcD3c84621e976D8185a91A922aE77ECEc30"),
				big.NewInt(1),
			)
			if err != nil {
				fmt.Printf("error sending tx: %s\n", err.Error())
				break
			}
			fmt.Printf("tx sent: %s\n", tx.Hash())
			time.Sleep(time.Second * 1)
		}
	}()

	block := uint64(0)

	eventCh := make(chan *token.TokenTransfer)
	sub, err := solver.Web3.Contracts.Token.WatchTransfer(
		&bind.WatchOpts{Start: &block, Context: ctx},
		eventCh,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Fatalf("Failed to watch MyEvent: %v", err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Subscription error: %v", err)
		case event := <-eventCh:
			// Do something with the event data
			log.Printf("New MyEvent. From: %s, Value: %d", event.From.Hex(), event.Value)
		}
	}
	// query := ethereum.FilterQuery{
	// 	Addresses: []common.Address{common.HexToAddress(solver.Web3.Options.TokenAddress)},
	// }

	// logs := make(chan types.Log)
	// sub, err := solver.Web3.Client.SubscribeFilterLogs(context.Background(), query, logs)
	// if err != nil {
	// 	log.Fatalf("Failed to subscribe to contract events: %v", err)
	// }

	// for {
	// 	select {
	// 	case err := <-sub.Err():
	// 		log.Fatalf("Received subscription error: %v", err)
	// 	case vLog := <-logs:
	// 		fmt.Printf("LOGS: %+v\n", vLog.Data)
	// 		//vLog.Topics
	// 		// event := new(YourEventTypeHere)                              // Change this to your specific event type
	// 		// err := contractAbi.Unpack(event, "EventNameHere", vLog.Data) // Change "EventNameHere" to your specific event name
	// 		// if err != nil {
	// 		// 	log.Fatalf("Failed to unpack event: %v", err)
	// 		// }

	// 		// Access event fields and do something
	// 	}
	// }
}
