package models

import "gorm.io/gorm"

// BusinessUser représente un utilisateur professionnel dans le système
type BusinessUser struct {
	gorm.Model
	UserID       uint   `gorm:"primaryKey"`        // Référence à l'utilisateur
	User         User   `gorm:"foreignKey:UserID"` // Relation avec le modèle User
	CompanyName  string `gorm:"not null"`          // Nom de l'entreprise
	SIRET        string // Numéro SIRET de l'entreprise
	Address      string // Adresse de l'entreprise
	City         string // Ville
	Postcode     string // Code postal
	Region       string // Région
	Country      string // Pays
	CreationDate string // Date de création de l'entreprise
	IsValidated  bool   `gorm:"default:false"` // Si l'entreprise est validée par un admin
	IsActivated  bool   `gorm:"default:false"` // Si le profil est activé
	CreatedAt    string `gorm:"not null"`      // Date de création
	UpdatedAt    string `gorm:"not null"`      // Date de mise à jour
}
