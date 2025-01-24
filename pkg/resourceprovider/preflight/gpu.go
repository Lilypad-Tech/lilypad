package preflight

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

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

type nvidiaSmiResponse struct {
	UUID          string
	Name          string
	MemoryTotal   string // Set to a string since CSV doesn't parse numbers
	DriverVersion string
}

// GetGPUInfo runs nvidia-smi to get information about available GPUs
func (p *preflightChecker) GetGPUInfo(ctx context.Context) ([]GPUInfo, error) {
	log.Info().Msg("Attempting to run nvidia-smi check")

	if err := checkNvidiaSMI(); err != nil {
		return nil, err
	}

	log.Info().Msg("Running nvidia-smi command")
	cmd := exec.CommandContext(ctx, "nvidia-smi",
		"--query-gpu=gpu_uuid,gpu_name,memory.total,driver_version",
		"--format=csv,noheader")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Error().Str("output", string(output)).Err(err).Msg("nvidia-smi command failed")
		return nil, fmt.Errorf("error running nvidia-smi: %w", err)
	}

	// Parse CSV output
	records := strings.Split(strings.TrimSpace(string(output)), "\n")
	gpus := make([]GPUInfo, 0)

	for _, record := range records {
		fields := strings.Split(record, ", ")
		if len(fields) != 4 {
			continue
		}

		// Parse memory string (e.g. "12282 MiB" -> 12282)
		memoryStr := strings.Split(fields[2], " ")[0]
		memoryMiB, _ := strconv.ParseInt(memoryStr, 10, 64)

		gpu := GPUInfo{
			UUID:          strings.TrimSpace(fields[0]),
			Name:          strings.TrimSpace(fields[1]),
			MemoryTotal:   memoryMiB,
			DriverVersion: strings.TrimSpace(fields[3]),
		}
		gpus = append(gpus, gpu)

		log.Info().
			Str("name", gpu.Name).
			Str("uuid", gpu.UUID).
			Int64("memory_mb", gpu.MemoryTotal).
			Msgf("🎮 GPU %d details", len(gpus))
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
