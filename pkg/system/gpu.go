package system

import (
	"github.com/jaypipes/ghw"
	"github.com/rs/zerolog/log"
)

func GetGPUInfo() []string {
	gpu, err := ghw.GPU()
	if err != nil {
		log.Warn().Msgf("failed to get GPU info: %s", err)
	}

	var gpu_info []string
	if gpu != nil {
		for _, card := range gpu.GraphicsCards {
			gpu_info = append(gpu_info, card.String())
		}
	}

	return gpu_info
}
