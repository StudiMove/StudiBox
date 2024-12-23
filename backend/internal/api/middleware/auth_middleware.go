// // backend/internal/api/middleware/auth_middleware.go
package middleware

import (
	"backend/config"
	"backend/internal/utils"
	"context"
	"log"
	"net/http"
)

// AuthMiddleware vérifie le token JWT et ajoute les claims au contexte
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			log.Println("No token found in the request")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Supprimer "Bearer " du token
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		} else {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		// Valider le JWT et obtenir les claims
		claims, err := utils.ValidateJWT(token, config.AppConfig.JwtSecretAccessKey)
		if err != nil {
			log.Printf("Failed to validate token: %v", err)
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Ajouter les claims entières au contexte
		log.Printf("Token validated successfully. UserID: %v", claims.UserID)
		ctx := context.WithValue(r.Context(), "user", claims)
		log.Printf("REGARDE ICI")

		log.Println("Requête reçue pour :", r.URL.Path)
		log.Println("Headers :", r.Header)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
