package solver

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/system"
	"github.com/rs/zerolog/log"
)

// Solver storage for tar input files
const INPUT_ARCHIVES_DIR = "job-input-archives"
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

func GetInputArchivesPath(id string) string {
	return system.GetDataDir(filepath.Join(INPUT_ARCHIVES_DIR, id))
}

func GetDealsFilePath(id string) string {
	return system.GetDataDir(filepath.Join(FILES_DIR, id))
}

func EnsureInputsArchivePath(id string) (string, error) {
	return system.EnsureDataDir(filepath.Join(INPUT_ARCHIVES_DIR, id))
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

// Check if timestamp is recent within a diff before or after now
func isTimestampRecent(timestamp int, diff int) bool {
	before := time.Now().UnixMilli() - int64(diff)
	after := time.Now().UnixMilli() + int64(diff)
	if int64(timestamp) < before || int64(timestamp) > after {
		return false
	}
	return true
}
