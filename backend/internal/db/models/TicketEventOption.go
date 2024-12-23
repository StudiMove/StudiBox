package models

import "gorm.io/gorm"

type TicketEventOption struct {
	gorm.Model
	TicketID uint        `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Référence au ticket
	OptionID uint        `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Référence à l'option
	Option   EventOption `gorm:"foreignKey:OptionID"`
}
