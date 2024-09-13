// package routes

// import (
//     "net/http"
//     "backend/internal/api/handlers/events"
//     "backend/internal/api/middlewares"
// )

// func RegisterEventRoutes(createHandler *events.CreateEventHandler, getHandler *events.GetEventHandler, updateHandler *events.UpdateEventHandler, deleteHandler *events.DeleteEventHandler) {
//     http.Handle("/events/create", middleware.AuthMiddleware(http.HandlerFunc(createHandler.HandleCreateEvent)))
//     http.Handle("/events/get", middleware.AuthMiddleware(http.HandlerFunc(getHandler.HandleGetEvent)))
//     http.Handle("/events/update", middleware.AuthMiddleware(http.HandlerFunc(updateHandler.HandleUpdateEvent)))
//     http.Handle("/events/delete", middleware.AuthMiddleware(http.HandlerFunc(deleteHandler.HandleDeleteEvent)))
// }
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
