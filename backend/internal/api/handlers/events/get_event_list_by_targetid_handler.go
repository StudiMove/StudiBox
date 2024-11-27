package events

import (
	"backend/internal/api/models/event/request"
	"backend/internal/services/event"
	"encoding/json"
	"log"
	"net/http"
)

type GetEventListByTargetIDHandler struct {
    eventService *event.EventService
}

func NewGetEventListByTargetIDHandler(eventService *event.EventService) *GetEventListByTargetIDHandler {
    return &GetEventListByTargetIDHandler{eventService: eventService}
}

// HandleGetEventListByTargetID récupère la liste des événements pour un UserTargetID donné
func (h *GetEventListByTargetIDHandler) HandleGetEventListByTargetID(w http.ResponseWriter, r *http.Request) {
    var reqBody request.GetEventListByTargetIDRequest

    // Décoder le corps de la requête pour obtenir UserTargetID
    if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    log.Printf("Received userTargetID: %d", reqBody.UserTargetID)

    // Récupérer la liste des événements pour cet utilisateur
    events, err := h.eventService.GetEventsByUserID(reqBody.UserTargetID)
    if err != nil {
        http.Error(w, "Failed to get events for user", http.StatusInternalServerError)
        return
    }

    // Retourner les événements sous forme de JSON
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(events)
}
