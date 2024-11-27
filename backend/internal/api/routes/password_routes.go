package routes

import (
	"backend/internal/api/handlers/passwords" // Chemin vers PasswordResetHandler
	"backend/internal/services/auth"
	"backend/internal/services/password"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterPasswordRoutes(router *mux.Router, passwordService *password.PasswordResetService, authService *auth.AuthService, db *gorm.DB) {
    handler := passwords.NewPasswordResetHandler(passwordService, authService, db)

    // Routes pour les actions sur les mots de passe
    router.HandleFunc("/request-reset", handler.HandleRequestPasswordReset).Methods("POST","OPTIONS")
    router.HandleFunc("/update", handler.HandleUpdatePassword).Methods("PUT","OPTIONS")
    router.HandleFunc("/get-reset-code", handler.HandleGetResetCode).Methods("POST","OPTIONS")
    router.HandleFunc("/verify-reset-code", handler.HandleVerifyResetCode).Methods("POST","OPTIONS")
}
