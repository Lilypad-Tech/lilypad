package preflight

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

type gpuCheckConfig struct {
	required     bool
	minGPUs      int
	minMemory    int64
	capabilities []string
}

func checkNvidiaSMI() error {
	_, err := exec.LookPath("nvidia-smi")
	return err
}

type nvidiaSmiResponse struct {
	uuid          string
	name          string
	memoryTotal   string
	driverVersion string
}

func parseGPURecord(record string) (*gpuInfo, error) {
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
	gpu := &gpuInfo{
		uuid:          strings.TrimSpace(fields[0]),
		name:          strings.TrimSpace(fields[1]),
		memoryTotal:   memoryMiB,
		driverVersion: strings.TrimSpace(fields[3]),
	}

	// Validate required fields
	if gpu.uuid == "" {
		return nil, fmt.Errorf("empty UUID")
	}
	if gpu.name == "" {
		return nil, fmt.Errorf("empty Name")
	}
	if gpu.driverVersion == "" {
		return nil, fmt.Errorf("empty DriverVersion")
	}

	return gpu, nil
}

func (p *preflightChecker) getGPUInfo(ctx context.Context) ([]gpuInfo, error) {
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
	gpus := make([]gpuInfo, 0, len(records))

	for _, record := range records {
		gpu, err := parseGPURecord(record)
		if err != nil {
			log.Warn().Err(err).Msgf("Failed to parse GPU record: %s", record)
			continue
		}

		gpus = append(gpus, *gpu)
		log.Info().
			Str("name", gpu.name).
			Str("uuid", gpu.uuid).
			Int64("memory_mb", gpu.memoryTotal).
			Msgf("🛠️ GPU %d details", len(gpus))
	}

	if len(gpus) == 0 {
		return nil, fmt.Errorf("no valid GPUs found in nvidia-smi output")
	}

	return gpus, nil
}

func (p *preflightChecker) checkGPU(ctx context.Context, config *gpuCheckConfig) checkResult {
	if !config.required {
		// Attempt to retrieve GPU info
		gpus, err := p.getGPUInfo(ctx)
		if err != nil {
			log.Warn().Msg("⚠️  Running without GPU support - Resource Provider will operate in CPU-only mode")
			return checkResult{
				passed:  true,
				message: "Operating in CPU-only mode",
			}
		}

		// If we found GPUs, log them but still continue
		log.Info().Msgf("🖥️ Found %d optional GPUs available for use", len(gpus))
		return checkResult{
			passed:  true,
			message: fmt.Sprintf("Found %d NVIDIA GPUs (optional)", len(gpus)),
		}
	}

	// Required GPU checks
	log.Info().Msg("Starting required GPU checks")
	gpus, err := p.getGPUInfo(ctx)
	if err != nil {
		return checkResult{
			passed:  false,
			error:   err,
			message: "Required GPU check failed - no NVIDIA GPUs detected",
		}
	}

	log.Info().Msg("✅ GPU requirements satisfied")
	return checkResult{
		passed:  true,
		message: fmt.Sprintf("Found %d suitable GPUs", len(gpus)),
	}
}
