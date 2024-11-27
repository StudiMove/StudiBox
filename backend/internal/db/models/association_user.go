package models

import (
	"gorm.io/gorm"
)

type AssociationUser struct {
	gorm.Model
	UserID          uint   `gorm:"uniqueIndex" json:"userId"`     // Référence unique à l'utilisateur
	User            User   `gorm:"foreignKey:UserID" json:"user"` // Relation avec le modèle User
	AssociationName string `gorm:"not null" json:"name"`
	SIRET           string `json:"siret"`
	Address         string `json:"address"`
	City            string `json:"city"`
	Postcode        string `json:"postcode"`
	Region          string `json:"region"`
	Country         string `json:"country"`
	Description     string `json:"description"`
	Status          string `json:"status" gorm:"default:En Attente" `
	IsValidated     bool   `gorm:"default:false" json:"isValidated"` // Validé par un admin
	IsActivated     bool   `gorm:"default:false" json:"isActivated"` // Activé par un admin
	IsPending       bool   `gorm:"default:true" json:"isPending"`    // En attente d'approbation
	MemberCount     int    `gorm:"default:0" json:"memberCount"`     // Compte des membres dans l'association
}
