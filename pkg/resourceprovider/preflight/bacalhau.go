package preflight

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/rs/zerolog/log"
)

type bacalhauInfo struct {
	Version    string
	APIVersion string
	NodeID     string
}

// checkBacalhauGPUAccess verifies that bacalhau CLI can access GPUs
func (p *preflightChecker) checkBacalhauGPUAccess(ctx context.Context) checkResult {
	// No need to check nvidia-smi again as we already have GPU info in the checker
	if len(p.gpuInfo) == 0 {
		return checkResult{
			passed:  true,
			message: "No GPUs detected, skipping Bacalhau GPU access check",
		}
	}

	for i, gpu := range p.gpuInfo {
		log.Info().
			Str("name", gpu.name).
			Str("uuid", gpu.uuid).
			Int64("memory_mb", gpu.memoryTotal).
			Msgf("🖥️ Verifying Bacalhau access to GPU %d", i+1)
	}

	// Get container name - in a real environment this might be configurable
	containerName := "bacalhau"

	// Check if Docker is installed
	_, err := exec.LookPath("docker")
	if err != nil {
		return checkResult{
			passed:  false,
			error:   fmt.Errorf("docker CLI not found: %w", err),
			message: "Docker CLI is not installed or not in PATH",
		}
	}

	// Check if the container exists
	cmdCheckContainer := exec.CommandContext(ctx, "docker", "ps", "-q", "-f", "name="+containerName)
	containerID, err := cmdCheckContainer.Output()
	if err != nil || len(containerID) == 0 {
		return checkResult{
			passed:  false,
			error:   fmt.Errorf("bacalhau container not found: %w", err),
			message: "Bacalhau container is not running",
		}
	}

	// Check if Bacalhau can connect to the network by listing nodes
	cmdList := exec.CommandContext(ctx, "docker", "exec", containerName, "bacalhau", "node", "list")
	output, err := cmdList.CombinedOutput()
	if err != nil {
		return checkResult{
			passed:  false,
			error:   fmt.Errorf("bacalhau network check failed: %w", err),
			message: "Unable to connect to Bacalhau network",
		}
	}

	// Look for GPU information in the output without logging the entire output
	nodeListInfo := string(output)
	if !containsGPUInfo(nodeListInfo) {
		return checkResult{
			passed:  false,
			error:   fmt.Errorf("no GPU information found in node list"),
			message: "Bacalhau nodes do not show GPU configuration",
		}
	}

	log.Info().
		Int("gpu_count", len(p.gpuInfo)).
		Msg("✅ Bacalhau GPU configuration verified")

	return checkResult{
		passed:  true,
		message: fmt.Sprintf("Bacalhau has access to %d GPUs", len(p.gpuInfo)),
	}
}

// containsGPUInfo checks if the output contains GPU-related configuration
func containsGPUInfo(nodeInfo string) bool {
	gpuIndicators := []string{
		"GPU",
		"NVIDIA",
		"nvidia.com/gpu",
		"compute-type=gpu",
		"GeForce",
		"RTX",
	}

	for _, indicator := range gpuIndicators {
		if strings.Contains(nodeInfo, indicator) {
			return true
		}
	}
	return false
}
