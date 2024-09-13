package events

import (
    "net/http"
    "encoding/json"
    "backend/internal/services/event"
    "backend/internal/db/models"
)

type UpdateEventHandler struct {
    eventService *event.EventService
}

func NewUpdateEventHandler(eventService *event.EventService) *UpdateEventHandler {
    return &UpdateEventHandler{eventService: eventService}
}

func (h *UpdateEventHandler) HandleUpdateEvent(w http.ResponseWriter, r *http.Request) {
    var event models.Event
    if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    if err := h.eventService.UpdateEvent(&event); err != nil {
        http.Error(w, "Failed to update event", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(event)
}
