package models

import (
	"time"

	"gorm.io/gorm"
)

// Constantes pour le type de changement de points
const (
	PointTypeEarned = "earned"
	PointTypeSpent  = "spent"
)

type PointHistory struct {
	gorm.Model
	UserID        uint `gorm:"not null"`
	Points        int  `gorm:"default:0"`
	TransactionID uint
	EventID       uint
	ChangeDate    time.Time
	Type          string `gorm:"size:20;not null;check:type IN ('earned', 'spent')"`
}
