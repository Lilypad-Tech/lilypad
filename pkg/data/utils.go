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

func GetJobOfferID(offer JobOffer) (string, error) {
	offer.ID = ""
	return CalculateCID(offer)
}

func GetResourceOfferID(offer ResourceOffer) (string, error) {
	offer.ID = ""
	return CalculateCID(offer)
}

func GetDealID(deal Deal) (string, error) {
	deal.ID = ""
	return CalculateCID(deal)
}

func GetModuleID(module ModuleConfig) (string, error) {
	return CalculateCID(module)
}
