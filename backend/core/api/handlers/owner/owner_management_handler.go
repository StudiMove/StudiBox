package owner

import (
	request "backend/core/api/request/owner"
	responseGlobal "backend/core/api/response"
	response "backend/core/api/response/owner"
	"backend/core/services/owner"
	"backend/core/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleUpdateOwner gère la mise à jour d'un profil owner
func HandleUpdateOwner(c *gin.Context, ownerService *owner.OwnerServiceType, userService *user.UserServiceType) {
	userID, err := userService.Retrieval.GetUserIDFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("ID utilisateur invalide ou token manquant", err))
		return
	}

	var updateProfileReq request.UpdateOwnerRequest
	if err := c.ShouldBindJSON(&updateProfileReq); err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Données de mise à jour invalides", err))
		return
	}

	if err := ownerService.Management.UpdateOwnerProfile(userID, updateProfileReq); err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Échec de la mise à jour du profil propriétaire", err))
		return
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Profil propriétaire mis à jour avec succès", &response.OwnerResponse{
		CompanyName: updateProfileReq.CompanyName,
		Address:     updateProfileReq.Address,
		City:        updateProfileReq.City,
		PostalCode:  updateProfileReq.PostalCode,
		Country:     updateProfileReq.Country,
		Phone:       updateProfileReq.Phone,
	}))
}
