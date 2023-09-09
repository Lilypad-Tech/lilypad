package system

import (
	"os"
	"os/exec"

	"github.com/rs/zerolog/log"
)

func RunCommand(
	name string,
	args []string,
	workingDir string,
) error {
	log.Info().Msgf("Running %s %s", name, args)
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = workingDir
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
