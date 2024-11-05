//go:build !(linux && amd64)

package resourceprovider

var cardBinary []byte = []byte("just mock a string, make it buildable for other components")
