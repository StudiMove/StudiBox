package middleware

import (
    "net/http"
    "backend/internal/services/auth"
    "backend/internal/utils"
)

// RoleMiddleware accepte une liste de rôles requis
func RoleMiddleware(authService *auth.AuthService, requiredRoles []string, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        userClaims := r.Context().Value("user").(*utils.JWTClaims)

        // Appel de CheckUserRole avec une liste de rôles
        hasRole, err := authService.CheckUserRole(userClaims.UserID, requiredRoles)
        if err != nil || !hasRole {
            http.Error(w, "Forbidden: Not access with role permission", http.StatusForbidden)
            return
        }

        next.ServeHTTP(w, r)
    })
}
