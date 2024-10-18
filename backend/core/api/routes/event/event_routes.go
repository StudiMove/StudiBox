package event_routes

import (
	eventHandlers "backend/core/api/handlers/events"
	"backend/core/api/middleware"
	"backend/core/services/event"

	"github.com/gin-gonic/gin"
)

// RegisterEventRoutes enregistre les routes des événements avec le service d'événements
func RegisterEventRoutes(routerGroup *gin.RouterGroup, eventService *event.EventService) {
	eventGroup := routerGroup.Group("/events")
	eventGroup.Use(middleware.AuthMiddleware())
	{
		eventGroup.POST("/", func(c *gin.Context) {
			eventHandlers.HandleCreateEvent(c, eventService)
		})
		eventGroup.GET("/:id", func(c *gin.Context) {
			eventHandlers.HandleGetEvent(c, eventService)
		})
		eventGroup.PUT("/:id", func(c *gin.Context) {
			eventHandlers.HandleUpdateEvent(c, eventService)
		})
		eventGroup.DELETE("/:id", func(c *gin.Context) {
			eventHandlers.HandleDeleteEvent(c, eventService)
		})
	}
}
