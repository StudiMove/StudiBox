package profilservice

import (
	"backend/internal/api/models/profil/request"
	"backend/internal/db/models"
	"errors"
	"log"
)

// UpdateUserProfile met à jour le profil de l'utilisateur en fonction des rôles
func (s *ProfilService) UpdateUserProfile(userID uint, input request.UpdateProfileRequest, roles []models.Role) error {
	log.Printf("Début de UpdateUserProfile avec userID: %d et rôles: %+v", userID, roles)

	// Mise à jour du modèle User pour les champs communs, comme l'Email
	var user models.User
	if err := s.DB.First(&user, "id = ?", userID).Error; err != nil {
		return err
	}

	// Mise à jour des champs User
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Phone != "" {
		user.Phone = input.Phone
	}
	// Ajoutez d'autres champs de User si nécessaire

	// Sauvegarde des modifications du modèle User
	if err := s.DB.Save(&user).Error; err != nil {
		return err
	}

	// Ensuite, mise à jour selon le rôle
	for _, role := range roles {
		log.Printf("Traitement du rôle: %s pour userID: %d", role.Name, userID)
		switch role.Name {
		case "admin":
			log.Println("Mise à jour en tant que admin user dans la table business")
			return s.updateBusinessUserProfile(userID, input) // Mise à jour dans la table business pour admin
		case "business":
			log.Println("Mise à jour en tant que business user")
			return s.updateBusinessUserProfile(userID, input)
		case "school":
			log.Println("Mise à jour en tant que school user")
			return s.updateSchoolUserProfile(userID, input)
		case "association":
			log.Println("Mise à jour en tant que association user")
			return s.updateAssociationUserProfile(userID, input)
		default:
			log.Printf("Rôle non reconnu: %s", role.Name)
		}
	}

	log.Println("Aucun rôle reconnu pour l'utilisateur, échec de la mise à jour.")
	return errors.New("user role not recognized")
}

// UpdateUserProfileByTargetID met à jour le profil d'un utilisateur par targetUserID
func (s *ProfilService) UpdateUserProfileByTargetID(targetUserID uint, input request.UpdateProfileRequest) error {
	// Mettre à jour les informations de base dans la table User
	var user models.User
	if err := s.DB.First(&user, "id = ?", targetUserID).Error; err != nil {
		log.Printf("User with ID %d not found in User table", targetUserID)
		return errors.New("user not found")
	}

	log.Printf("User found with ID %d. Updating email and phone if provided.", targetUserID)
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Phone != "" {
		user.Phone = input.Phone
	}

	// Sauvegarder les changements dans User
	if err := s.DB.Save(&user).Error; err != nil {
		log.Printf("Error saving user ID %d in User table: %v", targetUserID, err)
		return err
	}
	log.Printf("Basic information updated for user ID %d in User table.", targetUserID)

	// Mise à jour des informations spécifiques dans les tables de profil
	var businessUser models.BusinessUser
	if err := s.DB.First(&businessUser, "user_id = ?", targetUserID).Error; err == nil {
		log.Printf("User ID %d found in BusinessUser table. Updating BusinessUser profile.", targetUserID)
		return s.updateBusinessUserProfile(targetUserID, input)
	} else {
		log.Printf("User ID %d not found in BusinessUser table: %v", targetUserID, err)
	}

	var schoolUser models.SchoolUser
	if err := s.DB.First(&schoolUser, "user_id = ?", targetUserID).Error; err == nil {
		log.Printf("User ID %d found in SchoolUser table. Updating SchoolUser profile.", targetUserID)
		return s.updateSchoolUserProfile(targetUserID, input)
	} else {
		log.Printf("User ID %d not found in SchoolUser table: %v", targetUserID, err)
	}

	var associationUser models.AssociationUser
	if err := s.DB.First(&associationUser, "user_id = ?", targetUserID).Error; err == nil {
		log.Printf("User ID %d found in AssociationUser table. Updating AssociationUser profile.", targetUserID)
		return s.updateAssociationUserProfile(targetUserID, input)
	} else {
		log.Printf("User ID %d not found in AssociationUser table: %v", targetUserID, err)
	}

	log.Printf("User with ID %d not found in any specific organisation table.", targetUserID)
	return errors.New("user not found in any organisation")
}

// Fonction de mise à jour pour un BusinessUser
func (s *ProfilService) updateBusinessUserProfile(userID uint, input request.UpdateProfileRequest) error {
	var businessUser models.BusinessUser
	if err := s.DB.First(&businessUser, "user_id = ?", userID).Error; err != nil {
		return err
	}

	if input.Name != "" {
		businessUser.CompanyName = input.Name
	}
	if input.Address != "" {
		businessUser.Address = input.Address
	}
	if input.City != "" {
		businessUser.City = input.City
	}
	if input.Postcode != "" {
		businessUser.Postcode = input.Postcode
	}
	if input.Country != "" {
		businessUser.Country = input.Country
	}
	if input.Region != "" {
		businessUser.Region = input.Region
	}
	if input.SIRET != "" {
		businessUser.SIRET = input.SIRET
	}
	if input.Description != "" {
		businessUser.Description = input.Description
	}
	if input.Status != "" {
		businessUser.Status = input.Status
	}
	businessUser.IsActivated = input.IsActivated

	businessUser.IsPending = input.IsPending

	businessUser.IsValidated = input.IsValidated

	return s.DB.Save(&businessUser).Error
}

// Fonction de mise à jour pour un SchoolUser
func (s *ProfilService) updateSchoolUserProfile(userID uint, input request.UpdateProfileRequest) error {
	var schoolUser models.SchoolUser
	if err := s.DB.First(&schoolUser, "user_id = ?", userID).Error; err != nil {
		return err
	}

	if input.Name != "" {
		schoolUser.SchoolName = input.Name
	}
	if input.Address != "" {
		schoolUser.Address = input.Address
	}
	if input.City != "" {
		schoolUser.City = input.City
	}
	if input.Postcode != "" {
		schoolUser.Postcode = input.Postcode
	}
	if input.Country != "" {
		schoolUser.Country = input.Country
	}
	if input.Region != "" {
		schoolUser.Region = input.Region
	}
	if input.SIRET != "" {
		schoolUser.SIRET = input.SIRET
	}
	if input.Description != "" {
		schoolUser.Description = input.Description
	}
	if input.Status != "" {
		schoolUser.Status = input.Status
	}
	schoolUser.IsActivated = input.IsActivated

	schoolUser.IsPending = input.IsPending

	schoolUser.IsValidated = input.IsValidated

	return s.DB.Save(&schoolUser).Error
}

// Fonction de mise à jour pour un AssociationUser
func (s *ProfilService) updateAssociationUserProfile(userID uint, input request.UpdateProfileRequest) error {
	var associationUser models.AssociationUser
	if err := s.DB.First(&associationUser, "user_id = ?", userID).Error; err != nil {
		return err
	}

	if input.Name != "" {
		associationUser.AssociationName = input.Name
	}
	if input.Address != "" {
		associationUser.Address = input.Address
	}
	if input.City != "" {
		associationUser.City = input.City
	}
	if input.Postcode != "" {
		associationUser.Postcode = input.Postcode
	}
	if input.Country != "" {
		associationUser.Country = input.Country
	}
	if input.Region != "" {
		associationUser.Region = input.Region
	}
	if input.SIRET != "" {
		associationUser.SIRET = input.SIRET
	}
	if input.Description != "" {
		associationUser.Description = input.Description
	}
	if input.Status != "" {
		associationUser.Status = input.Status
	}
	associationUser.IsActivated = input.IsActivated

	associationUser.IsPending = input.IsPending

	associationUser.IsValidated = input.IsValidated

	return s.DB.Save(&associationUser).Error
}
