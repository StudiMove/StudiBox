// backend/internal/api/middleware/role_middleware.go
package middleware

import (
	"backend/internal/api/models/auth/request"
	"backend/internal/services/auth"
	"backend/internal/utils"
	"log"
	"net/http"
)

func RoleMiddleware(authService *auth.AuthService, requiredRoles []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Récupérer les claims utilisateur à partir du contexte
			userClaims, ok := r.Context().Value("user").(*utils.JWTClaims)
			if !ok || userClaims == nil {
				log.Println("Unauthorized: User claims not found in context")
				http.Error(w, "Unauthorized: User claims not found", http.StatusUnauthorized)
				return
			}

			userID := userClaims.UserID
			log.Printf("UserID from claims: %d", userID)

			// Créer la requête `CheckUserRoleRequest`
			roleReq := &request.CheckUserRoleRequest{
				UserID: userID,
				Roles:  requiredRoles,
			}

			// Vérifier le rôle requis en utilisant le service d'authentification
			hasRoleResp, err := authService.CheckUserRole(roleReq)
			if err != nil {
				log.Printf("Error checking role permissions: %v", err)
				http.Error(w, "Error checking role permissions", http.StatusInternalServerError)
				return
			}

			// Vérifier si l'utilisateur a le rôle requis via `HasRole` de `CheckUserRoleResponse`
			if !hasRoleResp.HasRole {
				log.Println("Forbidden: Insufficient role permissions")
				http.Error(w, "Forbidden: Insufficient role permissions", http.StatusForbidden)
				return
			}

			// Passer au handler suivant
			next.ServeHTTP(w, r)
		})
	}
}
