package preflight

import (
	"context"
	"fmt"
)

// NewPreflightChecker creates a new instance of PreflightChecker
func NewPreflightChecker() PreflightChecker {
	return &preflightChecker{}
}

func (p *preflightChecker) RunAllChecks(ctx context.Context, config PreflightConfig) error {

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
