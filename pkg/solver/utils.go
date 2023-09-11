package solver

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func LogSolverEvent(ev SolverEvent) {
	switch ev.EventType {
	case JobOfferAdded:
		log.Info().
			Str("solver event: JobOfferAdded", fmt.Sprintf("%+v", *ev.JobOffer)).
			Msgf("")
	case ResourceOfferAdded:
		log.Info().
			Str("solver event: ResourceOfferAdded", fmt.Sprintf("%+v", *ev.ResourceOffer)).
			Msgf("")
	case MatchFound:
		log.Info().
			Str("solver event: MatchFound", fmt.Sprintf("%+v", ev)).
			Msgf("")
	}
}
