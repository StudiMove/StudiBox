package middleware

import (
	"log"
	"net/http"

	"backend/config"
	"backend/core/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware for Gin
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			log.Println("No token found in the request")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Forbidden",
				"message": "No token provided",
			})
			return
		}

		// Remove "Bearer " from the token
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid token format",
			})
			return
		}

		// Validate the JWT
		claims, err := utils.ValidateJWT(token, config.AppConfig.JwtSecretAccessKey)
		if err != nil {
			log.Printf("Failed to validate token: %v", err)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Forbidden",
				"message": "Invalid token",
			})
			return
		}

		// Store the claims in the Gin context
		log.Printf("Token validated, user ID: %v", claims.UserID)
		c.Set("user", claims)

		// Continue to the next handler
		c.Next()
	}
}
