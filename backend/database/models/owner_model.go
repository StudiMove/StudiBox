package models

import (
	"gorm.io/gorm"
)

// Déclaration des statuts possibles pour Owner et les types d'Owner
const (
	StatusPending        = "En Attente"
	StatusValidated      = "Validé"
	StatusRejected       = "Rejeté"
	StatusInactive       = "Inactif"
	OwnerTypeAssociation = "association"
	OwnerTypeOwner       = "owner"
	OwnerTypeSchool      = "school"
)

type Owner struct {
	gorm.Model
	User        User   `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CompanyName string `gorm:"size:100;not null"`
	Type        string `gorm:"size:20;not null;check:type IN ('association', 'owner', 'school')"`
	Phone       int    `gorm:"type:int"`
	SIRET       string `gorm:"size:20"`
	Country     string `gorm:"size:100"`
	Region      string `gorm:"size:50"`
	PostalCode  int32  `gorm:"type:int4"`
	City        string `gorm:"size:100"`
	Address     string `gorm:"size:255"`
	Description string `gorm:"type:text"`
	Status      string `gorm:"size:20;default:'En Attente';check:status IN ('En Attente', 'Validé', 'Rejeté', 'Inactif')"`
	MemberCount int    `gorm:"default:0"`
}
