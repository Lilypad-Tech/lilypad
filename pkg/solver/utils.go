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

func getMutualTrustedParties(a []string, b []string) []string {
	mutual := []string{}
	for _, aParty := range a {
		for _, bParty := range b {
			if aParty == bParty {
				mutual = append(mutual, aParty)
			}
		}
	}
	return mutual
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

	mutualMediators := getMutualTrustedParties(resourceOffer.TrustedParties.Mediator, jobOffer.TrustedParties.Mediator)
	if len(mutualMediators) == 0 {
		return false
	}

	mutualDirectories := getMutualTrustedParties(resourceOffer.TrustedParties.Directory, jobOffer.TrustedParties.Directory)
	if len(mutualDirectories) == 0 {
		return false
	}

	return true
}

func getDeal(
	jobOffer data.JobOffer,
	resourceOffer data.ResourceOffer,
) (data.Deal, error) {
	mutualMediators := getMutualTrustedParties(resourceOffer.TrustedParties.Mediator, jobOffer.TrustedParties.Mediator)
	mutualDirectories := getMutualTrustedParties(resourceOffer.TrustedParties.Directory, jobOffer.TrustedParties.Directory)

	dealData := data.Deal{
		Members: data.DealMembers{
			JobCreator:       jobOffer.JobCreator,
			ResourceProvider: resourceOffer.ResourceProvider,
			Directory:        mutualDirectories[0],
			Mediators:        mutualMediators,
		},
		// TODO: this assumes marketing pricing for the client
		// this should be configurable
		Pricing: resourceOffer.DefaultPricing,
		// TODO: this assumes resource provider timeouts
		// this should be configurable
		Timeouts:      resourceOffer.DefaultTimeouts,
		JobOffer:      jobOffer,
		ResourceOffer: resourceOffer,
	}

	id, err := data.GetDealID(dealData)

	if err != nil {
		return dealData, err
	}

	dealData.ID = id
	return dealData, nil
}

func getJobOfferContainer(
	jobOffer data.JobOffer,
) (data.JobOfferContainer, error) {
	state, err := data.GetAgreementState("DealNegotiating")
	if err != nil {
		return data.JobOfferContainer{}, err
	}
	return data.JobOfferContainer{
		ID:         jobOffer.ID,
		DealID:     "",
		JobCreator: jobOffer.JobCreator,
		State:      state,
		JobOffer:   jobOffer,
	}, nil
}

func getResourceOfferContainer(
	resourceOffer data.ResourceOffer,
) (data.ResourceOfferContainer, error) {
	state, err := data.GetAgreementState("DealNegotiating")
	if err != nil {
		return data.ResourceOfferContainer{}, err
	}
	return data.ResourceOfferContainer{
		ID:               resourceOffer.ID,
		DealID:           "",
		ResourceProvider: resourceOffer.ResourceProvider,
		State:            state,
		ResourceOffer:    resourceOffer,
	}, nil
}

func getDealContainer(
	deal data.Deal,
) (data.DealContainer, error) {
	state, err := data.GetAgreementState("DealNegotiating")
	if err != nil {
		return data.DealContainer{}, err
	}
	return data.DealContainer{
		ID:               deal.ID,
		JobCreator:       deal.JobOffer.JobCreator,
		ResourceProvider: deal.ResourceOffer.ResourceProvider,
		JobOffer:         deal.JobOffer.ID,
		ResourceOffer:    deal.ResourceOffer.ID,
		State:            state,
		Deal:             deal,
	}, nil
}
