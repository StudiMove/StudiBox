package events

import (
	"backend/internal/services/event"
	"encoding/json"
	"log"
	"net/http"
)

type GetEventWithTagsHandler struct {
	eventService *event.EventService
}

func NewGetEventWithTagsHandler(eventService *event.EventService) *GetEventWithTagsHandler {
	return &GetEventWithTagsHandler{eventService: eventService}
}

// HandleGetEventWithTags récupère un événement, y compris les tags associés
func (h *GetEventWithTagsHandler) HandleGetEventWithTags(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		EventID uint `json:"eventId"`
	}

	// Décoder la requête JSON
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Récupérer l'événement par ID
	event, err := h.eventService.GetEventByID(reqBody.EventID)
	if err != nil {
		log.Printf("Error fetching event: %v", err)
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	// Retourner l'événement et les tags associés en JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}
