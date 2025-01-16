package preflight

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
)

type nvidiaSmiResponse struct {
	GPU []struct {
		UUID          string `json:"uuid"`
		ProductName   string `json:"product_name"`
		Memory        int64  `json:"memory_total"`
		DriverVersion string `json:"driver_version"`
	} `json:"gpu"`
}

func checkNvidiaSMI() error {
	_, err := exec.LookPath("nvidia-smi")
	return err
}

func (v *preflight) GetAvailableGPUs(ctx context.Context) ([]GPUInfo, error) {
	if err := checkNvidiaSMI(); err != nil {
		return nil, fmt.Errorf("nvidia-smi not found: %w", err)
	}

	cmd := exec.CommandContext(ctx, "nvidia-smi", "--query-gpu=gpu_uuid,gpu_name,memory.total,driver_version", "--format=json")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error running nvidia-smi: %w", err)
	}

	var response nvidiaSmiResponse
	if err := json.Unmarshal(output, &response); err != nil {
		return nil, fmt.Errorf("error parsing nvidia-smi output: %w", err)
	}

	gpus := make([]GPUInfo, len(response.GPU))
	for i, gpu := range response.GPU {
		gpus[i] = GPUInfo{
			UUID:          gpu.UUID,
			Name:          gpu.ProductName,
			MemoryTotal:   gpu.Memory,
			DriverVersion: gpu.DriverVersion,
		}
	}

	return gpus, nil
}

func (v *preflight) PreflightGPU(ctx context.Context, req *GPURequirements) PreflightResult {
	gpus, err := v.GetAvailableGPUs(ctx)
	if err != nil {
		return PreflightResult{
			Success: false,
			Error:   err,
			Details: "Failed to get GPU information",
		}
	}

	if len(gpus) < req.MinGPUs {
		return PreflightResult{
			Success: false,
			Details: fmt.Sprintf("Insufficient GPUs. Required: %d, Available: %d", req.MinGPUs, len(gpus)),
		}
	}

	// Additional validation logic for memory and capabilities can be added here

	return PreflightResult{
		Success: true,
		Details: fmt.Sprintf("Found %d suitable GPUs", len(gpus)),
	}
}
