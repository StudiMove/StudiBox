package events

import (
	"encoding/json"
	"log"
	"net/http"

	"backend/internal/api/models/event/response"
	"backend/internal/services/event"
)

type GetOnlineEventsHandler struct {
    eventService *event.EventService
}

func NewGetOnlineEventsHandler(eventService *event.EventService) *GetOnlineEventsHandler {
    return &GetOnlineEventsHandler{eventService: eventService}
}

func (h *GetOnlineEventsHandler) HandleGetOnlineEvents(w http.ResponseWriter, r *http.Request) {
    // Récupérer les événements en ligne via le service
    events, err := h.eventService.GetAllOnlineEvents()
    if err != nil {
        log.Printf("Failed to retrieve online events: %v", err)
        http.Error(w, "Failed to retrieve online events", http.StatusInternalServerError)
        return
    }
    
    eventResponses := make([]response.EventSummaryResponse, len(events))
    for i, event := range events {
        eventResponses[i] = response.EventSummaryResponse{
            ID:        event.ID,
            Title:     event.Title,
            StartDate: event.StartDate,
            StartTime:   event.StartTime,
            IsOnline:   event.IsOnline,
            ImageURLs: event.ImageURLs,

        }
    }

    // Retourner la liste des événements en ligne en JSON
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(eventResponses)
}
