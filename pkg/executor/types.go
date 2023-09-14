package executor

import (
	"github.com/bacalhau-project/lilypad/pkg/data"
)

type Executor interface {
	// run the given job and return a local folder
	// that contains the results
	RunJob(
		deal data.DealContainer,
		module data.Module,
	) (string, error)
}
