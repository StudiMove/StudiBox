package models

import "gorm.io/gorm"

type TicketEventTarif struct {
	gorm.Model
	TicketID uint       `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Référence au ticket
	TarifID  uint       `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Référence au tarif
	Tarif    EventTarif `gorm:"foreignKey:TarifID"`
}
