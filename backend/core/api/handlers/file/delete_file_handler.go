package file

import (
	"backend/core/services/storage"
	"log"
	"net/http"
)

// DeleteFileHandler gère la suppression des fichiers
type DeleteFileHandler struct {
	storage storage.StorageService
}

// NewDeleteFileHandler crée une nouvelle instance de DeleteFileHandler
func NewDeleteFileHandler(storage storage.StorageService) *DeleteFileHandler {
	return &DeleteFileHandler{storage: storage}
}

// HandleDelete gère la suppression des fichiers
func (h *DeleteFileHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path[len("/files/"):]

	err := h.storage.DeleteFile(fileName)
	if err != nil {
		log.Printf("Failed to delete file: %v", err)
		http.Error(w, "Failed to delete file", http.StatusInternalServerError)
		return
	}

	log.Printf("File deleted successfully: %s", fileName)
	w.Write([]byte("File deleted successfully"))
}
