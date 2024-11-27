package events

import (
	"backend/internal/api/models/event/request"
	"backend/internal/api/models/event/response"
	"backend/internal/services/event"
	"encoding/json"
	"net/http"
)

type DeleteEventHandler struct {
    eventService *event.EventService
}

func NewDeleteEventHandler(eventService *event.EventService) *DeleteEventHandler {
    return &DeleteEventHandler{eventService: eventService}
}

func (h *DeleteEventHandler) HandleDeleteEvent(w http.ResponseWriter, r *http.Request) {
    var req request.DeleteEventRequest

    // Décoder la requête JSON en utilisant l'ODT
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if req.TargetEventID == 0 {
        http.Error(w, "Event ID is required", http.StatusBadRequest)
        return
    }

    // Appel au service pour supprimer l'événement
    if err := h.eventService.DeleteEvent(req.TargetEventID); err != nil {
        http.Error(w, "Failed to delete event", http.StatusInternalServerError)
        return
    }

    // Structure de réponse en utilisant l'ODT
    res := response.DeleteEventResponse{
        Message: "Event deleted successfully",
    }

    // Réponse JSON de succès
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(res)
}
