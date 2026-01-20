package models

import (
	"gorm.io/gorm"
)

type RefundStatus string

const (
	RefundStatusPending   RefundStatus = "pending"
	RefundStatusCompleted RefundStatus = "completed"
	RefundStatusFailed    RefundStatus = "failed"
)

type Refund struct {
	gorm.Model
	PaymentID uint         `gorm:"index" json:"payment_id"`
	Amount    int64        `gorm:"column:amount" json:"amount"`
	Reason    string       `gorm:"column:reason;type:text" json:"reason"`
	Status    RefundStatus `gorm:"column:status" json:"status"`
	Payment   *Payment     `gorm:"foreignKey:PaymentID;references:ID" json:"payment,omitempty"`
}
