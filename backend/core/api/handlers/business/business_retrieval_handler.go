package business

import (
	responseGlobal "backend/core/api/response"
	response "backend/core/api/response/business"
	"backend/core/services/business"
	"backend/core/services/user"
	"log" // Pour la journalisation des erreurs
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleGetBusiness gère la récupération soit du profil de l'utilisateur connecté, soit d'un business par ID
func HandleGetBusiness(c *gin.Context, businessService *business.BusinessService, userService *user.UserService) {
	// Récupérer l'ID de l'utilisateur (connecté ou via paramètre ID)
	userID, err := userService.Retrieval.GetUserIDFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Invalid business ID or missing token", err))
		return
	}

	// Récupérer le profil de business en fonction de l'ID utilisateur
	businessProfile, err := businessService.Retrieval.GetBusinessUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to retrieve business", err))
		return
	}

	// Réponse avec le profil business
	resp := response.BusinessProfileResponse{
		CompanyName: businessProfile.CompanyName,
		Address:     businessProfile.Address,
		City:        businessProfile.City,
		Postcode:    businessProfile.Postcode,
		Country:     businessProfile.Country,
		Phone:       businessProfile.User.Phone,
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Business retrieved successfully", resp))
}

// HandleGetAllBusinesses récupère tous les utilisateurs business
func HandleGetAllBusinesses(c *gin.Context, businessService *business.BusinessService) {
	businessUsers, err := businessService.Retrieval.GetAllBusinessUsers()

	if err != nil {
		// Journaliser l'erreur pour le serveur
		log.Printf("Erreur lors de la récupération des utilisateurs business: %v", err)

		// Retourner une réponse d'erreur générique avec le statut HTTP 500
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to retrieve businesses", err))
		return
	}

	// Vérification s'il n'y a aucun utilisateur business à retourner
	if len(businessUsers) == 0 {
		c.JSON(http.StatusOK, responseGlobal.SuccessResponse("No businesses found", nil))
		return
	}

	// Si tout va bien, on retourne les utilisateurs business
	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Businesses retrieved successfully", businessUsers))
}
