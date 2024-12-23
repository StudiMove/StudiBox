package routes

import (
	ticketHandler "backend/internal/api/handlers/ticket"
	ticketService "backend/internal/services/ticket"

	"github.com/gorilla/mux"
)

// RegisterTicketRoutes enregistre les routes pour les tickets.
func RegisterTicketRoutes(router *mux.Router, ticketService *ticketService.TicketService) {
	// Crée une nouvelle instance de TicketHandler
	handler := ticketHandler.NewTicketHandler(ticketService)

	// Route pour récupérer les tickets d'un utilisateur avec détails
	router.HandleFunc("/user/{userID}", handler.GetTicketsWithDetailsByUserIDHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/user/{userID}/buy", handler.GetEventIDsAndUUIDsHandler).Methods("GET", "OPTIONS")

	// Route pour créer un nouveau ticket
	router.HandleFunc("", handler.CreateTicketHandler).Methods("POST", "OPTIONS")

	// Route pour récupérer un ticket par ID
	router.HandleFunc("/id/{id}", handler.GetTicketByIDHandler).Methods("GET", "OPTIONS")
	// Route pour récupérer un ticket par ID
	router.HandleFunc("/{uuid}", handler.GetTicketByUUIDHandler).Methods("GET", "OPTIONS")

	// Route pour annuler un ticket
	router.HandleFunc("/{id}/cancel", handler.CancelTicketHandler).Methods("PUT", "OPTIONS")

	// Route pour marquer un ticket comme utilisé
	router.HandleFunc("/{id}/use", handler.MarkTicketAsUsedHandler).Methods("PUT", "OPTIONS")
}
