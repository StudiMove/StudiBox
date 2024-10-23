package user

import (
	responseGlobal "backend/core/api/response"
	response "backend/core/api/response/user"
	"backend/core/services/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HandleGetUserByID gère la récupération d'un utilisateur par ID
func HandleGetUserByID(c *gin.Context, userService *user.UserService) {
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil || userID <= 0 {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Invalid user ID", err))
		return
	}

	// Récupérer l'utilisateur par ID
	user, err := userService.Retrieval.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, responseGlobal.ErrorResponse("User not found", err))
		return
	}

	// Créer la réponse avec le profil utilisateur
	resp := response.GetUserProfileResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("User retrieved successfully", resp))
}

// HandleGetAllUsers récupère tous les utilisateurs (Admin seulement)
func HandleGetAllUsers(c *gin.Context, userService *user.UserService) {
	users, err := userService.Retrieval.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to retrieve users", err))
		return
	}

	// Création de la réponse pour la liste des utilisateurs
	var userResponses []response.GetUserProfileResponse
	for _, u := range users {
		userResponses = append(userResponses, response.GetUserProfileResponse{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Phone:     u.Phone,
		})
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Users retrieved successfully", userResponses))
}
