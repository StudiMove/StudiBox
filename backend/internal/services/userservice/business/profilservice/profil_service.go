package profilservice

import "gorm.io/gorm"

// ProfilService struct, commun à tous les services
type ProfilService struct {
    DB *gorm.DB
}

// NewProfilService crée une nouvelle instance de ProfilService
func NewProfilService(db *gorm.DB) *ProfilService {
    return &ProfilService{DB: db}
}
