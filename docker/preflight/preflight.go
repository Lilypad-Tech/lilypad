package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] != "check-gpu" {
		fmt.Println("Usage: preflight check-gpu")
		os.Exit(1)
	}

	// Check nvidia-smi is available
	if _, err := exec.LookPath("nvidia-smi"); err != nil {
		fmt.Println("Error: nvidia-smi not found. Please ensure NVIDIA drivers are installed")
		os.Exit(1)
	}

	// Run nvidia-smi to check GPU availability
	cmd := exec.Command("nvidia-smi")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running nvidia-smi: %v\n", err)
		os.Exit(1)
	}

	// Check docker runtime
	cmd = exec.Command("docker", "info", "--format", "'{{.Runtimes}}'")
	runtimeOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error checking Docker runtimes: %v\n", err)
		os.Exit(1)
	}

	if !strings.Contains(string(runtimeOutput), "nvidia") {
		fmt.Println("Error: NVIDIA runtime not found in Docker")
		os.Exit(1)
	}

	fmt.Println("GPU Check Output:")
	fmt.Println(string(output))
	fmt.Println("NVIDIA runtime is available")
	fmt.Println("All preflight checks passed successfully")
}
