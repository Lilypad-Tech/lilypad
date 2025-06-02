package preflight

import (
	"context"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/executor/bacalhau"
	executor "github.com/Lilypad-Tech/lilypad/v2/pkg/executor/bacalhau"
	"github.com/rs/zerolog/log"
)

func (p *preflightChecker) checkBacalhau(_ context.Context, options executor.BacalhauExecutorOptions) checkResult {
	executor, err := bacalhau.NewBacalhauExecutor(options)
	if err != nil {
		return checkResult{
			passed:  false,
			message: "Bacalhau executor creation failed",
			error:   err,
		}
	}

	err = executor.Preflight()
	if err != nil {
		log.Error().Err(err).Msg("❌ Bacalhau preflight check failed")
		return checkResult{
			passed:  false,
			message: "Bacalhau preflight check failed",
			error:   err,
		}
	}

	log.Info().Msg("✅ Bacalhau passed preflight check")
	return checkResult{
		passed:  true,
		message: "Bacalhau is running and healthy",
		error:   nil,
	}
}
