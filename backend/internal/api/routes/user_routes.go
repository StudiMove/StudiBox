package routes

import (
	"backend/internal/api/handlers/user"
	"backend/internal/api/middleware"
	"backend/internal/services/userservice"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterUserRoutes enregistre les routes utilisateur.
func RegisterUserRoutes(router *mux.Router, userService *userservice.UserService) {
	// Créez une nouvelle instance de UserHandler
	userHandler := user.NewUserHandler(userService)

	// Route pour déduire les coins en fonction de l'ID
	router.HandleFunc("/{id}/studibox", userHandler.GetStudiboxCoins).Methods("GET", "OPTIONS")

	router.Handle("/me/studibox",
		middleware.AuthMiddleware(http.HandlerFunc(userHandler.UpdateCoinsForAuthenticatedUser)),
	).Methods("PUT", "OPTIONS")

	// Route pour déduire les coins en fonction de l'ID
	router.HandleFunc("/{id}/studibox", userHandler.UpdateCoinsByID).Methods("PUT", "OPTIONS")

	// Route pour ajouter des coins en fonction de l'ID
	router.HandleFunc("/{id}/studibox/add", userHandler.AddCoinsByID).Methods("PUT", "OPTIONS")

	router.HandleFunc("/email-to-id", userHandler.GetUserIDHandler).Methods("POST", "OPTIONS")

	// Route pour mettre à jour les informations de l'utilisateur
	router.Handle("/me",
		middleware.AuthMiddleware(http.HandlerFunc(userHandler.UpdateUserHandler)),
	).Methods("PUT", "OPTIONS")

	router.Handle("/me/delete",
		middleware.AuthMiddleware(http.HandlerFunc(userHandler.DeleteUserHandler)),
	).Methods("DELETE", "OPTIONS")
}
