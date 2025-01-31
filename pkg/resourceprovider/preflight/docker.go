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

func (p *preflightChecker) CheckDockerRuntime(ctx context.Context) CheckResult {
	cmd := exec.CommandContext(ctx, "docker", "info", "--format", "{{json .}}")
	output, err := cmd.Output()
	if err != nil {
		return CheckResult{
			Passed:  false,
			Error:   fmt.Errorf("failed to get Docker info: %w", err),
			Message: "Docker check failed",
		}
	}

	var info dockerInfo
	if err := json.Unmarshal(output, &info); err != nil {
		return CheckResult{
			Passed:  false,
			Error:   fmt.Errorf("failed to parse Docker info: %w", err),
			Message: "Docker info parsing failed",
		}
	}

	// Check for nvidia runtime
	_, hasNvidia := info.Runtimes["nvidia"]
	if !hasNvidia {
		return CheckResult{
			Passed:  false,
			Error:   fmt.Errorf("nvidia runtime not found in Docker configuration"),
			Message: "NVIDIA runtime not found in Docker",
		}
	}

	// Test nvidia runtime
	testCmd := exec.CommandContext(ctx, "docker", "run", "--rm", "--runtime=nvidia", "nvidia/cuda:11.8.0-base", "nvidia-smi")
	if err := testCmd.Run(); err != nil {
		return CheckResult{
			Passed:  false,
			Error:   fmt.Errorf("failed to run NVIDIA runtime test: %w", err),
			Message: "NVIDIA runtime test failed",
		}
	}

	return CheckResult{
		Passed:  true,
		Message: "NVIDIA runtime is available and functional",
	}
}
