package options

import (
	"github.com/lilypad-tech/lilypad/pkg/resourceprovider/powsignal"
)

func NewPowSignalOptions() powsignal.PowSignalOptions {
	options := powsignal.PowSignalOptions{
		Web3: GetDefaultWeb3Options(),
	}
	return options
}

func ProcessPowSignalrOptions(options powsignal.PowSignalOptions, network string) (powsignal.PowSignalOptions, error) {
	newWeb3Options, err := ProcessWeb3Options(options.Web3, network)
	if err != nil {
		return options, err
	}
	options.Web3 = newWeb3Options
	return options, CheckPowSignalOptions(options)
}

func CheckPowSignalOptions(options powsignal.PowSignalOptions) error {
	return CheckWeb3Options(options.Web3)
}
