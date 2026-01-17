package models

import (
	"fmt"
	"venue-service/internal/validation"

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

func (vt VenueType) IsValid() bool {
	switch vt {
	case VenueFootball, VenueBasketball, VenueTennis, VenueGym, VenueSwimming:
		return true
	default:
		return false
	}
}

func (vt VenueType) String() string {
	return string(vt)
}

type Venue struct {
	gorm.Model
	VenueType VenueType `json:"venue_type" gorm:"column:venue_type;type:varchar(50);not null"`
	OwnerID   uint      `json:"owner_id" gorm:"column:owner_id;not null;index"`
	IsActive  bool      `json:"is_active" gorm:"column:is_active;default:true"`
	HourPrice int       `json:"hour_price" gorm:"column:hour_price;not null;check:hour_price >= 0"`
	District  string    `json:"district" gorm:"column:district;type:varchar(50);not null"`
	StartTime string    `json:"start_time" gorm:"column:start_time;type:time;not null"` // Рабочее время начала (например "09:00")
	EndTime   string    `json:"end_time" gorm:"column:end_time;type:time;not null"`     // Рабочее время окончания (например "18:00")
}

func (Venue) TableName() string {
	return "venues"
}

// validateVenue проверяет валидность данных площадки
func (v *Venue) validateVenue() error {
	if !v.VenueType.IsValid() {
		return fmt.Errorf("неверный тип площадки: %s", v.VenueType)
	}

	// Валидация времени начала
	_, err := validation.ValidateTime(v.StartTime)
	if err != nil {
		return fmt.Errorf("время начала: %w", err)
	}

	// Валидация времени окончания
	_, err = validation.ValidateTime(v.EndTime)
	if err != nil {
		return fmt.Errorf("время окончания: %w", err)
	}

	// Проверка, что время начала раньше времени окончания
	isBefore, err := validation.CompareTimes(v.StartTime, v.EndTime)
	if err != nil {
		return err
	}
	if !isBefore {
		return fmt.Errorf("время начала должно быть раньше времени окончания")
	}

	return nil
}

func (v *Venue) BeforeCreate(tx *gorm.DB) error {
	return v.validateVenue()
}

func (v *Venue) BeforeUpdate(tx *gorm.DB) error {
	return v.validateVenue()
}
