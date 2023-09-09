package main

import (
	"os"
	"testing"

	optionsfactory "github.com/bacalhau-project/lilypad/pkg/options"
	"github.com/davecgh/go-spew/spew"
)

func TestStack(t *testing.T) {
	solverOptions := optionsfactory.NewSolverOptions()
	solverOptions.Web3.PrivateKey = os.Getenv("SOLVER_PRIVATE_KEY")

	// test that the solver private key is defined
	if solverOptions.Web3.PrivateKey == "" {
		t.Error("SOLVER_PRIVATE_KEY is not defined")
	}

	spew.Dump(solverOptions)
}
