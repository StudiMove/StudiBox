package routes

import (
	"backend/internal/api/handlers/referral"
	referralService "backend/internal/services/referral" // Alias pour éviter les conflits de noms

	"github.com/gorilla/mux"
)

// RegisterReferralRoutes enregistre les routes liées au parrainage
func RegisterReferralRoutes(router *mux.Router, service *referralService.ReferralService) {
	// Initialiser le handler avec le service de parrainage
	handler := referral.NewReferralHandler(service)

	// Définir les routes pour les parrainages
	router.HandleFunc("/filleuls", handler.HandleGetFilleulsByParrain).Methods("POST", "OPTIONS")
	router.HandleFunc("/count", handler.HandleCountFilleuls).Methods("POST", "OPTIONS")
	router.HandleFunc("/filleuls/ids", handler.HandleGetFilleulIDsByParrain).Methods("POST", "OPTIONS")

}
