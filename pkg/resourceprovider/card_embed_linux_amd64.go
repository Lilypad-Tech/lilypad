//go:build linux && amd64

package resourceprovider

import (
	_ "embed"
)

//go:embed card-linux-amd64
var cardBinary []byte
