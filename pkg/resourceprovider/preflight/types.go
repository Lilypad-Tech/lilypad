package preflight

import (
	"context"
)

// ValidationResult represents the outcome of a validation check
type ValidationResult struct {
	Success bool
	Details string
	Error   error
}

// GPUInfo represents information about an available GPU
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

// GPURequirements defines what GPU capabilities are required
type GPURequirements struct {
	MinMemory    int64
	MinGPUs      int
	Capabilities []string
}

type PreflightChecker interface {
    // CheckGPU verifies GPU requirements
    CheckGPU(ctx context.Context, config *GPUCheckConfig) CheckResult
    // CheckDockerRuntime verifies Docker runtime configuration
    CheckDockerRuntime(ctx context.Context) CheckResult
    // GetGPUInfo retrieves information about available GPUs
    GetGPUInfo(ctx context.Context) ([]GPUInfo, error)
    // RunAllChecks performs all configured checks
    RunAllChecks(ctx context.Context, config PreflightConfig) error
}

// Validator defines the interface for performing validations
type Validator interface {
	// ValidateGPU checks if the required GPU capabilities are available
	ValidateGPU(ctx context.Context, req *GPURequirements) ValidationResult

	// ValidateDockerRuntime checks if the nvidia runtime is available
	ValidateDockerRuntime(ctx context.Context) ValidationResult

	// GetAvailableGPUs returns information about all available GPUs
	GetAvailableGPUs(ctx context.Context) ([]GPUInfo, error)
}
