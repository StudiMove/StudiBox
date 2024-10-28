package user

import (
	responseGlobal "backend/core/api/response"
	response "backend/core/api/response/user"
	"backend/core/services/user"
	"backend/core/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HandleGetUser gère la récupération d'un utilisateur, que ce soit à partir du JWT ou du paramètre ID
func HandleGetUser(c *gin.Context, userService *user.UserService) {
	// Récupérer l'ID de l'utilisateur via JWT ou paramètre
	userID, err := userService.Retrieval.GetUserIDFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Invalid user ID or token", err))
		return
	}

	// Récupérer l'utilisateur par ID
	userProfile, err := userService.Retrieval.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, responseGlobal.ErrorResponse("User not found", err))
		return
	}

	// Extraire les rôles de l'utilisateur (si nécessaire)
	roleNames, err := userService.Management.ExtractRoleNames(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to retrieve user roles", err))
		return
	}

	// Créer la réponse avec le profil utilisateur
	resp := response.UserProfileResponse{
		ID:        userProfile.ID,
		FirstName: userProfile.FirstName,
		LastName:  userProfile.LastName,
		Email:     userProfile.Email,
		Phone:     userProfile.Phone,
		Roles:     roleNames,
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("User profile retrieved successfully", resp))
}

// HandleGetAllUsers récupère tous les utilisateurs (Admin seulement)
func HandleGetAllUsers(c *gin.Context, userService *user.UserService) {
	users, err := userService.Retrieval.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to retrieve users", err))
		return
	}

	// Création de la réponse pour la liste des utilisateurs
	var userResponses []response.UserProfileResponse
	for _, u := range users {
		userResponses = append(userResponses, response.UserProfileResponse{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Phone:     u.Phone,
		})
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Users retrieved successfully", userResponses))
}

// HandleExportUsersCSV gère l'exportation des utilisateurs en CSV
func HandleExportUsersCSV(c *gin.Context, userService *user.UserService) {
	// Récupérer tous les utilisateurs
	users, err := userService.Retrieval.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to retrieve users", err))
		return
	}

	// Vérifier si des utilisateurs existent
	if len(users) == 0 {
		c.JSON(http.StatusNotFound, responseGlobal.ErrorResponse("No users found", nil))
		return
	}

	// Préparer les données pour l'exportation
	headers := []string{"ID", "First Name", "Last Name", "Email", "Phone"}
	var rows [][]string
	for _, user := range users {
		rows = append(rows, []string{
			strconv.Itoa(int(user.ID)),
			user.FirstName,
			user.LastName,
			user.Email,
			user.Phone,
		})
	}

	// Utiliser l'utilitaire pour créer un fichier CSV avec point-virgule comme délimiteur
	fileName, err := utils.ExportToCSV(headers, rows, ';')
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to generate CSV", err))
		return
	}

	// Définir les en-têtes pour le téléchargement de fichier CSV
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=users.csv")
	c.Header("Content-Type", "text/csv")

	// Envoyer le fichier CSV
	c.File(fileName)
}
