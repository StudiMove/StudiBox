package event_routes

import (
	eventHandlers "backend/core/api/handlers/events"
	"backend/core/api/middleware"
	"backend/core/services/event"

	"github.com/gin-gonic/gin"
)

// EventRoutes enregistre les routes des événements avec le service d'événements
func EventRoutes(routerGroup *gin.RouterGroup, eventService *event.EventService) {
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
		eventGroup.GET("/all", func(c *gin.Context) {
			eventHandlers.HandleListEvents(c, eventService)
		})
		eventGroup.GET("/recommendations", func(c *gin.Context) {
			eventHandlers.HandleGetRecommendations(c, eventService)
		})
		eventGroup.POST("/like/:id", func(c *gin.Context) {
			eventHandlers.HandleLikeEvent(c, eventService)
		})
		eventGroup.DELETE("/like/:id", func(c *gin.Context) {
			eventHandlers.HandleUnlikeEvent(c, eventService)
		})
		eventGroup.GET("/liked", func(c *gin.Context) {
			eventHandlers.HandleGetLikedEvents(c, eventService)
		})
	}
}
