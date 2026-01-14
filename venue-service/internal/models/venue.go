package models

import (
	"gorm.io/gorm"
)

type VenueType string

const (
	VenueFootball   VenueType = "football"
	VenueBasketball VenueType = "basketball"
	VenueTennis     VenueType = "tennis"
	VenueGym        VenueType = "gym"
	VenueSwimming   VenueType = "swimming"
)

type Venue struct {
	gorm.Model
	VenueType VenueType `gorm:"column:venue_type;type:varchar(50);not null;check:venue_type IN ('football','basketball','tennis','gym','swimming')" json:"venue_type"`
	OwnerID   uint      `gorm:"column:owner_id;not null;index" json:"owner_id"`
	IsActive  bool      `gorm:"column:is_active;default:true" json:"is_active"`
	HourPrice int       `gorm:"column:hour_price;not null;check:hour_price >= 0" json:"hour_price"`
}

func (Venue) TableName() string {
	return "venues"
}
