package events

import (
	"backend/internal/api/models/event/request"
	"backend/internal/api/models/event/response"
	"backend/internal/services/event"
	"encoding/json"
	"net/http"
	"strconv"
)

type EventUserActionHandler struct {
	service *event.EventUserActionService
}

func NewEventUserActionHandler(service *event.EventUserActionService) *EventUserActionHandler {
	return &EventUserActionHandler{service: service}
}
func (h *EventUserActionHandler) UpdateUserAction(w http.ResponseWriter, r *http.Request) {
	var req request.UpdateUserActionRequest

	// Décoder la requête JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Appeler le service pour mettre à jour l'action utilisateur
	err := h.service.UpdateUserAction(req.UserID, req.EventID, req.IsInterested, req.IsFavorite, req.UpdateInterested, req.UpdateFavorite)
	if err != nil {
		http.Error(w, "Failed to update user action", http.StatusInternalServerError)
		return
	}

	// Retourner une réponse de succès
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.ActionUpdateResponse{
		Success: true,
		Message: "User action updated successfully",
	})
}

// Handler pour récupérer les événements favoris d'un utilisateur
func (h *EventUserActionHandler) GetFavoriteEvents(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	events, err := h.service.GetFavoriteEventsByUser(uint(userID))
	if err != nil {
		http.Error(w, "Failed to fetch favorite events", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(events)
}

// Handler pour récupérer les événements où l'utilisateur est intéressé
func (h *EventUserActionHandler) GetInterestedEvents(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	events, err := h.service.GetInterestedEventsByUser(uint(userID))
	if err != nil {
		http.Error(w, "Failed to fetch interested events", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(events)
}

// Handler pour supprimer une action utilisateur
func (h *EventUserActionHandler) RemoveUserAction(w http.ResponseWriter, r *http.Request) {
	var req request.RemoveUserActionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := h.service.RemoveUserAction(req.UserID, req.EventID)
	if err != nil {
		http.Error(w, "Failed to remove user action", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.ActionRemoveResponse{
		Success: true,
		Message: "User action removed successfully",
	})
}
