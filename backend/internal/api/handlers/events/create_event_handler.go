package events

import (
	"backend/internal/api/models/event/request"
	"backend/internal/services/event"
	"backend/internal/utils"
	"encoding/json"
	"log"
	"net/http"
)

type CreateEventHandler struct {
    eventService *event.EventService
}

func NewCreateEventHandler(eventService *event.EventService) *CreateEventHandler {
    return &CreateEventHandler{eventService: eventService}
}

func (h *CreateEventHandler) HandleCreateEvent(w http.ResponseWriter, r *http.Request) {

    var req request.CreateEventRequest

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        log.Printf("Error decoding request payload: %v", err)
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    // log.Println("Request payload decoded successfully")

    userClaims, ok := r.Context().Value("user").(*utils.JWTClaims)
    if !ok || userClaims == nil {
        // log.Println("Unauthorized request: invalid or missing JWT claims")
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    req.UserID = userClaims.UserID // Assignez l'ID de l'utilisateur récupéré

    // log.Printf("User ID retrieved from JWT claims: %d", req.UserID)

    eventID, err := h.eventService.CreateEvent(req)
    if err != nil {
        log.Printf("Error creating event: %v", err)
        http.Error(w, "Failed to create event", http.StatusInternalServerError)
        return
    }

    // log.Printf("Event created successfully with ID: %d", eventID)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "success": true,
        "eventID": eventID,
    })
    log.Println("Response sent to client")
}
