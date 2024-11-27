package events

import (
	"backend/internal/api/models/event/request"
	"backend/internal/services/event"
	"encoding/json"
	"net/http"
)

type GetEventHandler struct {
	eventService *event.EventService
}

func NewGetEventHandler(eventService *event.EventService) *GetEventHandler {
	return &GetEventHandler{eventService: eventService}
}


func (h *GetEventHandler) HandleGetEvent(w http.ResponseWriter, r *http.Request) {
	var req request.GetEventRequest

	// Décoder le body de la requête
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Récupérer les détails de l'événement
	event, err := h.eventService.GetEventByID(req.TargetEventID)
	if err != nil {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	// Réponse réussie avec l'événement
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}