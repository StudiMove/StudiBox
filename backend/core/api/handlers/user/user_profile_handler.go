// user_profile.go
package user

import (
	request "backend/core/api/request/user"
	responseGlobal "backend/core/api/response"
	response "backend/core/api/response/user"
	"backend/core/services/user"
	"backend/core/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleGetUserProfile gère la récupération du profil utilisateur
func HandleGetUserProfile(c *gin.Context, userService *user.UserService) {
	userClaims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, responseGlobal.ErrorResponse("Unauthorized: Missing or invalid token", nil))
		return
	}

	userID := userClaims.(*utils.JWTClaims).UserID

	userProfile, err := userService.Retrieval.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to retrieve user profile", err))
		return
	}

	resp := response.GetUserProfileResponse{
		FirstName: userProfile.FirstName,
		LastName:  userProfile.LastName,
		Email:     userProfile.Email,
		Phone:     userProfile.Phone,
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("User profile retrieved successfully", resp))
}

// HandleUpdateUserProfile gère la mise à jour du profil utilisateur
func HandleUpdateUserProfile(c *gin.Context, userService *user.UserService) {
	userClaims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, responseGlobal.ErrorResponse("Unauthorized: Missing or invalid token", nil))
		return
	}

	userID := userClaims.(*utils.JWTClaims).UserID

	var updateProfileReq request.UpdateUserProfileRequest
	if err := c.ShouldBindJSON(&updateProfileReq); err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Invalid input", err))
		return
	}

	updateProfileInput := request.UpdateUserProfileRequest{
		FirstName: updateProfileReq.FirstName,
		LastName:  updateProfileReq.LastName,
		Email:     updateProfileReq.Email,
		Phone:     updateProfileReq.Phone,
	}

	if err := userService.Management.UpdateUserProfile(userID, updateProfileInput); err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to update user profile", err))
		return
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Profile updated successfully", nil))
}
