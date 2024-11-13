package models

import (
	"time"

	"gorm.io/gorm"
)

// Constantes pour le statut du Ticket
const (
	TicketStatusValid     = "valid"
	TicketStatusCancelled = "cancelled"
	TicketStatusUsed      = "used"
	TicketStatusWaiting   = "waiting"
)

type Ticket struct {
	gorm.Model
	UserID     uint `gorm:"not null"`
	EventID    uint `gorm:"not null"`
	IssueDate  time.Time
	TicketCode string `gorm:"size:50;unique;not null"`
	Status     string `gorm:"size:20;default:'valid';check:status IN ('valid', 'cancelled', 'used', 'waiting')"`
}
