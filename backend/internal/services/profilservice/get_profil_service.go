package profilservice

import (
	"backend/internal/api/models/profil/response"
	"backend/internal/db/models"
	"errors"
)

// GetUserProfileByTargetID récupère le profil d'un utilisateur par son TargetUserID, incluant les informations utilisateur et organisation
func (s *ProfilService) GetUserProfileByTargetID(targetUserID uint) (response.UserProfileResponse, error) {
	var user models.User
	var userProfileResponse response.UserProfileResponse

		// Affiche le targetUserID pour s'assurer qu'il est passé correctement
		// fmt.Printf("Debug: targetUserID = %d\n", targetUserID)

		// Récupération des informations utilisateur avec le préchargement des rôles
		if err := s.DB.Preload("Roles").First(&user, "id = ?", targetUserID).Error; err != nil {
			// fmt.Printf("Erreur: impossible de trouver l'utilisateur avec l'ID %d. Erreur: %v\n", targetUserID, err)
			return userProfileResponse, errors.New("user not found ")
		}
	
		// Affiche les informations de l'utilisateur récupéré
		// fmt.Printf("Debug: Utilisateur trouvé: %+v\n", user)

	// Extraction des rôles de l'utilisateur
	var roleNames []string
	for _, role := range user.Roles {
		roleNames = append(roleNames, role.Name)
	}

	// Assemblage des informations utilisateur
	userProfileResponse = response.UserProfileResponse{
		UserID:       user.ID,
		Email:        user.Email,
		Phone:        user.Phone,
		ProfileImage: user.ProfileImage,
		RoleNames:    roleNames,
	}

	// Vérification des types d'organisation associés
	if s.populateOrganisationData(targetUserID, &userProfileResponse) {
		return userProfileResponse, nil
	}

	return userProfileResponse, errors.New("user not found in any organisation")
}

// populateOrganisationData remplit les données de l'organisation en fonction de l'utilisateur
func (s *ProfilService) populateOrganisationData(userID uint, profileResponse *response.UserProfileResponse) bool {
	var businessUser models.BusinessUser
	if err := s.DB.First(&businessUser, "user_id = ?", userID).Error; err == nil {
		profileResponse.Organisation = response.BusinessUserProfileResponse{
			Name:        businessUser.CompanyName,
			SIRET:       businessUser.SIRET,
			Address:     businessUser.Address,
			City:        businessUser.City,
			Postcode:    businessUser.Postcode,
			Country:     businessUser.Country,
			Description: businessUser.Description,
			Status: 	 businessUser.Status,	
		}
		return true
	}

	var schoolUser models.SchoolUser
	if err := s.DB.First(&schoolUser, "user_id = ?", userID).Error; err == nil {
		profileResponse.Organisation = response.SchoolUserProfileResponse{
			Name:        schoolUser.SchoolName,
			Address:     schoolUser.Address,
			City:        schoolUser.City,
			Postcode:    schoolUser.Postcode,
			Country:     schoolUser.Country,
			Description: schoolUser.Description,
			Status: 	 schoolUser.Status,	

		}
		return true
	}

	var associationUser models.AssociationUser
	if err := s.DB.First(&associationUser, "user_id = ?", userID).Error; err == nil {
		profileResponse.Organisation = response.AssociationUserProfileResponse{
			Name:        associationUser.AssociationName,
			Address:     associationUser.Address,
			City:        associationUser.City,
			Postcode:    associationUser.Postcode,
			Country:     associationUser.Country,
			Description: associationUser.Description,
			Status: 	 associationUser.Status,	

		}
		return true
	}

	return false
}