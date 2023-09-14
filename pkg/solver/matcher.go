package solver

import (
	"sort"
	"strings"

	"github.com/bacalhau-project/lilypad/pkg/data"
	"github.com/bacalhau-project/lilypad/pkg/solver/store"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/rs/zerolog/log"
)

type ListOfResourceOffers []data.ResourceOffer

func (a ListOfResourceOffers) Len() int { return len(a) }
func (a ListOfResourceOffers) Less(i, j int) bool {
	return a[i].DefaultPricing.InstructionPrice < a[j].DefaultPricing.InstructionPrice
}
func (a ListOfResourceOffers) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// the most basic of matchers
// basically just check if the resource offer >= job offer cpu, gpu & ram
// if the job offer is zero then it will match any resource offer
func doOffersMatch(
	resourceOffer data.ResourceOffer,
	jobOffer data.JobOffer,
) bool {
	if resourceOffer.Spec.CPU < jobOffer.Spec.CPU {
		log.Trace().
			Str("resource offer", resourceOffer.ID).
			Str("job offer", jobOffer.ID).
			Int("resource CPU", resourceOffer.Spec.CPU).
			Int("job CPU", jobOffer.Spec.CPU).
			Msgf("did not match CPU")
		return false
	}
	if resourceOffer.Spec.GPU < jobOffer.Spec.GPU {
		log.Trace().
			Str("resource offer", resourceOffer.ID).
			Str("job offer", jobOffer.ID).
			Int("resource GPU", resourceOffer.Spec.GPU).
			Int("job GPU", jobOffer.Spec.GPU).
			Msgf("did not match GPU")
		return false
	}
	if resourceOffer.Spec.RAM < jobOffer.Spec.RAM {
		log.Trace().
			Str("resource offer", resourceOffer.ID).
			Str("job offer", jobOffer.ID).
			Int("resource RAM", resourceOffer.Spec.RAM).
			Int("job RAM", jobOffer.Spec.RAM).
			Msgf("did not match RAM")
		return false
	}

	// if the resource provider has specified modules then check them
	if len(resourceOffer.Modules) > 0 {
		moduleID, err := data.GetModuleID(jobOffer.Module)
		if err != nil {
			log.Error().
				Err(err).
				Msgf("error getting module ID")
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
			log.Trace().
				Str("resource offer", resourceOffer.ID).
				Str("job offer", jobOffer.ID).
				Str("modules", strings.Join(resourceOffer.Modules, ", ")).
				Msgf("did not match modules")
			return false
		}
	}

	// we don't currently support market priced resource offers
	if resourceOffer.Mode == data.MarketPrice {
		log.Trace().
			Str("resource offer", resourceOffer.ID).
			Str("job offer", jobOffer.ID).
			Msgf("do not support market priced resource offers")
		return false
	}

	// if both are fixed price then we filter out "cannot afford"
	if resourceOffer.Mode == data.FixedPrice && jobOffer.Mode == data.FixedPrice {
		if resourceOffer.DefaultPricing.InstructionPrice > jobOffer.Pricing.InstructionPrice {
			log.Trace().
				Str("resource offer", resourceOffer.ID).
				Str("job offer", jobOffer.ID).
				Msgf("fixed price job offer cannot afford resource offer")
			return false
		}
	}

	mutualMediators := data.GetMutualServices(resourceOffer.Services.Mediator, jobOffer.Services.Mediator)
	if len(mutualMediators) == 0 {
		log.Trace().
			Str("resource offer", resourceOffer.ID).
			Str("job offer", jobOffer.ID).
			Msgf("no matching mutual mediators")
		return false
	}

	if resourceOffer.Services.Solver != jobOffer.Services.Solver {
		log.Trace().
			Str("resource offer", resourceOffer.ID).
			Str("job offer", jobOffer.ID).
			Msgf("no matching solver")
		return false
	}

	return true
}

func getMatchingDeals(
	db store.SolverStore,
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

			if doOffersMatch(resourceOffer.ResourceOffer, jobOffer.JobOffer) {
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
