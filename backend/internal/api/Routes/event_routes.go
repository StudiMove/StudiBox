package routes
import (
    "net/http"
    "backend/internal/api/handlers/events"
)

func RegisterEventRoutes(createHandler *events.CreateEventHandler, getHandler *events.GetEventHandler, updateHandler *events.UpdateEventHandler, deleteHandler *events.DeleteEventHandler) {
    http.HandleFunc("/events/create", createHandler.HandleCreateEvent)
    http.HandleFunc("/events/get", getHandler.HandleGetEvent)
    http.HandleFunc("/events/update", updateHandler.HandleUpdateEvent)
    http.HandleFunc("/events/delete", deleteHandler.HandleDeleteEvent)
}
