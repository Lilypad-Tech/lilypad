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
	MemoryTotal   string
	DriverVersion string
}

func (p *preflightChecker) GetGPUInfo(ctx context.Context) ([]GPUInfo, error) {
	log.Debug().Msg("Checking for nvidia-smi")

	if err := checkNvidiaSMI(); err != nil {
		log.Warn().Msg("⚠️  nvidia-smi not found - system appears to have no NVIDIA GPU")
		return nil, fmt.Errorf("nvidia-smi not available: %w", err)
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

	records := strings.Split(strings.TrimSpace(string(output)), "\n")
	gpus := make([]GPUInfo, 0)

	for _, record := range records {
		fields := strings.Split(record, ", ")
		if len(fields) != 4 {
			continue
		}

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
		// Attempt to retrieve GPU info
		gpus, err := p.GetGPUInfo(ctx)
		if err != nil {
			log.Warn().Msg("⚠️  Running without GPU support - Resource Provider will operate in CPU-only mode")
			return CheckResult{
				Passed:  true,
				Message: "Operating in CPU-only mode",
			}
		}

		// If we found GPUs, log them but still continue
		log.Info().Msgf("🎮 Found %d optional GPUs available for use", len(gpus))
		return CheckResult{
			Passed:  true,
			Message: fmt.Sprintf("Found %d NVIDIA GPUs (optional)", len(gpus)),
		}
	}

	// Required GPU checks
	log.Info().Msg("Starting required GPU checks")
	gpus, err := p.GetGPUInfo(ctx)
	if err != nil {
		return CheckResult{
			Passed:  false,
			Error:   err,
			Message: "Required GPU check failed - no NVIDIA GPUs detected",
		}
	}

	log.Info().
		Int("gpu_count", len(gpus)).
		Int("required_gpus", config.MinGPUs).
		Msg("Checking GPU requirements")

	if len(gpus) < config.MinGPUs {
		log.Warn().
			Int("available", len(gpus)).
			Int("required", config.MinGPUs).
			Msg("Insufficient GPUs for requirements")
		return CheckResult{
			Passed:  false,
			Message: fmt.Sprintf("Insufficient GPUs. Required: %d, Available: %d", config.MinGPUs, len(gpus)),
		}
	}

	log.Info().Msg("✅ GPU requirements satisfied")
	return CheckResult{
		Passed:  true,
		Message: fmt.Sprintf("Found %d suitable GPUs", len(gpus)),
	}
}
