package preflight

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

const RequiredGPUMemoryGB = 1 // 1GB of VRAM is required to startup if GPU is enabled

type GPUInfo struct {
	UUID          string
	Name          string
	MemoryTotal   int64
	DriverVersion string
}

type CheckResult struct {
	Passed  bool
	Message string
	Error   error
}

type preflightConfig struct {
	GPU struct {
		MinMemoryGB int64
	}
	Docker struct {
		CheckRuntime bool
	}
}

type preflightChecker struct {
	gpuInfo []GPUInfo
}

func RunPreflightChecks(ctx context.Context) error {
	log.Info().Msg("Starting preflight checks...")
	checker := &preflightChecker{}
	config := preflightConfig{
		GPU: struct {
			MinMemoryGB int64
		}{
			MinMemoryGB: RequiredGPUMemoryGB,
		},
	}

	// Logging GPU requirements
	gpuInfo, err := checker.GetGPUInfo(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("⚠️  No GPU detected - will operate in CPU-only mode")
	} else {
		log.Info().
			Int("gpu_count", len(gpuInfo)).
			Int64("min_memory_gb", config.GPU.MinMemoryGB).
			Msg("🎮 GPU requirements")
	}

	err = checker.RunAllChecks(ctx, config)
	if err != nil {
		log.Error().Err(err).Msg("❌ Preflight checks failed")
		return err
	}
	return nil
}

func (p *preflightChecker) RunAllChecks(ctx context.Context, config preflightConfig) error {

	gpuResult := p.CheckGPU(ctx, &GPUCheckConfig{
		MinMemory: config.GPU.MinMemoryGB * 1024 * 1024 * 1024,
	})
	if !gpuResult.Passed {
		return fmt.Errorf("GPU check failed: %s", gpuResult.Message)
	}

	if config.Docker.CheckRuntime {
		runtimeResult := p.CheckDockerRuntime(ctx)
		if !runtimeResult.Passed {
			return fmt.Errorf("Docker runtime check failed: %s", runtimeResult.Message)
		}
	}

	return nil
}
