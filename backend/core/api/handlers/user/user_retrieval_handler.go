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

// HandleGetUser gère la récupération d'un profil utilisateur à partir de l'ID ou du JWT
func HandleGetUser(c *gin.Context, userService *user.UserServiceType) {
	userID, err := userService.Retrieval.GetUserIDFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("ID utilisateur ou token invalide", err))
		return
	}

	userProfile, err := userService.Retrieval.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, responseGlobal.ErrorResponse("Utilisateur non trouvé", err))
		return
	}

	resp := response.UserResponse{
		ID:            userProfile.ID,
		FirstName:     userProfile.FirstName,
		LastName:      userProfile.LastName,
		Pseudo:        userProfile.Pseudo,
		Email:         userProfile.Email,
		Phone:         userProfile.Phone,
		Country:       userProfile.Country,
		Region:        userProfile.Region,
		PostalCode:    userProfile.PostalCode,
		Address:       userProfile.Address,
		BirthDate:     userProfile.BirthDate.Format("2006-01-02"),
		ProfileImage:  userProfile.ProfileImage,
		ProfileType:   userProfile.ProfileType,
		StudiboxCoins: userProfile.StudiboxCoins,
		Role:          userProfile.Role.Name,
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Profil utilisateur récupéré avec succès", resp))
}

// HandleGetAllUsers récupère tous les utilisateurs (Admin seulement)
func HandleGetAllUsers(c *gin.Context, userService *user.UserServiceType) {
	users, err := userService.Retrieval.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Échec de la récupération des utilisateurs", err))
		return
	}

	var userResponses []response.UserResponse
	for _, u := range users {
		userResponses = append(userResponses, response.UserResponse{
			ID:            u.ID,
			FirstName:     u.FirstName,
			LastName:      u.LastName,
			Pseudo:        u.Pseudo,
			Email:         u.Email,
			Phone:         u.Phone,
			Country:       u.Country,
			Region:        u.Region,
			PostalCode:    u.PostalCode,
			Address:       u.Address,
			BirthDate:     u.BirthDate.Format("2006-01-02"),
			ProfileImage:  u.ProfileImage,
			ProfileType:   u.ProfileType,
			StudiboxCoins: u.StudiboxCoins,
			Role:          u.Role.Name,
		})
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Utilisateurs récupérés avec succès", userResponses))
}

// HandleExportUsersCSV gère l'exportation des utilisateurs en format CSV
func HandleExportUsersCSV(c *gin.Context, userService *user.UserServiceType) {
	users, err := userService.Retrieval.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Échec de la récupération des utilisateurs", err))
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, responseGlobal.ErrorResponse("Aucun utilisateur trouvé", nil))
		return
	}

	headers := []string{"ID", "Prénom", "Nom", "Email", "Téléphone", "Country", "Region", "Code Postal", "City", "Address", "Pseudo", "Rôle"}
	var rows [][]string
	for _, user := range users {
		rows = append(rows, []string{
			strconv.Itoa(int(user.ID)),
			user.FirstName,
			user.LastName,
			user.Email,
			strconv.Itoa(user.Phone),
			user.Country,
			user.Region,
			strconv.Itoa(int(user.PostalCode)),
			user.City,
			user.Address,
			user.Pseudo,
			user.Role.Name,
		})
	}

	fileName, err := utils.ExportToCSV(headers, rows, ';')
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Échec de la génération du fichier CSV", err))
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=users.csv")
	c.Header("Content-Type", "text/csv")
	c.File(fileName)
}
