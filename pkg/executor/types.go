package executor

import (
	"github.com/bacalhau-project/lilypad/pkg/data"
)

// real time logs we post to the solver for clients to hear
type ExecutorLogChannels struct {
	Stdout chan string
	Stderr chan string
}

type Executor interface {
	// this is given the deal container and a function the will
	// upload a local folder as the "results" of the job
	// this function is likely to be the solver client "upload results"
	// and will return the CID of the results folder
	// the executor channels are used to stream the logs back to the client
	// (via the solver)
	RunJob(
		deal data.DealContainer,
		uploadResults func(string) (string, error),
		logChannels *ExecutorLogChannels,
	) (*data.Result, error)
}

func NewExecutorLogChannels() *ExecutorLogChannels {
	return &ExecutorLogChannels{
		Stdout: make(chan string),
		Stderr: make(chan string),
	}
}
