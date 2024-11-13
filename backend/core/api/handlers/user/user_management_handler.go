package user

import (
	request "backend/core/api/request/user"
	responseGlobal "backend/core/api/response"
	response "backend/core/api/response/user"
	"backend/core/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleUpdateUser gère la mise à jour du profil utilisateur
// HandleUpdateUser gère la mise à jour du profil utilisateur
func HandleUpdateUser(c *gin.Context, userService *user.UserServiceType) {
	userID, err := userService.Retrieval.GetUserIDFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("ID utilisateur ou token invalide", err))
		return
	}

	var updateProfileReq request.UserRequest
	if err := c.ShouldBindJSON(&updateProfileReq); err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Données invalides", err))
		return
	}

	// Mettre à jour le profil utilisateur
	if err := userService.Management.UpdateUserProfile(userID, updateProfileReq); err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Échec de la mise à jour du profil utilisateur", err))
		return
	}

	// Récupérer le profil mis à jour
	userProfile, err := userService.Retrieval.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Échec de la récupération du profil mis à jour", err))
		return
	}

	roleNames, err := userService.Management.ExtractRoleName(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Erreur lors de la récupération du rôle", err))
		return
	}

	resp := response.UserResponse{
		ID:            userProfile.ID,
		FirstName:     userProfile.FirstName,
		LastName:      userProfile.LastName,
		Pseudo:        userProfile.Pseudo,
		Email:         userProfile.Email,
		Phone:         userProfile.Phone,
		ProfileImage:  userProfile.ProfileImage,
		BirthDate:     userProfile.BirthDate.Format("2006-01-02"),
		Country:       userProfile.Country,
		City:          userProfile.City,
		Address:       userProfile.Address,
		PostalCode:    userProfile.PostalCode,
		ProfileType:   userProfile.ProfileType,
		StudiboxCoins: userProfile.StudiboxCoins,
		Role:          roleNames,
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Profil utilisateur mis à jour avec succès", resp))
}
