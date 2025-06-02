package preflight

import (
	"context"
	"fmt"

	executor "github.com/Lilypad-Tech/lilypad/v2/pkg/executor/bacalhau"
	"github.com/rs/zerolog/log"
)

const RequiredGPUMemoryGB = 1 // 1GB of VRAM is required to startup if GPU is enabled

type gpuInfo struct {
	uuid          string
	name          string
	memoryTotal   int64
	driverVersion string
}

type checkResult struct {
	passed  bool
	message string
	error   error
}

type preflightConfig struct {
	GPU struct {
		MinMemoryGB int64
	}
	Docker struct {
		CheckRuntime bool
	}
	Bacalhau struct {
		Options executor.BacalhauExecutorOptions
	}
}

type preflightChecker struct {
	gpuInfo []gpuInfo
}

func RunPreflightChecks(bacalhauOptions executor.BacalhauExecutorOptions) error {
	ctx := context.Background()
	log.Info().Msg("Starting preflight checks...")
	checker := &preflightChecker{}
	config := preflightConfig{
		GPU: struct {
			MinMemoryGB int64
		}{
			MinMemoryGB: RequiredGPUMemoryGB,
		},
		Bacalhau: struct {
			Options executor.BacalhauExecutorOptions
		}{
			Options: bacalhauOptions,
		},
	}

	err := checker.runAllChecks(ctx, config)
	if err != nil {
		log.Error().Err(err).Msg("❌ Preflight checks failed")
		return err
	}
	return nil
}

func (p *preflightChecker) runAllChecks(ctx context.Context, config preflightConfig) error {
	gpuResult := p.checkGPU(ctx, &gpuCheckConfig{
		minMemory: config.GPU.MinMemoryGB * 1024 * 1024 * 1024,
	})
	if !gpuResult.passed {
		return fmt.Errorf("GPU check failed: %s", gpuResult.message)
	}

	bacalhauResult := p.checkBacalhau(ctx, config.Bacalhau.Options)
	if !bacalhauResult.passed {
		return fmt.Errorf("Bacalhau check failed: %s", bacalhauResult.message)
	}

	if config.Docker.CheckRuntime {
		runtimeResult := p.checkDockerRuntime(ctx)
		if !runtimeResult.passed {
			return fmt.Errorf("Docker runtime check failed: %s", runtimeResult.message)
		}
	}

	return nil
}
