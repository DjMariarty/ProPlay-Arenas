package dto

import (
	"time"

	"payment-service/internal/models"
)

type RefundRequest struct {
	Amount int64  `json:"amount" binding:"required,gt=0"`
	Reason string `json:"reason" binding:"required,min=5,max=500"`
}

type RefundResponse struct {
	ID        uint                `json:"id"`
	PaymentID uint                `json:"payment_id"`
	Amount    int64               `json:"amount"`
	Reason    string              `json:"reason"`
	Status    models.RefundStatus `json:"status"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}
