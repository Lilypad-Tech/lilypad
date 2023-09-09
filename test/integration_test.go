package main

import (
	"path/filepath"
	"testing"

	optionsfactory "github.com/bacalhau-project/lilypad/pkg/options"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/davecgh/go-spew/spew"
)

func stack(args []string) error {
	// the tests run where the working dir is "test"
	// so to run the stack commands - we need to go up one level
	rootPath, err := filepath.Abs("..")
	if err != nil {
		return err
	}
	return system.RunCommand("./stack", args, rootPath)
}

func bootStack() error {
	return stack([]string{"boot"})
}

func stopStack() error {
	return stack([]string{"clean"})
}

func TestStack(t *testing.T) {
	// system.SetupLogging()
	// err := bootStack()
	// if err != nil {
	// 	t.Error(err)
	// }
	// defer stopStack()

	solverOptions := optionsfactory.NewSolverOptions()

	spew.Dump(solverOptions)
}
