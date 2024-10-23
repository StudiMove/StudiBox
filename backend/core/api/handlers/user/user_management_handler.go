// user_management.go
package user

import (
	request "backend/core/api/request/user"
	responseGlobal "backend/core/api/response"
	"backend/core/services/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HandleUpdateUserByID gère la mise à jour d'un utilisateur par ID (Admin seulement)
func HandleUpdateUserByID(c *gin.Context, userService *user.UserService) {
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Invalid user ID", err))
		return
	}

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

	if err := userService.Management.UpdateUserProfile(uint(userID), updateProfileInput); err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to update user profile", err))
		return
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("User updated successfully", nil))
}
