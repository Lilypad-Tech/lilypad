package preflight

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/rs/zerolog/log"
)

type nvidiaSmiResponse struct {
	GPU []struct {
		UUID          string `json:"uuid"`
		ProductName   string `json:"product_name"`
		Memory        int64  `json:"memory_total"`
		DriverVersion string `json:"driver_version"`
	} `json:"gpu"`
}

type preflightChecker struct {
	gpuInfo []GPUInfo
}

type GPUCheckConfig struct {
	Required     bool
	MinGPUs      int
	MinMemory    int64
	Capabilities []string
}

func checkNvidiaSMI() error {
	_, err := exec.LookPath("nvidia-smi")
	return err
}

func (p *preflightChecker) GetGPUInfo(ctx context.Context) ([]GPUInfo, error) {
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

func (p *preflightChecker) CheckGPU(ctx context.Context, config *GPUCheckConfig) CheckResult {
	if !config.Required {
		log.Info().Msg("GPU checks skipped - not required")
		return CheckResult{
			Passed:  true,
			Message: "GPU not required for this configuration",
		}
	}

	log.Info().Msg("Starting GPU preflight check")

	gpus, err := p.GetGPUInfo(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to get GPU information")
		// If GPUs are required but not found, return failure
		if config.Required {
			return CheckResult{
				Passed:  false,
				Error:   err,
				Message: "Failed to get GPU information",
			}
		}
		// If GPUs aren't required, pass even if we can't find them
		return CheckResult{
			Passed:  true,
			Message: "No GPUs found but none required",
		}
	}

	log.Info().
		Int("gpu_count", len(gpus)).
		Int("required_gpus", config.MinGPUs).
		Msg("Found GPUs")

	for i, gpu := range gpus {
		log.Info().
			Str("uuid", gpu.UUID).
			Str("name", gpu.Name).
			Int64("memory", gpu.MemoryTotal).
			Int("index", i).
			Msg("GPU details")
	}

	if len(gpus) < config.MinGPUs {
		log.Warn().
			Int("available", len(gpus)).
			Int("required", config.MinGPUs).
			Msg("Insufficient GPUs")
		return CheckResult{
			Passed:  false,
			Message: fmt.Sprintf("Insufficient GPUs. Required: %d, Available: %d", config.MinGPUs, len(gpus)),
		}
	}

	log.Info().Msg("GPU check passed successfully")
	return CheckResult{
		Passed:  true,
		Message: fmt.Sprintf("Found %d suitable GPUs", len(gpus)),
	}
}
