package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	OwnerID      uint
	OwnerType    string `gorm:"size:20;not null;check:owner_type IN ('association', 'owner', 'school')"`
	ImageURL     string `gorm:"size:255"`
	VideoURL     string `gorm:"size:255"`
	Title        string `gorm:"size:100;not null"`
	Subtitle     string `gorm:"size:255"`
	StartDate    time.Time
	EndDate      time.Time
	StartTime    time.Time
	EndTime      time.Time
	Address      string `gorm:"size:100"`
	City         string `gorm:"size:100"`
	PostalCode   int32  `gorm:"type:int4"`
	Region       string `gorm:"size:50"`
	Country      string `gorm:"size:50"`
	IsOnline     bool   `gorm:"default:false"`
	IsPublic     bool   `gorm:"default:false"`
	IsValidated  bool   `gorm:"default:false"`
	UseStudibox  bool
	CategoryIDs  pq.Int64Array      `gorm:"type:integer[]"`
	TagIDs       pq.Int64Array      `gorm:"type:integer[]"`
	Categories   []EventCategory    `gorm:"many2many:event_event_categories;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tags         []EventTag         `gorm:"many2many:event_event_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Descriptions []EventDescription `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Options      []EventOption      `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Tarifs       []EventTarif       `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Tickets      []Ticket           `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE"`
}
