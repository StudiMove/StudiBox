package models

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	UUID                 string `gorm:"type:uuid;uniqueIndex"`
	UserID               uint
	User                 User `gorm:"foreignKey:UserID"`
	EventID              uint
	Event                Event  `gorm:"foreignKey:EventID"`
	TicketNumberReadable string `gorm:"unique"`
	Status               string
	Tarifs               []TicketEventTarif  `gorm:"foreignKey:TicketID"`
	Options              []TicketEventOption `gorm:"foreignKey:TicketID"`
}
