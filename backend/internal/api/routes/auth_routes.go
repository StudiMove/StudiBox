package routes

import (
	"backend/internal/api/handlers/authentification"
	"backend/internal/services/auth"
	"backend/pkg/httpclient"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(router *mux.Router, authService *auth.AuthService, httpClient *httpclient.APIClient) {
	handler := authentification.NewAuthHandler(authService)
	registerHandler := authentification.NewRegisterHandler(authService, httpClient)

	// Définir les routes d'authentification
	router.HandleFunc("/register/organisation", registerHandler.HandleRegisterOrganisationUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", handler.HandleLogin).Methods("POST", "OPTIONS")
	router.HandleFunc("/login/user", handler.HandleLoginNormalUser).Methods("POST", "OPTIONS")

	router.HandleFunc("/register/user", registerHandler.HandleRegisterNormalUser).Methods("POST", "OPTIONS")

	router.HandleFunc("/get-user-id", handler.HandleGetUserIDByEmail).Methods("POST", "OPTIONS")
}
