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

func parseGPURecord(record string) (*GPUInfo, error) {
	fields := strings.Split(record, ", ")
	if len(fields) != 4 {
		return nil, fmt.Errorf("invalid record format: expected 4 fields, got %d", len(fields))
	}

	// Parse memory, handling potential empty fields
	memoryParts := strings.Split(strings.TrimSpace(fields[2]), " ")
	if len(memoryParts) != 2 {
		return nil, fmt.Errorf("invalid memory format: %s", fields[2])
	}

	memoryStr := memoryParts[0]
	if memoryStr == "" {
		return nil, fmt.Errorf("empty memory value")
	}

	memoryMiB, err := strconv.ParseInt(memoryStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse memory value '%s': %w", memoryStr, err)
	}

	// Create GPU info with trimmed fields and validated memory
	gpu := &GPUInfo{
		UUID:          strings.TrimSpace(fields[0]),
		Name:          strings.TrimSpace(fields[1]),
		MemoryTotal:   memoryMiB,
		DriverVersion: strings.TrimSpace(fields[3]),
	}

	// Validate required fields
	if gpu.UUID == "" {
		return nil, fmt.Errorf("empty UUID")
	}
	if gpu.Name == "" {
		return nil, fmt.Errorf("empty Name")
	}
	if gpu.DriverVersion == "" {
		return nil, fmt.Errorf("empty DriverVersion")
	}

	return gpu, nil
}

func (p *preflightChecker) GetGPUInfo(ctx context.Context) ([]GPUInfo, error) {
	if err := checkNvidiaSMI(); err != nil {
		return nil, fmt.Errorf("nvidia-smi not available: %w", err)
	}

	cmd := exec.CommandContext(ctx, "nvidia-smi",
		"--query-gpu=gpu_uuid,gpu_name,memory.total,driver_version",
		"--format=csv,noheader")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Error().Str("output", string(output)).Err(err).Msg("nvidia-smi command failed")
		return nil, fmt.Errorf("error running nvidia-smi: %w", err)
	}

	records := strings.Split(strings.TrimSpace(string(output)), "\n")
	gpus := make([]GPUInfo, 0, len(records))

	for i, record := range records {
		gpu, err := parseGPURecord(record)
		if err != nil {
			log.Warn().Err(err).Int("index", i).Msg("Failed to parse GPU record")
			continue
		}

		gpus = append(gpus, *gpu)
		log.Info().
			Str("name", gpu.Name).
			Str("uuid", gpu.UUID).
			Int64("memory_mb", gpu.MemoryTotal).
			Msgf("🎮 GPU %d details", len(gpus))
	}

	if len(gpus) == 0 {
		return nil, fmt.Errorf("no valid GPUs found in nvidia-smi output")
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

	log.Info().Msg("✅ GPU requirements satisfied")
	return CheckResult{
		Passed:  true,
		Message: fmt.Sprintf("Found %d suitable GPUs", len(gpus)),
	}
}
