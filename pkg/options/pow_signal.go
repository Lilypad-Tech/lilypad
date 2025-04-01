package options

import (
	"github.com/lilypad-tech/lilypad/v2/pkg/web3"
)

type PowSignalOptions struct {
	Web3 web3.Web3Options
}

func NewPowSignalOptions() PowSignalOptions {
	options := PowSignalOptions{
		Web3: GetDefaultWeb3Options(),
	}
	return options
}

func ProcessPowSignalOptions(options PowSignalOptions, network string) (PowSignalOptions, error) {
	newWeb3Options, err := ProcessWeb3Options(options.Web3, network)
	if err != nil {
		return options, err
	}
	options.Web3 = newWeb3Options
	return options, CheckPowSignalOptions(options)
}

func CheckPowSignalOptions(options PowSignalOptions) error {
	return CheckWeb3Options(options.Web3)
}
