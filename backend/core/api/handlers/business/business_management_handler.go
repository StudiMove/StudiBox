package business

import (
	request "backend/core/api/request/business"
	responseGlobal "backend/core/api/response"
	"backend/core/services/business"
	"backend/core/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleUpdateBusiness gère la mise à jour d'un business soit pour l'utilisateur connecté, soit par ID (Admin)
func HandleUpdateBusiness(c *gin.Context, businessService *business.BusinessService, userService *user.UserService) {
	// Récupérer l'ID de l'utilisateur (connecté ou via paramètre ID)
	userID, err := userService.Retrieval.GetUserIDFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Invalid business ID or missing token", err))
		return
	}

	// Validation des données du profil à mettre à jour
	var updateProfileReq request.UpdateBusinessProfileRequest
	if err := c.ShouldBindJSON(&updateProfileReq); err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Invalid input data", err))
		return
	}

	// Mise à jour du profil business
	if err := businessService.Management.UpdateBusinessUserProfile(userID, updateProfileReq); err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to update business profile", err))
		return
	}

	// Réponse en cas de succès
	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Business updated successfully", nil))
}
