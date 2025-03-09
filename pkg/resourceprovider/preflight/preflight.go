package preflight

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

const RequiredGPUMemoryGB = 1      // 1GB of VRAM is required to startup if GPU is enabled
const RequiredNetworkSpeedMBps = 1 // 25 MB/s is the minimum required network speed

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
	Network struct {
		CheckSpeed   bool
		MinSpeedMBps int
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
		Network: struct {
			CheckSpeed   bool
			MinSpeedMBps int
		}{
			CheckSpeed:   true, // Always check network speed
			MinSpeedMBps: RequiredNetworkSpeedMBps,
		},
	}

	// Try to get GPU info, but continue with checks regardless
	gpuInfo, err := checker.getGPUInfo(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("⚠️  No GPU detected - will operate in CPU-only mode")
		// Continue with empty gpuInfo - don't return early
	} else {
		checker.gpuInfo = gpuInfo // Store GPU info in the checker
		log.Info().
			Int("gpu_count", len(gpuInfo)).
			Int64("min_memory_gb", config.GPU.MinMemoryGB).
			Msg("🛠️ GPU requirements")
	}

	// Run all checks regardless of GPU availability
	err = checker.runAllChecks(ctx, config)
	if err != nil {
		log.Error().Err(err).Msg("❌ Preflight checks failed")
		return err
	}
	return nil
}

func (p *preflightChecker) runAllChecks(ctx context.Context, config preflightConfig) error {
	// Network Speed Check - Run always and enforce the requirement
	if config.Network.CheckSpeed {
		networkResult := p.checkNetworkSpeed(ctx, &speedTestConfig{
			required:       true,
			minBandwidthMB: config.Network.MinSpeedMBps,
			testURL:        SpeedTestURL,
		})
		if !networkResult.passed {
			return fmt.Errorf("Network speed check failed: %s", networkResult.message)
		}
		log.Info().Msg(networkResult.message)
	}

	// GPU Check
	gpuResult := p.checkGPU(ctx, &gpuCheckConfig{
		minMemory: config.GPU.MinMemoryGB * 1024 * 1024 * 1024,
	})
	if !gpuResult.passed {
		return fmt.Errorf("GPU check failed: %s", gpuResult.message)
	}

	// Docker Runtime Check
	if config.Docker.CheckRuntime {
		runtimeResult := p.checkDockerRuntime(ctx)
		if !runtimeResult.passed {
			return fmt.Errorf("Docker runtime check failed: %s", runtimeResult.message)
		}
	}

	// Bacalhau Check - Run regardless of GPU availability
	if config.Bacalhau.CheckGPUAccess {
		bacalhauResult := p.checkBacalhauGPUAccess(ctx)
		if !bacalhauResult.passed {
			return fmt.Errorf("Bacalhau GPU access check failed: %s", bacalhauResult.message)
		}
		log.Info().Msg(bacalhauResult.message)
	}

	return nil
}
