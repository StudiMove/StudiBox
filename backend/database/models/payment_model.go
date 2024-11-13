package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	PaymentStatusPending   = "pending"
	PaymentStatusConfirmed = "confirmed"
	PaymentStatusCancelled = "cancelled"
)

type Payment struct {
	gorm.Model
	UserID            uint   `gorm:"not null"`
	Amount            int    `gorm:"default:0"`
	Status            string `gorm:"size:20;default:'pending';check:status IN ('pending', 'confirmed', 'cancelled')"`
	PaymentDate       time.Time
	InstallmentNumber int `gorm:"default:1"`
	TotalInstallments int `gorm:"default:1"`
	CancellationDate  time.Time
	Transaction       PaymentTransaction `gorm:"foreignKey:PaymentID;constraint:OnDelete:CASCADE"`
}
