package options

import "github.com/bacalhau-project/lilypad/pkg/mediator"

func NewMediatorOptions() mediator.MediatorOptions {
	return mediator.MediatorOptions{
		Web3: GetDefaultWeb3Options(),
	}
}
