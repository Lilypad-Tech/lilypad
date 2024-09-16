package matcher

import (
	"sort"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/solver/store"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/rs/zerolog/log"
)

type ListOfResourceOffers []data.ResourceOffer

func (a ListOfResourceOffers) Len() int { return len(a) }
func (a ListOfResourceOffers) Less(i, j int) bool {
	return a[i].DefaultPricing.InstructionPrice < a[j].DefaultPricing.InstructionPrice
}
func (a ListOfResourceOffers) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func GetMatchingDeals(
	db store.SolverStore,
	updateJobOfferState func(string, string, uint8) (*data.JobOfferContainer, error),
) ([]data.Deal, error) {
	deals := []data.Deal{}

	resourceOffers, err := db.GetResourceOffers(store.GetResourceOffersQuery{
		NotMatched: true,
	})
	if err != nil {
		return nil, err
	}

	jobOffers, err := db.GetJobOffers(store.GetJobOffersQuery{
		NotMatched: true,
	})
	if err != nil {
		return nil, err
	}

	// loop over job offers
	for _, jobOffer := range jobOffers {

		// See if our jobOffer targets a specific address. If so, we will create a deal automatically
		// with the matcing resourceOffer.
		if jobOffer.JobOffer.Target.Address != "" {
			resourceOffer, err := db.GetResourceOfferByAddress(jobOffer.JobOffer.Target.Address)
			if err != nil {
				return nil, err
			}

			// We don't have a resource provider for this address
			if resourceOffer == nil {
				log.Trace().
					Str("job offer", jobOffer.ID).
					Str("target address", jobOffer.JobOffer.Target.Address).
					Msgf("No resource provider found for address")
				updateJobOfferState(jobOffer.ID, "", data.GetAgreementStateIndex("JobOfferCancelled"))
				continue
			}

			deal, err := data.GetDeal(jobOffer.JobOffer, resourceOffer.ResourceOffer)
			if err != nil {
				return nil, err
			}
			deals = append(deals, deal)
			continue
		}

		// loop over resource offers
		matchingResourceOffers := []data.ResourceOffer{}
		for _, resourceOffer := range resourceOffers {
			decision, err := db.GetMatchDecision(resourceOffer.ID, jobOffer.ID)
			if err != nil {
				return nil, err
			}

			// if this exists it means we've already tried to match the two elements and should not try again
			if decision != nil {
				continue
			}

			result := matchOffers(resourceOffer.ResourceOffer, jobOffer.JobOffer)
			logMatch(result)

			if result.matched() {
				matchingResourceOffers = append(matchingResourceOffers, resourceOffer.ResourceOffer)
			} else {
				_, err := db.AddMatchDecision(resourceOffer.ID, jobOffer.ID, "", false)
				if err != nil {
					return nil, err
				}
			}
		}

		// yay - we've got some matching resource offers
		// let's choose the cheapest one
		if len(matchingResourceOffers) > 0 {
			// now let's order the matching resource offers by price
			sort.Sort(ListOfResourceOffers(matchingResourceOffers))

			cheapestResourceOffer := matchingResourceOffers[0]
			deal, err := data.GetDeal(jobOffer.JobOffer, cheapestResourceOffer)
			if err != nil {
				return nil, err
			}

			// add the match decision for this job offer
			for _, matchingResourceOffer := range matchingResourceOffers {

				addDealID := ""
				if cheapestResourceOffer.ID == matchingResourceOffer.ID {
					addDealID = deal.ID
				}

				_, err := db.AddMatchDecision(matchingResourceOffer.ID, jobOffer.ID, addDealID, true)
				if err != nil {
					return nil, err
				}
			}

			deals = append(deals, deal)
		}
	}

	log.Debug().
		Int("jobOffers", len(jobOffers)).
		Int("resourceOffers", len(resourceOffers)).
		Int("deals", len(deals)).
		Msgf(system.GetServiceString(system.SolverService, "Solver solving"))

	return deals, nil
}
