// backend/internal/api/routes/auth_routes.go

package auth_routes

import (
    "net/http"
    "backend/internal/api/handlers/authentification"
)

// RegisterAuthRoutes enregistre les routes d'authentification
func RegisterAuthRoutes(mux *http.ServeMux, authHandler *authentification.AuthHandler, registerHandler *authentification.RegisterHandler) {
    mux.HandleFunc("/auth/login", authHandler.HandleLogin)
    mux.HandleFunc("/auth/register/user", registerHandler.HandleRegisterUser)
    mux.HandleFunc("/auth/register/business", registerHandler.HandleRegisterBusinessUser)
}
