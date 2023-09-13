package system

import "fmt"

type Service string

const (
	SolverService           Service = "solver"
	ResourceProviderService Service = "resource-provider"
	JobCreatorService       Service = "job-creator"
	DirectoryService        Service = "directory"
	MediatorService         Service = "mediator"
)

func GetServiceBadge(service Service) string {
	switch service {
	case SolverService:
		return "🔴 SOL"
	case ResourceProviderService:
		return "🔵 RP"
	case JobCreatorService:
		return "🟢 JC"
	case DirectoryService:
		return "🟡 DIR"
	case MediatorService:
		return "🟠 MED"
	default:
		return "⚪"
	}
}

func GetServiceString(service Service, st string) string {
	return fmt.Sprintf("%s %s", GetServiceBadge(service), st)
}
