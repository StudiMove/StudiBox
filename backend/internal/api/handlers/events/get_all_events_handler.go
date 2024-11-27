package events

import (
	"encoding/json"
	"log"
	"net/http"

	// Import des ODT pour structurer la réponse
	"backend/internal/api/models/event/response"
	"backend/internal/services/event"
)

type GetAllEventsHandler struct {
    eventService *event.EventService
}

func NewGetAllEventsHandler(eventService *event.EventService) *GetAllEventsHandler {
    return &GetAllEventsHandler{eventService: eventService}
}
func (h *GetAllEventsHandler) HandleGetAllEvents(w http.ResponseWriter, r *http.Request) {
    // Récupérer tous les événements via le service
    events, err := h.eventService.GetAllEvents()
    if err != nil {
        log.Printf("Failed to retrieve events: %v", err)
        http.Error(w, "Failed to retrieve events", http.StatusInternalServerError)
        return
    }

    eventResponses := make([]response.EventSummaryResponse, len(events))

    for i, event := range events {
        imageURLs := event.ImageURLs 
        
        eventResponses[i] = response.EventSummaryResponse{
            ID:        event.ID,
            Title:     event.Title,
            StartDate: event.StartDate,
            StartTime: event.StartTime,
            IsOnline:  event.IsOnline,
            ImageURLs: imageURLs,
        }
    }

    // Retourner la liste des événements en JSON
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(eventResponses)
}

