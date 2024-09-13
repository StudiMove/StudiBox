package routes

import (
    "net/http"
    "backend/internal/api/handlers/authentification"
)

func RegisterAuthRoutes(authHandler *authentification.AuthHandler, registerHandler *authentification.RegisterHandler) {
    http.HandleFunc("/login", authHandler.HandleLogin)
    http.HandleFunc("/register/user", registerHandler.HandleRegisterUser)
    http.HandleFunc("/register/business", registerHandler.HandleRegisterBusinessUser)
}
