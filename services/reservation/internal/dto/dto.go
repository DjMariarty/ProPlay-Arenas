package dto

import "time"

type ReservationCreate struct {
	ClientID uint   `json:"client_id"`
	OwnerID  uint   `json:"owner_id"`
	StartAt  time.Time `json:"start_at"`
	EndAt    time.Time `json:"end_at"`
	Price    int    `json:"price_cents,omitempty"`
	Status   string `json:"status"`
}