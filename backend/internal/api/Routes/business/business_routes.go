// backend/internal/api/routes/business/business_routes.go

package business_routes

import (
    "net/http"
    "backend/internal/api/handlers/user/business/profil"
    "backend/internal/api/middlewares"  // Import du middleware
    "backend/internal/services/auth"
)

func RegisterBusinessRoutes(mux *http.ServeMux, getProfilHandler *profil.GetProfilHandler, updateProfilHandler *profil.UpdateProfilHandler, authService *auth.AuthService) {
    // Route pour récupérer les informations de profil (rôles : business, admin)
    mux.Handle("/business/profil", middleware.AuthMiddleware(middleware.RoleMiddleware(authService, []string{"business", "admin"}, http.HandlerFunc(getProfilHandler.HandleGetProfil))))

    // Route pour mettre à jour les informations de profil (rôle : admin uniquement)
    mux.Handle("/business/profil/update", middleware.AuthMiddleware(middleware.RoleMiddleware(authService, []string{"admin"}, http.HandlerFunc(updateProfilHandler.HandleUpdateProfil))))

    // Route de test sans middleware
    mux.Handle("/business/profil/test", middleware.AuthMiddleware(http.HandlerFunc(getProfilHandler.HandleGetProfil)))
}
