package solver

import (
	"sort"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/solver/store"
	"github.com/rs/zerolog/log"
)

type ListOfResourceOffers []data.ResourceOffer

func (a ListOfResourceOffers) Len() int { return len(a) }
func (a ListOfResourceOffers) Less(i, j int) bool {
	return a[i].DefaultPricing.InstructionPrice < a[j].DefaultPricing.InstructionPrice
}
func (a ListOfResourceOffers) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func getMatches(controller *SolverController) ([]data.Match, error) {
	resourceOffers, err := controller.store.GetResourceOffers(store.GetResourceOffersQuery{})
	if err != nil {
		return nil, err
	}

	jobOffers, err := controller.store.GetJobOffers(store.GetJobOffersQuery{})
	if err != nil {
		return nil, err
	}

	// matches, err := controller.store.GetMatches(store.GetMatchesQuery{})
	// if err != nil {
	// 	return err
	// }

	// get maps of the resource offers and job offers by ids so we can filter them out
	// as we go
	resourceOffersMap := map[string]data.ResourceOffer{}
	for _, resourceOffer := range resourceOffers {
		resourceOffersMap[resourceOffer.ID] = resourceOffer
	}

	jobOffersMap := map[string]data.JobOffer{}
	for _, jobOffer := range jobOffers {
		jobOffersMap[jobOffer.ID] = jobOffer
	}

	deals, err := controller.store.GetDeals(store.GetDealsQuery{})
	if err != nil {
		return nil, err
	}

	// get maps of deal resource offers and job offers
	// so we can filter them out as we go
	dealsResourceOffersMap := map[string]data.ResourceOffer{}
	dealsJobOffersMap := map[string]data.JobOffer{}

	for _, deal := range deals {
		dealsResourceOffersMap[deal.ResourceOffer.ID] = deal.ResourceOffer
		dealsJobOffersMap[deal.JobOffer.ID] = deal.JobOffer
	}

	log.Debug().
		Int("jobOffers", len(jobOffers)).
		Int("resourceOffers", len(resourceOffers)).
		Msgf("Solver solving")

	// loop over job offers
	for _, jobOffer := range jobOffers {

		// loop over resource offers
		matchingResourceOffers := []data.ResourceOffer{}
		for _, resourceOffer := range resourceOffers {
			if doOffersMatch(resourceOffer, jobOffer) {
				matchingResourceOffers = append(matchingResourceOffers, resourceOffer)
				break
			}
		}

		// yay - we've got some matching resource offers
		// let's choose the cheapest one
		if len(matchingResourceOffers) > 0 {
			// now let's order the matching resource offers by price
			sort.Sort(ListOfResourceOffers(matchingResourceOffers))
			// offer := matchingResourceOffers[0]

			// err := controller.match(jobOffer, offer)
			// if err != nil {
			// 	log.Error().Err(err).Msgf("error matching")
			// } else {
			// 	// we don't need to keep looping we have matched
			// 	break
			// }

		}
	}

	return nil, nil
}
