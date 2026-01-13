package models

import (
	"time"
)

type status string

const (
	pending   status = "pending"
	confirmed status = "confirmed"
	cancelled status = "cancelled"
	completed status = "completed"
)

type ReservationDetails struct {
	Base
	ClientID    uint      `json:"client_id"`
	OwnerID     uint      `json:"owner_id"`  
	StartAt     time.Time `json:"start_at" gorm:"not null"`
	EndAt       time.Time `json:"end_at" gorm:"not null"`
	Price       float64   `json:"price_cents,omitempty"`
	Paid        bool      `json:"paid"`
	IsAvailable bool      `json:"is_available"`
	Status      status    `json:"status"`
}

type Reservation struct {
	ClientID uint   `json:"client_id"`
	StartAt  time.Time  `json:"start_at"`
	EndAt    time.Time  `json:"end_at"`
}
