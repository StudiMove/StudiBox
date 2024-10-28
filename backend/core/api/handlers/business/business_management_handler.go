package business

import (
	request "backend/core/api/request/business"
	responseGlobal "backend/core/api/response"
	response "backend/core/api/response/business"
	"backend/core/services/business"
	"backend/core/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleUpdateBusiness gère la mise à jour d'un business
func HandleUpdateBusiness(c *gin.Context, businessService *business.BusinessService, userService *user.UserService) {
	userID, err := userService.Retrieval.GetUserIDFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Invalid business ID or missing token", err))
		return
	}

	var updateProfileReq request.UpdateBusinessProfileRequest
	if err := c.ShouldBindJSON(&updateProfileReq); err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Invalid input data", err))
		return
	}

	if err := businessService.Management.UpdateBusinessUserProfile(userID, updateProfileReq); err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to update business profile", err))
		return
	}

	// Réponse après mise à jour
	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Business updated successfully", &response.BusinessProfileResponse{
		CompanyName: updateProfileReq.CompanyName,
		Address:     updateProfileReq.Address,
		City:        updateProfileReq.City,
		Postcode:    updateProfileReq.Postcode,
		Country:     updateProfileReq.Country,
		Phone:       updateProfileReq.Phone,
	}))
}
