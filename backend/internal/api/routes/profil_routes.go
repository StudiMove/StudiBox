package routes

import (
	"backend/internal/api/handlers/profil"
	"backend/internal/api/handlers/user"
	"backend/internal/api/middleware"
	"backend/internal/services/auth"
	"backend/internal/services/profilservice"
	"backend/internal/services/userservice" // Import du userService
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterProfilRoutes(
	router *mux.Router,
	profilService *profilservice.ProfilService,
	authService *auth.AuthService,
	userService *userservice.UserService,
	jwtSecret string,
) {
	profilHandler := profil.NewProfilHandler(profilService)
	updateProfilpHandler := profil.NewUpdateProfilHandler(profilService, userService)
	uploadProfilImageHandler := profil.NewUploadProfilImageHandler(profilService, jwtSecret)

	// Définissez les rôles requis pour cette route
	requiredRoles := []string{"admin", "business", "school", "association"}
	adminRoles := []string{"admin"}

	// Route pour récupérer le rôle de l'utilisateur
	router.HandleFunc("/user/role", user.GetUserRoleHandler(userService)).Methods("GET", "POST", "OPTIONS")

	// Route pour récupérer le profil de l'utilisateur avec AuthMiddleware et RoleMiddleware en cascade
	router.Handle("/organisation/profile",
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(authService, requiredRoles)(http.HandlerFunc(profilHandler.GetUserProfile)),
		),
	).Methods("GET", "OPTIONS")

	router.Handle("/organisation/profile/targetId",
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(authService, adminRoles)(http.HandlerFunc(profilHandler.GetUserProfileByTargetID)),
		),
	).Methods("POST", "OPTIONS", "GET")

	// Route pour mise à jour du profil de l'utilisateur connecté
	router.Handle("/organisation/profile/update",
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(authService, requiredRoles)(http.HandlerFunc(updateProfilpHandler.HandleUpdateOwnProfile)),
		),
	).Methods("PUT", "OPTIONS")

	// Route pour mise à jour du profil d'un utilisateur cible
	router.Handle("/organisation/profile/update/targetId",
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(authService, adminRoles)(http.HandlerFunc(updateProfilpHandler.HandleUpdateTargetProfile)),
		),
	).Methods("PUT", "OPTIONS")

	// Route pour upload de l'image de profil de l'utilisateur connecté
	router.Handle("/organisation/profile/upload-image",
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(authService, requiredRoles)(http.HandlerFunc(uploadProfilImageHandler.HandleUploadProfileImage)),
		),
	).Methods("POST", "OPTIONS")

	// Route pour upload de l'image de profil d'un utilisateur cible via targetId
	router.Handle("/organisation/profile/upload-image/targetId",
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(authService, adminRoles)(http.HandlerFunc(uploadProfilImageHandler.HandleUploadProfileImageWithTargetID)),
		),
	).Methods("POST", "OPTIONS")
}
