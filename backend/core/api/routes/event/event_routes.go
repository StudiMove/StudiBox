package event_routes

import (
	eventHandlers "backend/core/api/handlers/event"
	"backend/core/api/middleware"
	"backend/core/services/event"
	"backend/core/services/owner"
	"backend/core/services/user"

	"github.com/gin-gonic/gin"
)

// EventRoutes enregistre les routes des événements avec le service d'événements
func EventRoutes(routerGroup *gin.RouterGroup, eventService *event.EventServiceType, ownerService *owner.OwnerServiceType, userService *user.UserServiceType) {
	eventGroup := routerGroup.Group("/events")
	eventGroup.Use(middleware.AuthMiddleware())
	{
		eventGroup.POST("/", func(c *gin.Context) {
			eventHandlers.HandleCreateEvent(c, eventService, ownerService)
		})
		eventGroup.GET("/:id", func(c *gin.Context) {
			eventHandlers.HandleGetEvent(c, eventService)
		})
		eventGroup.PUT("/:id", func(c *gin.Context) {
			eventHandlers.HandleUpdateEvent(c, eventService, ownerService)
		})
		eventGroup.DELETE("/:id", middleware.RoleMiddleware(userService, []string{"Admin", "Owner"}), func(c *gin.Context) {
			eventHandlers.HandleDeleteEvent(c, eventService, userService)
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
