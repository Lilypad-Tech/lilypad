package store

import (
	"github.com/lilypad-tech/lilypad/pkg/data"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type JobOffer struct {
	gorm.Model
	CID        string `gorm:"index"`
	JobCreator string `gorm:"index"`
	DealID     string `gorm:"index"`
	State      uint8
	Attributes datatypes.JSONType[data.JobOfferContainer]
}

type ResourceOffer struct {
	gorm.Model
	CID              string `gorm:"index"`
	ResourceProvider string `gorm:"index"`
	DealID           string `gorm:"index"`
	State            uint8
	Attributes       datatypes.JSONType[data.ResourceOfferContainer]
}

type Deal struct {
	gorm.Model
	CID              string `gorm:"index"`
	JobCreator       string `gorm:"index"`
	ResourceProvider string `gorm:"index"`
	Mediator         string
	State            uint8
	Attributes       datatypes.JSONType[data.DealContainer]
}

type Result struct {
	gorm.Model
	DealID     string `gorm:"index"` // We query with deal ID for now
	CID        string
	Attributes datatypes.JSONType[data.Result]
}

type MatchDecision struct {
	gorm.Model
	ResourceOffer string `gorm:"primaryKey"`
	JobOffer      string `gorm:"primaryKey"`
	Attributes    datatypes.JSONType[data.MatchDecision]
}
