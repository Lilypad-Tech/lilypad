package store

import (
	"github.com/Lilypad-Tech/lilypad/v2/pkg/data"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type JobOffer struct {
	gorm.Model `gorm:"type:timestamp(3) with time zone"`
	CID        string `gorm:"index"`
	JobCreator string `gorm:"index"`
	DealID     string `gorm:"index"`
	State      uint8
	Attributes datatypes.JSONType[data.JobOfferContainer]
}

type ResourceOffer struct {
	gorm.Model       `gorm:"type:timestamp(3) with time zone"`
	CID              string `gorm:"index"`
	ResourceProvider string `gorm:"index"`
	DealID           string `gorm:"index"`
	State            uint8
	Attributes       datatypes.JSONType[data.ResourceOfferContainer]
}

type Deal struct {
	gorm.Model       `gorm:"type:timestamp(3) with time zone"`
	CID              string `gorm:"index"`
	JobCreator       string `gorm:"index"`
	ResourceProvider string `gorm:"index"`
	Mediator         string
	State            uint8
	Attributes       datatypes.JSONType[data.DealContainer]
}

type Result struct {
	gorm.Model `gorm:"type:timestamp(3) with time zone"`
	DealID     string `gorm:"index"` // We query with deal ID for now
	CID        string
	Attributes datatypes.JSONType[data.Result]
}

type MatchDecision struct {
	gorm.Model    `gorm:"type:timestamp(3) with time zone"`
	ResourceOffer string `gorm:"primaryKey;index:idx_resource_offer_job_offer,priority:1"`
	JobOffer      string `gorm:"primaryKey;index:idx_resource_offer_job_offer,priority:2"`
	Attributes    datatypes.JSONType[data.MatchDecision]
}

type AllowedResourceProvider struct {
	gorm.Model
	ResourceProvider string `gorm:"index"`
}
