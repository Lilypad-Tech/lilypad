package main

import (
	"testing"

	optionsfactory "github.com/bacalhau-project/lilypad/pkg/options"
	"github.com/davecgh/go-spew/spew"
)

func TestStack(t *testing.T) {
	solverOptions := optionsfactory.NewSolverOptions()

	spew.Dump(solverOptions)
}
