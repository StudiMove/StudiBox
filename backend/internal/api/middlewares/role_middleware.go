package middleware

import (
    "net/http"
    "backend/internal/services/auth"
	"backend/internal/utils"
)

func RoleMiddleware(role string, authService *auth.AuthService) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            userClaims := r.Context().Value("user")
            if userClaims == nil {
                http.Error(w, "Forbidden", http.StatusForbidden)
                return
            }

            userID := userClaims.(*utils.JWTClaims).UserID
            hasRole, err := authService.CheckUserRole(userID, role)
            if err != nil || !hasRole {
                http.Error(w, "Forbidden", http.StatusForbidden)
                return
            }

            next.ServeHTTP(w, r)
        })
    }
}
