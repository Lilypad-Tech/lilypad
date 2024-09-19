package resourceprovider

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"os/exec"

	"github.com/rs/zerolog/log"
)

//go:embed ipfs.tar.gz
var ipfsBinary []byte

func StartIpfs() error {
	tmpFile, err := os.CreateTemp("", "ipfs-*")
	if err != nil {
		log.Debug().
			Str("ipfs", "tempFile").
			Msgf("create temp file failed: %v", err)
		return err
	}
	// defer os.Remove(tmpFile.Name())
	if err := decompressAndExtract(ipfsBinary, tmpFile.Name()); err != nil {
		fmt.Println("decompressAndExtract failed", err)
		return err
	}
	tmpFile.Close()

	ipfsRepoPath := os.Getenv("IPFS_PATH")
	if ipfsRepoPath == "" {
		ipfsRepoPath = os.Getenv("HOME") + "/.ipfs"
	}

	if _, err := os.Stat(ipfsRepoPath); os.IsNotExist(err) {
		// Initialize IPFS
		initCmd := exec.Command(tmpFile.Name(), "init")
		initCmd.Stdout = os.Stdout
		initCmd.Stderr = os.Stderr
		if err := initCmd.Run(); err != nil {
			log.Debug().
				Str("ipfs", "initCmd.Run").
				Msgf("IPFS initialization failed: %v", err)
			return err
		}
		// Configure the IPFS gateway address
		configCmd := exec.Command(tmpFile.Name(), "config", "Addresses.Gateway", "/ip4/127.0.0.1/tcp/8090")
		configCmd.Stdout = os.Stdout
		configCmd.Stderr = os.Stderr
		if err := configCmd.Run(); err != nil {
			log.Debug().
				Str("ipfs", "configCmd.Run").
				Msgf("IPFS configuration failed: %v", err)
			return err
		}

	}

	
	ipfscmd := exec.Command(tmpFile.Name(), "daemon")

	println(ipfscmd.String())
	stdout, err := ipfscmd.StdoutPipe()
	if err != nil {
		log.Debug().
			Str("ipfs", "StdoutPipe").
			Msgf("failed to get stdout pipe: %v", err)
		return err
	}

	stderr, err := ipfscmd.StderrPipe()
	if err != nil {
		log.Debug().
			Str("ipfs", "StderrPipe").
			Msgf("failed to get stderr pipe: %v", err)
		return err
	}

	if err := ipfscmd.Start(); err != nil {
		log.Debug().
			Str("ipfs", "Start").
			Msgf("failed to start command: %v", err)
		return err
	}

	// Create a scanner to read stdout
	stdoutScanner := bufio.NewScanner(stdout)
	go func() {
		for stdoutScanner.Scan() {
			println(stdoutScanner.Text())
			log.Debug().
				Str("ipfs", "stdoutScanner").
				Msgf("stdout: %s", stdoutScanner.Text())
		}
	}()

	// Create a scanner to read stderr
	stderrScanner := bufio.NewScanner(stderr)
	go func() {
		for stderrScanner.Scan() {
			log.Debug().
				Str("ipfs", "stderrScanner").
				Msgf("stderr: %s", stderrScanner.Text())
		}
	}()

	if err := ipfscmd.Wait(); err != nil {
		log.Debug().
			Str("ipfs", "Wait").
			Msgf("command execution failed: %v", err)
		return err
	}

	return nil
}
