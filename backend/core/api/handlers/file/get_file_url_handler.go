package file

import (
	"backend/core/services/storage"
	"log"
	"net/http"
)

// GetFileURLHandler gère la récupération de l'URL des fichiers
type GetFileURLHandler struct {
	storage storage.StorageService
}

// NewGetFileURLHandler crée une nouvelle instance de GetFileURLHandler
func NewGetFileURLHandler(storage storage.StorageService) *GetFileURLHandler {
	return &GetFileURLHandler{storage: storage}
}

// HandleGetFileURL gère la récupération de l'URL d'un fichier
func (h *GetFileURLHandler) HandleGetFileURL(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path[len("/files/"):]

	url := h.storage.GetFileURL(fileName)
	log.Printf("File URL: %s", url)
	w.Write([]byte("File URL: " + url))
}
