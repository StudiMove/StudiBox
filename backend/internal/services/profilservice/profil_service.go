package profilservice

import (
	"backend/internal/services/storage"

	"gorm.io/gorm"
)

// ProfilService struct, commun à tous les services de profil
type ProfilService struct {
	DB             *gorm.DB
	storageService storage.StorageService
}

// NewProfilService crée une nouvelle instance de ProfilService avec DB et StorageService
func NewProfilService(db *gorm.DB, storageService storage.StorageService) *ProfilService {
	return &ProfilService{
		DB:             db,
		storageService: storageService,
	}
}
