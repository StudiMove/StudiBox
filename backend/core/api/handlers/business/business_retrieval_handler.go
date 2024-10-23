package business

import (
	responseGlobal "backend/core/api/response"
	"backend/core/services/business"
	"log" // Pour la journalisation des erreurs
	"net/http"

	"github.com/gin-gonic/gin"
)

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
