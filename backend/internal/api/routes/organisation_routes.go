// package routes

package routes

import (
	"net/http"

	"backend/internal/api/handlers/organisation" // Handlers pour les routes organisation
	"backend/internal/api/middleware"
	"backend/internal/services/auth"
	orgService "backend/internal/services/organisation" // Service pour gérer les organisations

	"github.com/gorilla/mux"
)

// RegisterOrganisationRoutes enregistre les routes pour la gestion des organisations
func RegisterOrganisationRoutes(router *mux.Router, orgService *orgService.OrganisationService, authService *auth.AuthService) {
	// Initialisation des rôles nécessaires pour l'accès aux routes d'organisation
	requiredRoles := []string{"admin", "business", "school", "association"}

	// Handler pour récupérer toutes les organisations
	router.Handle("/all",
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(authService, requiredRoles)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				allOrgHandler := organisation.NewGetAllOrganisationsHandler(orgService)
				allOrgHandler.HandleGetAllOrganisations(w, r)
			})),
		),
	).Methods("GET", "OPTIONS")

	// Handler pour récupérer les organisations actives
	router.Handle("/all/active",
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(authService, requiredRoles)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				activeOrgHandler := organisation.NewGetAllOrganisationsHandler(orgService)
				activeOrgHandler.HandleGetActiveOrganisations(w, r)
			})),
		),
	).Methods("GET", "OPTIONS")

	// Handler pour récupérer les organisations inactives
	router.Handle("/all/inactive",
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(authService, requiredRoles)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				inactiveOrgHandler := organisation.NewGetAllOrganisationsHandler(orgService)
				inactiveOrgHandler.HandleGetInactiveOrganisations(w, r)
			})),
		),
	).Methods("GET", "OPTIONS")

	// Handler pour récupérer les organisations en attente
	router.Handle("/all/pending",
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(authService, requiredRoles)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				pendingOrgHandler := organisation.NewGetAllOrganisationsHandler(orgService)
				pendingOrgHandler.HandleGetPendingOrganisations(w, r)
			})),
		),
	).Methods("GET", "OPTIONS")

	// Handler pour récupérer les organisations suspendues
	router.Handle("/all/suspended",
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(authService, requiredRoles)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				suspendedOrgHandler := organisation.NewGetAllOrganisationsHandler(orgService)
				suspendedOrgHandler.HandleGetSuspendedOrganisations(w, r)
			})),
		),
	).Methods("GET", "OPTIONS")
}
