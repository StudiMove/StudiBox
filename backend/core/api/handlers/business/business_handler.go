package business

import (
	request "backend/core/api/request/business"
	response "backend/core/api/response/business"
	"backend/core/services/business"
	"backend/core/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HandleGetBusinessProfile gère la récupération du profil business
func HandleGetBusinessProfile(c *gin.Context, businessService *business.BusinessService) {
	userClaims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing or invalid token"})
		return
	}

	userID := userClaims.(*utils.JWTClaims).UserID

	// Récupérer le profil de l'utilisateur business via le service
	businessProfile, err := businessService.GetBusinessUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user profile", "details": err.Error()})
		return
	}

	// Créer la réponse et retourner les données
	resp := response.GetBusinessProfileResponse{
		CompanyName: businessProfile.CompanyName,
		Address:     businessProfile.Address,
		City:        businessProfile.City,
		Postcode:    businessProfile.Postcode,
		Country:     businessProfile.Country,
		Phone:       businessProfile.User.Phone,
	}
	c.JSON(http.StatusOK, resp)
}

// HandleUpdateBusinessProfile gère la mise à jour du profil business
func HandleUpdateBusinessProfile(c *gin.Context, businessService *business.BusinessService) {
	userClaims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing or invalid token"})
		return
	}

	userID := userClaims.(*utils.JWTClaims).UserID

	var updateProfileReq request.UpdateBusinessProfileRequest
	if err := c.ShouldBindJSON(&updateProfileReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// Convertir UpdateBusinessProfileRequest en UpdateProfileInput pour l'utiliser dans le service
	updateProfileInput := business.UpdateProfileInput{
		CompanyName: updateProfileReq.CompanyName,
		Address:     updateProfileReq.Address,
		City:        updateProfileReq.City,
		Postcode:    updateProfileReq.Postcode,
		Country:     updateProfileReq.Country,
		Phone:       updateProfileReq.Phone,
	}

	// Appeler le service pour mettre à jour le profil business avec les nouveaux champs
	if err := businessService.UpdateBusinessUserProfile(userID, updateProfileInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

// HandleGetBusinessByID gère la récupération d'un business par ID
func HandleGetBusinessByID(c *gin.Context, businessService *business.BusinessService) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid business ID", "details": err.Error()})
		return
	}

	businessUser, err := businessService.GetBusinessUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve business", "details": err.Error()})
		return
	}

	resp := response.GetBusinessProfileResponse{
		CompanyName: businessUser.CompanyName,
		Address:     businessUser.Address,
		City:        businessUser.City,
		Postcode:    businessUser.Postcode,
		Country:     businessUser.Country,
		Phone:       businessUser.User.Phone,
	}

	c.JSON(http.StatusOK, resp)
}

// HandleGetAllBusinesses récupère tous les business users
func HandleGetAllBusinesses(c *gin.Context, businessService *business.BusinessService) {
	businessUsers, err := businessService.GetAllBusinessUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve businesses", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, businessUsers)
}

// HandleUpdateBusinessByID gère la mise à jour d'un business par ID (Admin seulement)
func HandleUpdateBusinessByID(c *gin.Context, businessService *business.BusinessService) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid business ID", "details": err.Error()})
		return
	}

	var updateProfileReq request.UpdateBusinessProfileRequest
	if err := c.ShouldBindJSON(&updateProfileReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	updateProfileInput := business.UpdateProfileInput{
		CompanyName: updateProfileReq.CompanyName,
		Address:     updateProfileReq.Address,
		City:        updateProfileReq.City,
		Postcode:    updateProfileReq.Postcode,
		Country:     updateProfileReq.Country,
		Phone:       updateProfileReq.Phone,
	}

	if err := businessService.UpdateBusinessUserProfile(uint(id), updateProfileInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update business profile", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Business updated successfully"})
}
