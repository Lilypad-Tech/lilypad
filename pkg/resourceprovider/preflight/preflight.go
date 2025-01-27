package preflight

import (
	"context"
	"fmt"
)

// NewPreflightChecker creates a new instance of PreflightChecker
func NewPreflightChecker() PreflightChecker {
	return &preflightChecker{}
}

// RunAllChecks performs all configured preflight checks
func (p *preflightChecker) RunAllChecks(ctx context.Context, config PreflightConfig) error {
	if config.GPU.Enabled {
		gpuResult := p.CheckGPU(ctx, &GPUCheckConfig{
			MinMemory:    config.GPU.MinMemoryGB * 1024 * 1024 * 1024,
			Capabilities: config.GPU.Capabilities,
		})
		if !gpuResult.Passed {
			return fmt.Errorf("GPU check failed: %s", gpuResult.Message)
		}
	}

	if config.Docker.CheckRuntime {
		runtimeResult := p.CheckDockerRuntime(ctx)
		if !runtimeResult.Passed {
			return fmt.Errorf("Docker runtime check failed: %s", runtimeResult.Message)
		}
	}

	return nil
}
