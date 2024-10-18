package middleware

import (
	"backend/core/services/auth"
	"backend/core/utils"
	"net/http"
)

// RoleMiddleware accepte une liste de rôles requis et vérifie si l'utilisateur a l'un de ces rôles
func RoleMiddleware(authService *auth.AuthService, requiredRoles []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Récupérer les informations de l'utilisateur à partir du contexte
		userClaims, ok := r.Context().Value("user").(*utils.JWTClaims)
		if !ok || userClaims == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Vérifier si l'utilisateur a l'un des rôles requis
		hasRole := false
		for _, role := range requiredRoles {
			hasRole, _ = authService.CheckUserRole(userClaims.UserID, role)
			if hasRole {
				break
			}
		}

		// Si l'utilisateur n'a pas le rôle requis
		if !hasRole {
			http.Error(w, "Forbidden: Not access with role permission", http.StatusForbidden)
			return
		}

		// Si tout est bon, continuer avec la requête
		next.ServeHTTP(w, r)
	})
}
