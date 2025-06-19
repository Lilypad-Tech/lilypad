package options

import (
	"github.com/Lilypad-Tech/lilypad/v2/pkg/anura"
)

func GetDefaultAnuraOptions() anura.AnuraClientOptions {
	return anura.AnuraClientOptions{
		BaseURL: GetDefaultServeOptionString("ANURA_BASE_URL", "https://anura-testnet.lilypad.tech/api/v1"),
		ApiKey:  GetDefaultServeOptionString("ANURA_API_KEY", ""),
	}
}
