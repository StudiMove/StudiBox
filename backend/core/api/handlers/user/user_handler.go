package user

import (
	"backend/core/models"
	"backend/core/services/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Variable globale pour le service utilisateur
var UserService *user.UserService

// HandleGetUserByID gère la récupération d'un utilisateur par ID.
func HandleGetUserByID(c *gin.Context) {
	// Extraire l'ID de l'utilisateur des paramètres de l'URL
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := UserService.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// HandleUpdateUser gère la mise à jour d'un utilisateur
func HandleUpdateUser(c *gin.Context) {
	var updateRequest models.User
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Récupérer l'utilisateur par son ID
	user, err := UserService.GetUserByID(updateRequest.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Mettre à jour les champs de l'utilisateur
	user.FirstName = updateRequest.FirstName
	user.LastName = updateRequest.LastName
	user.Email = updateRequest.Email

	// Mettre à jour l'utilisateur dans la base de données
	if err := UserService.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
