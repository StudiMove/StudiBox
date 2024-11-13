package models

import (
	"time"

	"gorm.io/gorm"
)

// Constantes pour ProfileType
const (
	ProfileTypeStudent    = "etudiant"
	ProfileTypeNonStudent = "non_etudiant"
)

type User struct {
	gorm.Model
	FirstName     string `gorm:"size:50;not null"`
	LastName      string `gorm:"size:50;not null"`
	Pseudo        string `gorm:"size:30;unique;not null"`
	Email         string `gorm:"unique;not null"`
	Password      string `gorm:"not null"`
	Phone         int    `gorm:"type:int"`
	ProfileImage  string `gorm:"size:255"`
	BirthDate     time.Time
	Country       string `gorm:"size:100"`
	Region        string `gorm:"size:50"`
	PostalCode    int32  `gorm:"type:int4"`
	City          string `gorm:"size:100"`
	Address       string `gorm:"size:255"`
	ProfileType   string `gorm:"size:20;not null;check:profile_type IN ('etudiant', 'non_etudiant')"`
	StudiboxCoins int    `gorm:"default:0"`
	OwnerID       *uint  `gorm:"default:null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	RoleID        uint   `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Role          Role   `gorm:"foreignKey:RoleID"`
}
