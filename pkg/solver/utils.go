package solver

import (
	"fmt"
	"path/filepath"

	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/rs/zerolog/log"
)

const FILES_DIR = "job-files"
const DOWNLOADS_DIR = "downloaded-files"

func LogSolverEvent(badge string, ev SolverEvent) {
	switch ev.EventType {
	case JobOfferAdded:
		log.Debug().
			Str(fmt.Sprintf("%s -> JobOfferAdded", badge), fmt.Sprintf("%+v", *ev.JobOffer)).
			Msgf("")
	case ResourceOfferAdded:
		log.Debug().
			Str(fmt.Sprintf("%s -> ResourceOfferAdded", badge), fmt.Sprintf("%+v", *ev.ResourceOffer)).
			Msgf("")
	case DealAdded:
		log.Debug().
			Str(fmt.Sprintf("%s -> DealAdded", badge), fmt.Sprintf("%+v", ev)).
			Msgf("")
	case JobOfferStateUpdated:
		log.Debug().
			Str(fmt.Sprintf("%s -> JobOfferStateUpdated", badge), fmt.Sprintf("%+v", ev)).
			Msgf("")
	case ResourceOfferStateUpdated:
		log.Debug().
			Str(fmt.Sprintf("%s -> ResourceOfferStateUpdated", badge), fmt.Sprintf("%+v", ev)).
			Msgf("")
	case DealStateUpdated:
		log.Debug().
			Str(fmt.Sprintf("%s -> DealStateUpdated", badge), fmt.Sprintf("%+v", ev)).
			Msgf("")
	case ResourceProviderTransactionsUpdated:
		log.Debug().
			Str(fmt.Sprintf("%s -> ResourceProviderTransactionsUpdated", badge), fmt.Sprintf("%+v", ev)).
			Msgf("")
	case JobCreatorTransactionsUpdated:
		log.Debug().
			Str(fmt.Sprintf("%s -> JobCreatorTransactionsUpdated", badge), fmt.Sprintf("%+v", ev)).
			Msgf("")
	}
}

func ServiceLogSolverEvent(service system.Service, ev SolverEvent) {
	LogSolverEvent(system.GetServiceBadge(service), ev)
}

func getDealIDs(deals []data.Deal) []string {
	var ids []string
	for _, deal := range deals {
		ids = append(ids, deal.ID)
	}
	return ids
}

func getResourceOfferIDs(resourceOffers []data.ResourceOfferContainer) []string {
	var ids []string
	for _, offer := range resourceOffers {
		ids = append(ids, offer.ID)
	}
	return ids
}

func getJobOfferIDs(jobOffers []data.JobOfferContainer) []string {
	var ids []string
	for _, offer := range jobOffers {
		ids = append(ids, offer.ID)
	}
	return ids
}

func GetDealsFilePath(id string) string {
	return system.GetDataDir(filepath.Join(FILES_DIR, id))
}

func EnsureDealsFilePath(id string) (string, error) {
	return system.EnsureDataDir(filepath.Join(FILES_DIR, id))
}

func GetDownloadsFilePath(id string) string {
	return system.GetDataDir(filepath.Join(DOWNLOADS_DIR, id))
}

func EnsureDownloadsFilePath(id string) (string, error) {
	return system.EnsureDataDir(filepath.Join(DOWNLOADS_DIR, id))
}
