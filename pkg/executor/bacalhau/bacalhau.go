package bacalhau

import (
	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/executor"
)

type BacalhauExecutorOptions struct {
	ApiHost string
}

type BacalhauExecutor struct {
	Options BacalhauExecutorOptions
}

func NewBacalhauExecutor(options BacalhauExecutorOptions) (*BacalhauExecutor, error) {
	return &BacalhauExecutor{
		Options: options,
	}, nil
}

func (executor *BacalhauExecutor) RunJob(
	deal data.DealContainer,
	uploadResults func(string) (string, error),
	logChannels *executor.ExecutorLogChannels,
) (*data.Result, error) {
	return nil, nil
}

// Compile-time interface check:
var _ executor.Executor = (*BacalhauExecutor)(nil)
