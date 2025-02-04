package preflight

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
)

type dockerInfo struct {
	Runtimes map[string]interface{} `json:"Runtimes"`
}

func (p *preflightChecker) checkDockerRuntime(ctx context.Context) checkResult {
	cmd := exec.CommandContext(ctx, "docker", "info", "--format", "{{json .}}")
	output, err := cmd.Output()
	if err != nil {
		return checkResult{
			passed:  false,
			error:   fmt.Errorf("failed to get Docker info: %w", err),
			message: "Docker check failed",
		}
	}

	var info dockerInfo
	if err := json.Unmarshal(output, &info); err != nil {
		return checkResult{
			passed:  false,
			error:   fmt.Errorf("failed to parse Docker info: %w", err),
			message: "Docker info parsing failed",
		}
	}

	// Check for nvidia runtime
	_, hasNvidia := info.Runtimes["nvidia"]
	if !hasNvidia {
		return checkResult{
			passed:  false,
			error:   fmt.Errorf("nvidia runtime not found in Docker configuration"),
			message: "NVIDIA runtime not found in Docker",
		}
	}

	// Test nvidia runtime
	testCmd := exec.CommandContext(ctx, "docker", "run", "--rm", "--runtime=nvidia", "nvidia/cuda:11.8.0-base", "nvidia-smi")
	if err := testCmd.Run(); err != nil {
		return checkResult{
			passed:  false,
			error:   fmt.Errorf("failed to run NVIDIA runtime test: %w", err),
			message: "NVIDIA runtime test failed",
		}
	}

	return checkResult{
		passed:  true,
		message: "NVIDIA runtime is available and functional",
	}
}
