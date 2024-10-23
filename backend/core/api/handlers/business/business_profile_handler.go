package business

import (
	responseGlobal "backend/core/api/response"
	response "backend/core/api/response/business"
	"backend/core/services/business"
	"backend/core/services/user"
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
	resp := response.GetBusinessProfileResponse{
		CompanyName: businessProfile.CompanyName,
		Address:     businessProfile.Address,
		City:        businessProfile.City,
		Postcode:    businessProfile.Postcode,
		Country:     businessProfile.Country,
		Phone:       businessProfile.User.Phone,
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Business retrieved successfully", resp))
}
