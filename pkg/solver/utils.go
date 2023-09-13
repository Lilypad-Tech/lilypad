package solver

import (
	"fmt"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/rs/zerolog/log"
)

func LogSolverEvent(badge string, ev SolverEvent) {
	switch ev.EventType {
	case JobOfferAdded:
		log.Info().
			Str(fmt.Sprintf("%s -> JobOfferAdded", badge), fmt.Sprintf("%+v", *ev.JobOffer)).
			Msgf("")
	case ResourceOfferAdded:
		log.Info().
			Str(fmt.Sprintf("%s -> ResourceOfferAdded", badge), fmt.Sprintf("%+v", *ev.ResourceOffer)).
			Msgf("")
	case DealAdded:
		log.Info().
			Str(fmt.Sprintf("%s -> DealAdded", badge), fmt.Sprintf("%+v", ev)).
			Msgf("")
	case JobOfferStateUpdated:
		log.Info().
			Str(fmt.Sprintf("%s -> JobOfferStateUpdated", badge), fmt.Sprintf("%+v", ev)).
			Msgf("")
	case ResourceOfferStateUpdated:
		log.Info().
			Str(fmt.Sprintf("%s -> ResourceOfferStateUpdated", badge), fmt.Sprintf("%+v", ev)).
			Msgf("")
	case DealStateUpdated:
		log.Info().
			Str(fmt.Sprintf("%s -> DealStateUpdated", badge), fmt.Sprintf("%+v", ev)).
			Msgf("")
	case ResourceProviderTransactionsUpdated:
		log.Info().
			Str(fmt.Sprintf("%s -> ResourceProviderTransactionsUpdated", badge), fmt.Sprintf("%+v", ev)).
			Msgf("")
	case JobCreatorTransactionsUpdated:
		log.Info().
			Str(fmt.Sprintf("%s -> JobCreatorTransactionsUpdated", badge), fmt.Sprintf("%+v", ev)).
			Msgf("")
	}
}

func ServiceLogSolverEvent(service system.Service, ev SolverEvent) {
	LogSolverEvent(system.GetServiceBadge(service), ev)
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
) data.JobOfferContainer {
	return data.JobOfferContainer{
		ID:         jobOffer.ID,
		DealID:     "",
		JobCreator: jobOffer.JobCreator,
		State:      data.GetDefaultAgreementState(),
		JobOffer:   jobOffer,
	}
}

func getResourceOfferContainer(
	resourceOffer data.ResourceOffer,
) data.ResourceOfferContainer {
	return data.ResourceOfferContainer{
		ID:               resourceOffer.ID,
		DealID:           "",
		ResourceProvider: resourceOffer.ResourceProvider,
		State:            data.GetDefaultAgreementState(),
		ResourceOffer:    resourceOffer,
	}
}

func getDealContainer(
	deal data.Deal,
) data.DealContainer {
	return data.DealContainer{
		ID:               deal.ID,
		JobCreator:       deal.JobOffer.JobCreator,
		ResourceProvider: deal.ResourceOffer.ResourceProvider,
		JobOffer:         deal.JobOffer.ID,
		ResourceOffer:    deal.ResourceOffer.ID,
		State:            data.GetDefaultAgreementState(),
		Deal:             deal,
	}
}

func checkResourceOffer(resourceOffer data.ResourceOffer) error {
	if resourceOffer.Mode == data.MarketPrice {
		return fmt.Errorf("resource offer mode cannot be market price")
	}

	if len(resourceOffer.TrustedParties.Directory) <= 0 {
		return fmt.Errorf("resource offer must have at least one trusted directory")
	}

	if len(resourceOffer.TrustedParties.Mediator) <= 0 {
		return fmt.Errorf("resource offer must have at least one trusted mediator")
	}

	return nil
}

func checkJobOffer(jobOffer data.JobOffer) error {
	if len(jobOffer.TrustedParties.Directory) <= 0 {
		return fmt.Errorf("job offer must have at least one trusted directory")
	}

	if len(jobOffer.TrustedParties.Mediator) <= 0 {
		return fmt.Errorf("job offer must have at least one trusted mediator")
	}

	return nil
}
