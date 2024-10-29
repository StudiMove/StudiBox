package middleware

import (
	"backend/core/api/response"
	"backend/core/services/auth"
	"backend/core/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoleMiddleware vérifie si l'utilisateur a l'un des rôles requis
func RoleMiddleware(authService *auth.AuthService, requiredRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Récupérer les informations de l'utilisateur à partir du contexte
		claimsValue, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse("Unauthorized", nil))
			return
		}

		claims, ok := claimsValue.(*utils.JWTClaims)
		if !ok || claims == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse("Token is missing or invalid", nil))
			return
		}

		// Vérifier si l'utilisateur a l'un des rôles requis
		hasRole := false
		for _, role := range requiredRoles {
			// Vérifier si l'utilisateur a ce rôle
			roleExists, err := authService.Role.CheckUserRole(claims.UserID, role)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse("Internal error checking roles", err))
				return
			}
			if roleExists {
				hasRole = true
				break
			}
		}

		// Si l'utilisateur n'a pas le rôle requis
		if !hasRole {
			c.AbortWithStatusJSON(http.StatusForbidden, response.ErrorResponse("Insufficient permissions", nil))
			return
		}

		// Continuer la requête si tout va bien
		c.Next()
	}
}
