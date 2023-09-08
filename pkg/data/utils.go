package data

import (
	"encoding/json"

	mdag "github.com/ipfs/go-merkledag"
)

// CalculateCID takes an interface, serializes it to JSON, and returns its IPFS CID
func CalculateCID(v interface{}) (string, error) {
	// Serialize the struct to JSON
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	// Create a Dag Node from the JSON bytes
	node := mdag.NodeWithData(data)

	// Compute CID of the Dag Node
	c := node.Cid()

	return c.String(), nil
}
