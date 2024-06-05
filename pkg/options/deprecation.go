package options

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/web3"
)

// TODO(bgins) Remove once users have migrated away from old testnet
func CheckDeprecation(serviceOptions data.ServiceConfig, web3Options web3.Web3Options) {
	if web3Options.RpcURL == "ws://testnet.lilypad.tech:8546" {
		web3Options.PrivateKey = "*****"
		log.SetFlags(0)
		log.Fatal(fmt.Sprintf(`This testnet has been deprecated. Please remove all environment variables for RPC URL, chain ID, and contract addresses.

Your environment currently contains:

%v
%v
The new testnet environment will be set by default, and you will only need to set your private key.
`, spew.Sdump(web3Options), spew.Sdump(serviceOptions)))
	}
}
