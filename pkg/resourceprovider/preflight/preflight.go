package preflight

import (
	"context"
	"fmt"

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
		CheckGPUAccess bool
	}
}

type preflightChecker struct {
	gpuInfo []gpuInfo
}

func RunPreflightChecks() error {
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
			CheckGPUAccess bool
		}{
			CheckGPUAccess: true,
		},
	}

	// Logging GPU requirements
	gpuInfo, err := checker.getGPUInfo(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("⚠️  No GPU detected - will operate in CPU-only mode")
		return nil
	} else {
		checker.gpuInfo = gpuInfo // Store GPU info in the checker
		log.Info().
			Int("gpu_count", len(gpuInfo)).
			Int64("min_memory_gb", config.GPU.MinMemoryGB).
			Msg("🛠️ GPU requirements")
	}

	err = checker.runAllChecks(ctx, config)
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

	if config.Docker.CheckRuntime {
		runtimeResult := p.checkDockerRuntime(ctx)
		if !runtimeResult.passed {
			return fmt.Errorf("Docker runtime check failed: %s", runtimeResult.message)
		}
	}

	// Check Bacalhau GPU access if GPUs are detected and Bacalhau check is enabled

	if len(p.gpuInfo) > 0 && config.Bacalhau.CheckGPUAccess {
		bacalhauResult := p.checkBacalhauGPUAccess(ctx)
		if !bacalhauResult.passed {
			return fmt.Errorf("Bacalhau GPU access check failed: %s", bacalhauResult.message)
		}
		log.Info().Msg(bacalhauResult.message)
	}

	return nil
}
