package resourceprovider

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/lilypad-tech/lilypad/pkg/powLogs"
	"github.com/rs/zerolog/log"
)

//go:embed card
var cardBinary []byte

func PostCard(id string, challange string, difficulty string) error {
	tmpFile, err := os.CreateTemp("", "card-*")
	if err != nil {
		log.Error().
			Msgf("create temp file failed: %v", err)
		return err
	}
	defer os.Remove(tmpFile.Name())

	// Write the embedded binary to the temporary file
	if _, err := tmpFile.Write(cardBinary); err != nil {
		log.Error().
			Msgf("write temp file failed: %v", err)
		return err
	}

	// Close the file to ensure all data is written
	if err := tmpFile.Close(); err != nil {
		log.Error().
			Msgf("close temp file failed: %v", err)
		return err
	}

	// Make the temporary file executable
	if err := os.Chmod(tmpFile.Name(), 0755); err != nil {
		log.Error().
			Msgf("chmod temp file failed: %v", err)
		return err
	}

	fmt.Print(tmpFile.Name())
	// Execute the temporary file
	cardcmd := exec.Command(tmpFile.Name(), id, challange, difficulty)
	output1, err := cardcmd.Output()
	if err != nil {
		log.Error().
			Msgf("execute temp file failed: %v", err)
		return err
	}

	// Print the output
	println(string(output1))

	cardInfo := make(map[string]string)
	err = json.Unmarshal(output1, &cardInfo)
	if err != nil {
		log.Debug().
			Msgf("unmarshal cardinfo failed: %v", err)
		return err
	}

	powLogs.TrackCardInfo(cardInfo)

	return nil
}
