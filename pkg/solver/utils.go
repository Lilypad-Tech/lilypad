package solver

import (
	"fmt"

	"github.com/bacalhau-project/lilypad/pkg/data"
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

// the most basic of matchers
// basically just check if the resource offer >= job offer cpu, gpu & ram
// if the job offer is zero then it will match any resource offer
func doOffersMatch(
	resourceOffer data.ResourceOffer,
	jobOffer data.JobOffer,
) bool {
	if resourceOffer.Spec.CPU < jobOffer.Spec.CPU {
		return false
	}
	if resourceOffer.Spec.GPU < jobOffer.Spec.GPU {
		return false
	}
	if resourceOffer.Spec.RAM < jobOffer.Spec.RAM {
		return false
	}

	// if the resource provider has specified modules then check them
	if len(resourceOffer.Modules) > 0 {
		moduleID, err := data.GetModuleID(jobOffer.Module)
		if err != nil {
			return false
		}
		// if the resourceOffer.Modules array does not contain the moduleID then we don't match
		hasModule := false
		for _, module := range resourceOffer.Modules {
			if module == moduleID {
				hasModule = true
				break
			}
		}

		if !hasModule {
			return false
		}
	}

	// if both are fixed price then we filter out "cannot afford"
	if resourceOffer.Mode == data.FixedPrice && jobOffer.Mode == data.FixedPrice {
		if resourceOffer.DefaultPricing.InstructionPrice > jobOffer.Pricing.InstructionPrice {
			return false
		}
	}

	return true
}

func getMatch(
	resourceOffer data.ResourceOffer,
	jobOffer data.JobOffer,
) data.Match {

	// deal := data.Deal{}

	return data.Match{}
}
