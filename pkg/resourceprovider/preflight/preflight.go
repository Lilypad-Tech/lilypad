package preflight

import (
	"context"
	"fmt"
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

type PreflightConfig struct {
	GPU struct {
		MinMemoryGB int64
	}
	Docker struct {
		CheckRuntime bool
	}
}

type PreflightChecker struct {
	gpuInfo []GPUInfo
}

func (p *PreflightChecker) RunAllChecks(ctx context.Context, config PreflightConfig) error {

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
