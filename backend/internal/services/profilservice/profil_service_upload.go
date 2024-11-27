package profilservice

import (
	"backend/internal/api/models/upload/request"
	"backend/internal/db/models"
	"errors"
	"log"
)

// UploadProfileImage gère le téléchargement de l'image de profil
func (s *ProfilService) UploadProfileImage(req *request.ProfileImageUploadRequest) (string, error) {
	log.Printf("Starting file upload for file: %s", req.FileName)

	// Utilisation de storageService pour télécharger l'image
	url, err := s.storageService.UploadFile(req.File, req.FileName)
	if err != nil {
		log.Printf("Error while uploading file to storage service: %v", err)
		return "", errors.New("failed to upload file to storage")
	}

	log.Printf("File uploaded successfully. URL: %s", url)
	return url, nil
}

// UpdateUserProfileImage met à jour l'URL de l'image de profil dans la base de données
func (s *ProfilService) UpdateUserProfileImage(userID uint, imageURL string) error {
	var user models.User
	if err := s.DB.Model(&user).Where("id = ?", userID).Update("profile_image", imageURL).Error; err != nil {
		log.Printf("Error updating profile image URL in the database: %v", err)
		return errors.New("failed to update profile image URL in the database")
	}
	return nil
}
