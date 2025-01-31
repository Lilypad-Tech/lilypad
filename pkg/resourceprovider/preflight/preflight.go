package preflight

import (
	"context"
	"fmt"
)

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

type PreflightChecker interface {
	CheckGPU(ctx context.Context, config *GPUCheckConfig) CheckResult
	CheckDockerRuntime(ctx context.Context) CheckResult
	GetGPUInfo(ctx context.Context) ([]GPUInfo, error)
	RunAllChecks(ctx context.Context, config PreflightConfig) error
}

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
