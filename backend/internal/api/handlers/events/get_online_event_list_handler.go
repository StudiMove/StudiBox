package events

import (
	"backend/internal/api/models/event/response"
	"backend/internal/services/event"
	"backend/internal/utils"
	"encoding/json"
	"net/http"
)

type GetOnlineEventListHandler struct {
    eventService *event.EventService
}

func NewGetOnlineEventListHandler(eventService *event.EventService) *GetOnlineEventListHandler {
    return &GetOnlineEventListHandler{eventService: eventService}
}

func (h *GetOnlineEventListHandler) HandleGetOnlineEvents(w http.ResponseWriter, r *http.Request) {
    // Récupérer l'ID utilisateur à partir du token JWT
    userClaims, ok := r.Context().Value("user").(*utils.JWTClaims)
    if !ok || userClaims == nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Récupérer la liste des événements en ligne pour l'utilisateur
    events, err := h.eventService.GetOnlineEventsByUserID(userClaims.UserID)
    if err != nil {
        http.Error(w, "Failed to retrieve events", http.StatusInternalServerError)
        return
    }

    // Transformer les événements en réponses utilisant les ODT de réponse
    eventResponses := make([]response.EventSummaryResponse, len(events))
    for i, event := range events {
        eventResponses[i] = response.EventSummaryResponse{
            ID:        event.ID,
            Title:     event.Title,
            StartDate: event.StartDate,
            StartTime:   event.StartTime,
            IsOnline:  event.IsOnline,
        }
    }

    // Réponse réussie avec la liste des événements
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(eventResponses)
}
