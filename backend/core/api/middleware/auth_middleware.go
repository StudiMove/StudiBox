// func AuthMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         token := r.Header.Get("Authorization")
//         if token == "" {
//             http.Error(w, "Forbidden", http.StatusForbidden)
//             return
//         }

//         // Valider le JWT
//         claims, err := utils.ValidateJWT(token, config.AppConfig.JwtSecretAccessKey)
//         if err != nil {
//             http.Error(w, "Forbidden", http.StatusForbidden)
//             return
//         }

//         // Ajouter les claims au contexte
//         ctx := context.WithValue(r.Context(), "user", claims)
//         next.ServeHTTP(w, r.WithContext(ctx))
//     })
// }
// func AuthMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         token := r.Header.Get("Authorization")
//         if token == "" {
//             log.Println("No token found in the request")
//             http.Error(w, "Forbidden", http.StatusForbidden)
//             return
//         }

//         // Valider le JWT
//         claims, err := utils.ValidateJWT(token, config.AppConfig.JwtSecretAccessKey)
//         if err != nil {
//             log.Printf("Failed to validate token: %v", err)
//             http.Error(w, "Forbidden", http.StatusForbidden)
//             return
//         }

//         // Ajouter les claims au contexte
//         log.Printf("Token validated, user ID: %v", claims.UserID)
//         ctx := context.WithValue(r.Context(), "user", claims)
//         next.ServeHTTP(w, r.WithContext(ctx))
//     })
// }

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
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		// Remove "Bearer " from the token
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}

		// Validate the JWT
		claims, err := utils.ValidateJWT(token, config.AppConfig.JwtSecretAccessKey)
		if err != nil {
			log.Printf("Failed to validate token: %v", err)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		// Store the claims in the Gin context
		log.Printf("Token validated, user ID: %v", claims.UserID)
		c.Set("user", claims)

		// Continue to the next handler
		c.Next()
	}
}
