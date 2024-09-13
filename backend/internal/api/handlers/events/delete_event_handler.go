package events

import (
    "net/http"
    "backend/internal/services/event"
    "strconv"
)

type DeleteEventHandler struct {
    eventService *event.EventService
}

func NewDeleteEventHandler(eventService *event.EventService) *DeleteEventHandler {
    return &DeleteEventHandler{eventService: eventService}
}

func (h *DeleteEventHandler) HandleDeleteEvent(w http.ResponseWriter, r *http.Request) {
    eventIDStr := r.URL.Query().Get("id")
    eventID, err := strconv.Atoi(eventIDStr)
    if err != nil {
        http.Error(w, "Invalid event ID", http.StatusBadRequest)
        return
    }

    if err := h.eventService.DeleteEvent(uint(eventID)); err != nil {
        http.Error(w, "Failed to delete event", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
