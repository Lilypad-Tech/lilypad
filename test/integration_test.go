package main

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	optionsfactory "github.com/bacalhau-project/lilypad/pkg/options"
	"github.com/bacalhau-project/lilypad/pkg/solver"
	solvermemorystore "github.com/bacalhau-project/lilypad/pkg/solver/store/memory"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/davecgh/go-spew/spew"
)

func getSolver(t *testing.T, systemContext *system.CommandContext) (*solver.Solver, error) {
	solverOptions := optionsfactory.NewSolverOptions()
	solverOptions.Web3.PrivateKey = os.Getenv("SOLVER_PRIVATE_KEY")

	// test that the solver private key is defined
	if solverOptions.Web3.PrivateKey == "" {
		return nil, fmt.Errorf("SOLVER_PRIVATE_KEY is not defined")
	}

	spew.Dump(solverOptions)

	web3SDK, err := web3.NewContractSDK(solverOptions.Web3)
	if err != nil {
		return nil, err
	}

	solverStore, err := solvermemorystore.NewSolverStoreMemory()
	if err != nil {
		return nil, err
	}

	return solver.NewSolver(solverOptions, solverStore, web3SDK)
}

func TestStack(t *testing.T) {
	commandCtx := system.NewSystemContext(context.Background())
	defer commandCtx.Cleanup()

	solver, err := getSolver(t, commandCtx)
	if err != nil {
		t.Error(err)
		return
	}

	err = solver.Start(commandCtx.Ctx, commandCtx.Cm)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("Solver started\n")

	time.Sleep(time.Second * 5)

}
