package user

import (
	request "backend/core/api/request/user"
	responseGlobal "backend/core/api/response"
	response "backend/core/api/response/user"
	"backend/core/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleUpdateUser gère la mise à jour du profil utilisateur
func HandleUpdateUser(c *gin.Context, userService *user.UserService) {
	userID, err := userService.Retrieval.GetUserIDFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Invalid user ID or token", err))
		return
	}

	var updateProfileReq request.UpdateUserProfileRequest
	if err := c.ShouldBindJSON(&updateProfileReq); err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Invalid input", err))
		return
	}

	if err := userService.Management.UpdateUserProfile(userID, updateProfileReq); err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to update user profile", err))
		return
	}

	// Réponse après mise à jour
	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("User profile updated successfully", &response.UserProfileResponse{
		ID:        userID,
		FirstName: updateProfileReq.FirstName,
		LastName:  updateProfileReq.LastName,
		Email:     updateProfileReq.Email,
		Phone:     updateProfileReq.Phone,
		Roles:     []string{"user"}, // Par exemple, si les rôles doivent être inclus
	}))
}
