package events

import (
    "net/http"
    "backend/internal/services/event"
    "strconv"
    "encoding/json"
)

type GetEventHandler struct {
    eventService *event.EventService
}

func NewGetEventHandler(eventService *event.EventService) *GetEventHandler {
    return &GetEventHandler{eventService: eventService}
}

func (h *GetEventHandler) HandleGetEvent(w http.ResponseWriter, r *http.Request) {
    eventIDStr := r.URL.Query().Get("id")
    eventID, err := strconv.Atoi(eventIDStr)
    if err != nil {
        http.Error(w, "Invalid event ID", http.StatusBadRequest)
        return
    }

    event, err := h.eventService.GetEvent(uint(eventID))
    if err != nil {
        http.Error(w, "Event not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(event)
}
