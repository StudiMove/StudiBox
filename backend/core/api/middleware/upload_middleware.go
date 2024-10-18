package middleware

import (
	"backend/core/services/storage"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadFileMiddleware gère l'upload des images pour les événements
func UploadFileMiddleware(storageService storage.StorageService) gin.HandlerFunc {
	return func(c *gin.Context) {
		imageURLs := []string{}

		// Gestion des images d'événements (jusqu'à 4 images)
		for i := 0; i < 4; i++ {
			fileKey := fmt.Sprintf("image%d", i+1)
			header, err := c.FormFile(fileKey) // Récupère l'en-tête du fichier
			if err == nil && header != nil {   // Si le fichier est bien récupéré
				// Ouvrir le fichier via l'en-tête
				file, err := header.Open()
				if err != nil {
					log.Printf("Failed to open file: %v", err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
					c.Abort()
					return
				}
				defer file.Close()

				// Génère un nom de fichier unique avec l'horodatage
				fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), header.Filename)
				// Téléverse l'image
				url, err := uploadFileToStorage(storageService, file, fileName)
				if err != nil {
					log.Printf("Failed to upload image: %v", err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
					c.Abort()
					return
				}
				// Ajouter l'URL de l'image à la liste
				imageURLs = append(imageURLs, url)
			}
		}

		// Ajouter les URLs des images dans le contexte pour les handlers
		c.Set("image_urls", imageURLs)
		c.Next()
	}
}

// uploadFileToStorage utilise le service de stockage pour téléverser les fichiers
func uploadFileToStorage(storageService storage.StorageService, file multipart.File, fileName string) (string, error) {
	url, err := storageService.UploadFile(file, fileName)
	if err != nil {
		return "", err
	}
	return url, nil
}
