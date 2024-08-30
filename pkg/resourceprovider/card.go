package resourceprovider

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"

	"github.com/rs/zerolog/log"
)

//go:embed card
var cardBinary []byte

func PostCard() error {
	tmpFile, err := os.CreateTemp("", "card-*")
	if err != nil {
		log.Debug().
			Str("pow->event", "PowNewPowRound").
			Msgf("create temp file failed: %v", err)
		return err
	}
	defer os.Remove(tmpFile.Name())

	// Write the embedded binary to the temporary file
	if _, err := tmpFile.Write(cardBinary); err != nil {
		log.Debug().
			Str("pow->event", "PowNewPowRound").
			Msgf("write temp file failed: %v", err)
		return err
	}

	// Close the file to ensure all data is written
	if err := tmpFile.Close(); err != nil {
		log.Debug().
			Str("pow->event", "PowNewPowRound").
			Msgf("close temp file failed: %v", err)
		return err
	}

	// Make the temporary file executable
	if err := os.Chmod(tmpFile.Name(), 0755); err != nil {
		log.Debug().
			Str("pow->event", "PowNewPowRound").
			Msgf("chmod temp file failed: %v", err)
		return err
	}

	fmt.Print(tmpFile.Name())
	// Execute the temporary file
	cardcmd := exec.Command(tmpFile.Name())
	output1, err := cardcmd.Output()
	if err != nil {
		log.Debug().
			Str("pow->event", "PowNewPowRound").
			Msgf("execute temp file failed: %v", err)
		return err
	}

	// Print the output
	println(string(output1))
	return nil
}
