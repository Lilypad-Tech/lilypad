package resourceprovider

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"os/exec"

	"github.com/rs/zerolog/log"
)

//go:embed bacalhau.tar.gz
var bacalhauBinary []byte

func StartBacalhau() error {
	tmpFile, err := os.CreateTemp("", "bacalhau-*")
	if err != nil {
		log.Debug().
			Str("bacalhau", "CreateTemp").
			Msgf("create temp file failed: %v", err)
		return err
	}
	// defer os.Remove(tmpFile.Name())
	println(tmpFile.Name())
	if err := decompressAndExtract(bacalhauBinary, tmpFile.Name()); err != nil {
		fmt.Println("decompressAndExtract failed", err)
		return err
	}
	tmpFile.Close()

	// Execute the temporary file
	bacalhaucmd := exec.Command(tmpFile.Name(), "serve", "--node-type", "compute,requester", "--peer", "none", "--private-internal-ipfs=false", "--job-selection-accept-networked", "--ipfs-connect", "/ip4/127.0.0.1/tcp/5001")
	println(bacalhaucmd.String())
	stdout, err := bacalhaucmd.StdoutPipe()
	if err != nil {
		log.Debug().
			Str("bacalhau", "StdoutPipe").
			Msgf("failed to get stdout pipe: %v", err)
		return err
	}

	stderr, err := bacalhaucmd.StderrPipe()
	if err != nil {
		log.Debug().
			Str("bacalhau", "StderrPipe").
			Msgf("failed to get stderr pipe: %v", err)
		return err
	}

	if err := bacalhaucmd.Start(); err != nil {
		log.Debug().
			Str("bacalhau", "Start").
			Msgf("failed to start command: %v", err)
		return err
	}

	// Create a scanner to read stdout
	stdoutScanner := bufio.NewScanner(stdout)
	go func() {
		for stdoutScanner.Scan() {
			println(stdoutScanner.Text())
			log.Debug().
				Str("bacalhau", "ScanOut").
				Msgf("stdout: %s", stdoutScanner.Text())
		}
	}()

	// Create a scanner to read stderr
	stderrScanner := bufio.NewScanner(stderr)
	go func() {
		for stderrScanner.Scan() {
			err := stderrScanner.Text()
			if len(err) > 0 {
				log.Debug().
					Str("bacalhau", "log").
					Msgf("log: %s", err)
			}
		}
	}()

	if err := bacalhaucmd.Wait(); err != nil {
		log.Debug().
			Str("bacalhau", "Wait").
			Msgf("command execution failed: %v", err)
		return err
	}

	return nil
}
