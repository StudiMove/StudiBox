package ticket

import (
	"backend/internal/services/ticket"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// TicketHandler représente le gestionnaire HTTP pour les tickets.
type TicketHandler struct {
	service *ticket.TicketService
}

// NewTicketHandler crée une nouvelle instance de TicketHandler.
func NewTicketHandler(service *ticket.TicketService) *TicketHandler {
	return &TicketHandler{service: service}
}

// CreateTicketHandler gère la création d'un ticket.
func (h *TicketHandler) CreateTicketHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		UserID    uint   `json:"user_id"`
		EventID   uint   `json:"event_id"`
		TarifIDs  []uint `json:"tarif_ids"`
		OptionIDs []uint `json:"option_ids"`
	}

	// Décoder le corps de la requête
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Appeler le service pour créer un ticket
	ticket, err := h.service.CreateTicketWithDetails(payload.UserID, payload.EventID, payload.TarifIDs, payload.OptionIDs)
	if err != nil {
		http.Error(w, "Failed to create ticket: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Retourner le ticket créé
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ticket)
}

// GetTicketByIDHandler récupère un ticket par son ID.
func (h *TicketHandler) GetTicketByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ticketID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	// Appeler le service pour récupérer le ticket
	ticket, err := h.service.GetTicketByID(uint(ticketID))
	if err != nil {
		http.Error(w, "Ticket not found: "+err.Error(), http.StatusNotFound)
		return
	}

	// Retourner le ticket
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ticket)
}

// GetTicketsByUserIDHandler récupère les tickets d'un utilisateur donné.
func (h *TicketHandler) GetTicketsByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["userID"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Appeler le service pour récupérer les tickets
	tickets, err := h.service.GetTicketsByUserID(uint(userID))
	if err != nil {
		http.Error(w, "Failed to fetch tickets: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Retourner les tickets
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

// CancelTicketHandler annule un ticket.
func (h *TicketHandler) CancelTicketHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ticketID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	// Appeler le service pour annuler le ticket
	if err := h.service.CancelTicket(uint(ticketID)); err != nil {
		http.Error(w, "Failed to cancel ticket: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Retourner une réponse de succès
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ticket successfully cancelled"))
}

// MarkTicketAsUsedHandler marque un ticket comme utilisé.
func (h *TicketHandler) MarkTicketAsUsedHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ticketID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	// Appeler le service pour marquer le ticket comme utilisé
	if err := h.service.MarkTicketAsUsed(uint(ticketID)); err != nil {
		http.Error(w, "Failed to mark ticket as used: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Retourner une réponse de succès
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ticket successfully marked as used"))
}

// GetTicketsWithDetailsByUserIDHandler récupère les tickets d'un utilisateur avec les détails.
func (h *TicketHandler) GetTicketsWithDetailsByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["userID"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Appeler le service pour récupérer les tickets détaillés
	tickets, err := h.service.GetTicketsWithDetailsByUserID(uint(userID))
	if err != nil {
		http.Error(w, "Failed to fetch detailed tickets: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Retourner les tickets détaillés
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

// GetTicketByUUIDHandler récupère un ticket par son UUID.
func (h *TicketHandler) GetTicketByUUIDHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer le UUID depuis les paramètres de l'URL
	params := mux.Vars(r)
	ticketUUID := params["uuid"] // Pas de conversion nécessaire, le UUID est déjà une chaîne

	// Appeler le service pour récupérer le ticket
	ticket, err := h.service.GetTicketByUUID(ticketUUID)
	if err != nil {
		http.Error(w, "Ticket not found: "+err.Error(), http.StatusNotFound)
		return
	}

	// Retourner le ticket au format JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ticket)
}

// GetEventIDsAndUUIDsHandler gère la requête pour récupérer les eventID et UUID associés à un userID.
func (h *TicketHandler) GetEventIDsAndUUIDsHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer le userID depuis les paramètres de l'URL.
	vars := mux.Vars(r)
	userIDStr, ok := vars["userID"]
	if !ok {
		http.Error(w, "userID is required", http.StatusBadRequest)
		return
	}

	// Convertir le userID en uint.
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}

	// Appeler le service pour récupérer les données.
	results, err := h.service.GetEventIDsAndUUIDsByUserID(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convertir les résultats en JSON et les renvoyer.
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
