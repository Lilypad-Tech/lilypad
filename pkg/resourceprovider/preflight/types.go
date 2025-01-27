package preflight

import (
	"context"
)

type GPUInfo struct {
	UUID          string
	Name          string
	MemoryTotal   int64
	DriverVersion string
	Capabilities  []string
}

type CheckResult struct {
	Passed  bool
	Message string
	Error   error
}

type PreflightConfig struct {
	GPU struct {
		Required     bool
		Enabled      bool
		MinMemoryGB  int64
		Capabilities []string
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
