package models

import (
	"time"
)

type Status string

const (
	pending   Status = "pending"
	confirmed Status = "confirmed"
	cancelled Status = "cancelled"
	completed Status = "completed"
)

type ReservationDetails struct {
	Base
	ClientID uint          `json:"client_id"`
	OwnerID  uint          `json:"owner_id"`
	StartAt  time.Time     `json:"start_at" gorm:"not null"`
	EndAt    time.Time     `json:"end_at" gorm:"not null"`
	Price    float64       `json:"price_cents,omitempty"`
	Duration time.Duration `json:"duration_minutes,omitempty"`

	Status Status `json:"status"`
}

type Reservation struct {
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
	Status  Status    `json:"status"`
}
