package solver

import (
	"sort"
	"strings"
	"fmt"

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

type AllowlistItem struct {
	Enabled  bool
	ModuleID string
	Version  string
}

func loadAllowlist(filepath string) ([]AllowlistItem, error) {
		file, err := os.Open(filepath)
		if err != nil {
			return nil, err
		}
		defer file.Close()
	
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
	
		var allowlist []AllowlistItem
		if err := json.Unmarshal(bytes, &allowlist); err != nil {
			return nil, err
		}
	
		return allowlist, nil
	}

	// strip out https- keep the org and module in place
	// https://lilypad.tech/cowsay -> lilypad.tech/cowsay
	// Assume module ID has the shape of the org/repo(eg lilypad-tech/cowsay) if string
	func moduleMatch(moduleID, allowlistModuleID string) bool {
		return moduleID == allowlistModuleID
	}
	
	func versionMatch(version, allowlistVersion string) bool {
		if strings.Contains(allowlistVersion, "*") {
			prefix := strings.TrimSuffix(allowlistVersion, "*")
			return strings.HasPrefix(version, prefix)
		}
		return version == allowlistVersion
	}

	
// the most basic of matchers
// basically just check if the resource offer >= job offer cpu, gpu & ram

/*// if the job offer is zero then it will match any resource offer
func doOffersMatch(
	resourceOffer data.ResourceOffer,
	jobOffer data.JobOffer,
) bool */
func doOffersMatch(resourceOffer data.ResourceOffer, jobOffer data.JobOffer, allowlist []AllowlistItem, allowlistEnabled bool) bool {
	if !allowlistEnabled {
		log.Info().Msg("Allowlist is disabled. Skipping allowlist checks.")
		return true // Assuming that other checks will still be processed if needed
	}

	// Initial resource checks
	if resourceOffer.Spec.CPU < jobOffer.Spec.CPU || resourceOffer.Spec.GPU < jobOffer.Spec.GPU || resourceOffer.Spec.RAM < jobOffer.Spec.RAM {
		log.Trace().
			Str("resource offer", resourceOffer.ID).
			Str("job offer", jobOffer.ID).
			Msg("Resource does not meet job specifications")
		return false
	}

	// Check against the allowlist
	var allDisabled, allEnabled, foundMatch bool
	allDisabled, allEnabled = true, true
	for _, item := range allowlist {
		if item.Enabled {
			allDisabled = false
		} else {
			allEnabled = false
		}
		if moduleMatch(jobOffer.Module, item.ModuleID) && versionMatch(jobOffer.Version, item.Version) && item.Enabled {
			foundMatch = true
		}
	}

	// Evaluate allowlist conditions
	if allDisabled && foundMatch {
		log.Info().Msg("All items are disabled, but found a match.")
		return false
	}
	if allEnabled && !foundMatch {
		log.Info().Msg("All items are enabled, but no match was found.")
		return false
	}
	if !foundMatch { // General case for mixed list or no match found
		log.Info().Msg("No match found in mixed allowlist condition.")
		return false
	}

	// Additional checks for pricing, mediators, and solvers
	if resourceOffer.Mode == data.MarketPrice || (resourceOffer.Mode == data.FixedPrice && jobOffer.Mode == data.FixedPrice && resourceOffer.DefaultPricing.InstructionPrice > jobOffer.Pricing.InstructionPrice) {
		return false
	}
	mutualMediators := data.GetMutualServices(resourceOffer.Services.Mediator, jobOffer.Services.Mediator)
	if len(mutualMediators) == 0 || resourceOffer.Services.Solver != jobOffer.Services.Solver {
		return false
	}
	if len(resourceOffer.Modules) > 0 { 

		// Need to check against the new schema for allowlist aka "resourceOffer.Modules" 

		moduleID, err := data.GetModuleID(jobOffer.Module)
		if err != nil {
			log.Error().
				Err(err).
				Msgf("error getting module ID")
			return false
		}
	return true
}}
	// if the resource provider has specified to to run allowlist checker on modules then check them
	
	/*
		// if the resourceOffer.Modules array does not contain the moduleID then we don't match
		hasModule := false
		for _, module := range resourceOffer.Modules {
			if module == moduleID {
			hasModule = true
			break
		}
		
	*/
		// We need to check if the ModuleID matches but also if the version matches the module id and the job offer and make sure we have a match
		// If enabled false, match everything 

// LOGAN: Allowlist only has enabled false items, matches everything, except the prohibited items

/*
if allowlist.Every(item=>item.enabled==false) {
	if allowlist.find(item = > moduleMatch (item.ModuleID) and versionMatch(item.Version)) {
		// this means the module item is disabled
	}
	// if you have reached this point, there is a match
}
*/
func getMatchingDeals(db store.SolverStore, allowlistFilePath string, allowlistEnabled bool) ([]data.Deal, error) {
	allowlist, err := loadAllowlist(allowlistFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load allowlist: %w", err)
	}

	resourceOffers, err := db.GetResourceOffers(store.GetResourceOffersQuery{NotMatched: true})
	if err != nil {
		return nil, err
	}

	jobOffers, err := db.GetJobOffers(store.GetJobOffersQuery{NotMatched: true})
	if err != nil {
		return nil, err
	}

	var deals []data.Deal
	for _, jobOffer := range jobOffers {
		var matchingResourceOffers []data.ResourceOffer
		for _, resourceOffer := range resourceOffers {
			if doOffersMatch(resourceOffer, jobOffer, allowlist, allowlistEnabled) {
				matchingResourceOffers = append(matchingResourceOffers, resourceOffer)
			}
		}

		if len(matchingResourceOffers) > 0 {
			sort.Sort(ListOfResourceOffers(matchingResourceOffers))
			cheapestOffer := matchingResourceOffers[0]
			deal, err := data.GetDeal(jobOffer, cheapestOffer)
			if err != nil {
				return nil, err
			}
			deals = append(deals, deal)
		}
	}

	log.Debug().Int("jobOffers", len(jobOffers)).Int("resourceOffers", len(resourceOffers)).Int("deals", len(deals)).Msg("Deals processed.")
	return deals, nil
}

/*

// LOGAN: Allowlist only has enabled true items, matches nothing, except the allowed items

if allowlist.Every(item=>item.enabled==true) {
	if allowlist.find(item = > moduleMatch (item.ModuleID) and versionMatch(item.Version)) {
		// this means the module should be executed
	}
	// if you have reached this point, there is no match no deal
}
*/

//func versionMatch(version, allowedVersion string) bool {
    
//    return version == allowedVersion
//}





/*
// LOGAN: Allowlist has both, you dont care about disabled items, only the allowed items can go through, matches nothing
if allowlist.find(item = > moduleMatch (item.ModuleID) and versionMatch(item.Version)) {
	// this means the module should be executed
	
}
*/
func doOffersMatch(resourceOffer data.ResourceOffer, jobOffer data.JobOffer, allowlist []data.AllowlistEntry) bool {
    // Initial resource checks
    if resourceOffer.Spec.CPU < jobOffer.Spec.CPU || resourceOffer.Spec.GPU < jobOffer.Spec.GPU || resourceOffer.Spec.RAM < jobOffer.Spec.RAM {
        log.Trace().
            Str("resource offer", resourceOffer.ID).
            Str("job offer", jobOffer.ID).
            Msg("Resource does not meet job specifications")
        return false
    }

		// 
    // Check against the allowlist
    var allDisabled, allEnabled, foundMatch bool
    allDisabled, allEnabled = true, true
    for _, item := range allowlist {
        if item.Enabled {
            allDisabled = false
        } else {
            allEnabled = false
        }
        if item.ModuleID == jobOffer.Module && item.Enabled { // Simple module check for demonstration
            foundMatch = true
        }
    }

    // Evaluate allowlist conditions
    if allDisabled && foundMatch {
        log.Info().Msg("All items are disabled, but found a match.")
        return false
    }
    if allEnabled && !foundMatch {
        log.Info().Msg("All items are enabled, but no match was found.")
        return false
    }
    if !foundMatch { // General case for mixed list or no match found
        log.Info().Msg("No match found in mixed allowlist condition.")
        return false
    }

    // Additional checks for pricing, mediators, and solvers
    if resourceOffer.Mode == data.MarketPrice || (resourceOffer.Mode == data.FixedPrice && jobOffer.Mode == data.FixedPrice && resourceOffer.DefaultPricing.InstructionPrice > jobOffer.Pricing.InstructionPrice) {
        return false
    }
    mutualMediators := data.GetMutualServices(resourceOffer.Services.Mediator, jobOffer.Services.Mediator)
    if len(mutualMediators) == 0 || resourceOffer.Services.Solver != jobOffer.Services.Solver {
        return false
    }

    return true
}

func getMatchingDeals(db store.SolverStore) ([]data.Deal, error) {
    // Fetch offers that haven't been matched yet
    resourceOffers, err := db.GetResourceOffers(store.GetResourceOffersQuery{NotMatched: true})
    if err != nil {
        return nil, err
    }

    jobOffers, err := db.GetJobOffers(store.GetJobOffersQuery{NotMatched: true})
    if err != nil {
        return nil, err
    }

    var deals []data.Deal
    // Loop over job offers and resource offers to find matches
    for _, jobOffer := range jobOffers {
        var matchingResourceOffers []data.ResourceOffer
        for _, resourceOffer := range resourceOffers {
            if doOffersMatch(resourceOffer, jobOffer, nil /* pass actual allowlist here */) {
                matchingResourceOffers = append(matchingResourceOffers, resourceOffer)
            }
        }

        // Find the cheapest matching offer
        if len(matchingResourceOffers) > 0 {
            sort.Sort(ListOfResourceOffers(matchingResourceOffers))
            cheapestOffer := matchingResourceOffers[0]
            deal, err := data.GetDeal(jobOffer, cheapestOffer)
            if err != nil {
                return nil, err
            }
            deals = append(deals, deal)
        }
    }
    return deals, nil
}



		// the way we want to run this is lilypad resource-provider --offer-gpu=0 --enabled-modules=./modules.json
		/* 
		{ ModuleID: lilypad.tech/cowsay, version: 0.0.4, enabled: true }
		{ ModuleID: lilypad.tech/sdxl, version: 0.0.4, enabled: false }
		{ ModuleID: lilypad.tech/sdv, version: 0.0.*, enabled: false }
		
		
		*/


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
// hey does the offer match- 
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

				// Here is where we check for the module allowlist

			 if jobOffer.JobOffer.Module (not in)  moduleAllowlist { return nil, err} 

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
