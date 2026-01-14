package models

import "time"

type VenueType string

const (
	VenueFootball   VenueType = "football"
	VenueBasketball VenueType = "basketball"
	VenueTennis     VenueType = "tennis"
	VenueGym        VenueType = "gym"
	VenueSwimming   VenueType = "swimming"
)

type Venue struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Type      VenueType `gorm:"type:varchar(50);not null" json:"type"`
	Owner     uint      `gorm:"not null;index" json:"owner"`
	Status    bool      `gorm:"default:true" json:"status"`
	HourPrice int       `gorm:"not null;check:hour_price >= 0" json:"hour_price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Venue) TableName() string {
	return "venues"
}
