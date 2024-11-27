package events

import (
	"backend/internal/services/event"
	"backend/internal/utils"
	"encoding/json"
	"log"
	"net/http"
)

type GetEventListHandler struct {
    eventService *event.EventService
}

func NewGetEventListHandler(eventService *event.EventService) *GetEventListHandler {
    return &GetEventListHandler{eventService: eventService}
}

func (h *GetEventListHandler) HandleGetEventList(w http.ResponseWriter, r *http.Request) {
    // Récupérer l'utilisateur à partir du contexte (JWT middleware)
    userClaims := r.Context().Value("user")
    if userClaims == nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    claims, ok := userClaims.(*utils.JWTClaims)
    if !ok {
        http.Error(w, "Invalid token claims", http.StatusForbidden)
        return
    }

    // Utiliser le UserID extrait des claims
    userID := claims.UserID

    // Récupérer la liste des événements pour cet utilisateur
    events, err := h.eventService.GetEventsByUserID(uint(userID))
    if err != nil {
        log.Printf("Failed to retrieve events: %v", err)
        http.Error(w, "Failed to retrieve events", http.StatusInternalServerError)
        return
    }

    // Retourner la liste des événements en JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(events)
}

