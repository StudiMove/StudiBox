package file

import (
	"backend/core/services/storage"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

const maxFileSize = 5 * 1024 * 1024 // 5 Mo en octets
var allowedImageExtensions = []string{".jpg", ".jpeg", ".png", ".gif", ".PNG"}

func isImageAllowed(fileName string) bool {
	ext := filepath.Ext(fileName)
	for _, allowedExt := range allowedImageExtensions {
		if ext == allowedExt {
			return true
		}
	}
	return false
}

// UploadFileHandler gère le téléchargement des fichiers
type UploadFileHandler struct {
	storage storage.StorageService
}

// NewUploadFileHandler crée une nouvelle instance de UploadFileHandler
func NewUploadFileHandler(storage storage.StorageService) *UploadFileHandler {
	return &UploadFileHandler{storage: storage}
}

// HandleUpload gère le téléchargement des fichiers
func (h *UploadFileHandler) HandleUpload(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling file upload...")

	file, header, err := r.FormFile("file")
	if err != nil {
		log.Printf("Failed to get file from form: %v", err)
		http.Error(w, "Failed to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Vérifier la taille du fichier
	if header.Size > maxFileSize {
		log.Printf("File size exceeds the limit of 5 MB")
		http.Error(w, "File size exceeds the limit of 5 MB", http.StatusBadRequest)
		return
	}

	// Vérifier le format de l'image
	if !isImageAllowed(header.Filename) {
		log.Printf("Invalid file type: %s", header.Filename)
		http.Error(w, "Invalid file type. Only images are allowed.", http.StatusBadRequest)
		return
	}

	// Générer un nom de fichier unique
	fileName := fmt.Sprintf("%d%s", time.Now().Unix(), filepath.Ext(header.Filename))
	log.Printf("Generated file name: %s", fileName)

	// Télécharger le fichier vers le service de stockage
	url, err := h.storage.UploadFile(file, fileName)
	if err != nil {
		log.Printf("Failed to upload file to storage: %v", err)
		http.Error(w, "Failed to upload file", http.StatusInternalServerError)
		return
	}

	log.Printf("File uploaded successfully: %s", url)

	// Retourner l'URL du fichier en JSON
	response := map[string]string{"url": url}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
