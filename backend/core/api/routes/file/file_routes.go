package file_routes

import (
	"backend/core/api/handlers/file"

	"github.com/gin-gonic/gin"
)

// RegisterFileRoutes enregistre les routes liées aux fichiers
func RegisterFileRoutes(routerGroup *gin.RouterGroup) {
	fileGroup := routerGroup.Group("/files")
	{
		fileGroup.POST("/upload", file.HandleUpload)
		fileGroup.DELETE("/:id", file.HandleDeleteFile)
		fileGroup.GET("/:id", file.HandleGetFileURL)
	}
}
