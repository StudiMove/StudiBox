package owner

import (
	responseGlobal "backend/core/api/response"
	response "backend/core/api/response/owner"
	"backend/core/services/owner"
	"backend/core/services/user"
	"backend/database/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HandleGetOwner gère la récupération soit du profil de l'utilisateur connecté, soit d'un owner par ID
func HandleGetOwner(c *gin.Context, ownerService *owner.OwnerServiceType, userService *user.UserServiceType) {
	var ownerProfile *models.Owner
	var err error

	// Récupérer le paramètre ID de l'URL (si présent)
	idParam := c.Param("id")

	// Si un ID est fourni dans l'URL, récupérer le propriétaire par cet ID
	if idParam != "" {
		ownerID, parseErr := strconv.Atoi(idParam)
		if parseErr != nil {
			c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("ID invalide", parseErr))
			return
		}
		ownerProfile, err = ownerService.Retrieval.GetOwnerByID(uint(ownerID))
		if err != nil {
			c.JSON(http.StatusNotFound, responseGlobal.ErrorResponse("Propriétaire non trouvé", err))
			return
		}
	} else {
		// Sinon, récupérer le profil du propriétaire connecté à partir du token
		userID, err := userService.Retrieval.GetUserIDFromRequest(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Token invalide ou utilisateur non connecté", err))
			return
		}
		ownerProfile, err = ownerService.Retrieval.GetOwnerByUserID(userID)
		if err != nil {
			c.JSON(http.StatusNotFound, responseGlobal.ErrorResponse("Propriétaire non trouvé pour l'utilisateur connecté", err))
			return
		}
	}

	// Construire la réponse avec les informations du propriétaire
	resp := response.OwnerResponse{
		ID:          ownerProfile.ID,
		CompanyName: ownerProfile.CompanyName,
		Type:        ownerProfile.Type,
		Address:     ownerProfile.Address,
		City:        ownerProfile.City,
		PostalCode:  ownerProfile.PostalCode,
		Region:      ownerProfile.Region,
		Country:     ownerProfile.Country,
		Phone:       ownerProfile.Phone,
		Description: ownerProfile.Description,
		Status:      ownerProfile.Status,
		SIRET:       ownerProfile.SIRET,
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Profil propriétaire récupéré avec succès", resp))
}

// HandleGetAllOwners récupère tous les utilisateurs owner
func HandleGetAllOwners(c *gin.Context, ownerService *owner.OwnerServiceType) {
	ownerUsers, err := ownerService.Retrieval.GetAllOwnerUsers()
	if err != nil {
		log.Printf("Erreur lors de la récupération des propriétaires: %v", err)
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Échec de la récupération des propriétaires", err))
		return
	}

	if len(ownerUsers) == 0 {
		c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Aucun propriétaire trouvé", nil))
		return
	}

	var ownerResponses []response.OwnerResponse
	for _, ownerProfile := range ownerUsers {
		ownerResponses = append(ownerResponses, response.OwnerResponse{
			ID:          ownerProfile.ID,
			CompanyName: ownerProfile.CompanyName,
			Type:        ownerProfile.Type,
			Address:     ownerProfile.Address,
			City:        ownerProfile.City,
			PostalCode:  ownerProfile.PostalCode,
			Region:      ownerProfile.Region,
			Country:     ownerProfile.Country,
			Phone:       ownerProfile.User.Phone,
			Description: ownerProfile.Description,
			Status:      ownerProfile.Status,
			SIRET:       ownerProfile.SIRET,
		})
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Propriétaires récupérés avec succès", ownerResponses))
}
