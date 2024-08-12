package system

import (
	"fmt"

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
			gpu_info = append(
				gpu_info,
				fmt.Sprintf("%s %s (%s)", card.DeviceInfo.Vendor.Name, card.DeviceInfo.Product.Name, card.DeviceInfo.Driver),
			)
		}
	}

	return gpu_info
}
