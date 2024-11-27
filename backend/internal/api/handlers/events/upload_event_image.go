package events

import (
	"backend/internal/api/models/event/response"
	"backend/internal/services/event"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"time"
)

type UploadEventImageHandler struct {
	eventService *event.EventService
}

// NewUploadEventImageHandler initialise et retourne un nouveau UploadEventImageHandler
func NewUploadEventImageHandler(eventService *event.EventService) *UploadEventImageHandler {
	return &UploadEventImageHandler{eventService: eventService}
}
func (h *UploadEventImageHandler) HandleUploadEventImage(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	fileHeaders := r.MultipartForm.File["file"]
	var files []multipart.File
	var fileNames []string

	for _, fileHeader := range fileHeaders {
		file, err := fileHeader.Open()
		if err != nil {
			continue
		}
		defer file.Close()
		files = append(files, file)
		fileNames = append(fileNames, fmt.Sprintf("event_image_%d_%s", time.Now().Unix(), fileHeader.Filename))
	}

	results, err := h.eventService.UploadEventImages(files, fileNames)
	if err != nil {
		http.Error(w, "Failed to upload event images", http.StatusInternalServerError)
		return
	}

	var imageUrls []string
	for _, result := range results {
		imageUrls = append(imageUrls, result.URL)
	}

	res := response.UploadEventImageResponse{URLs: imageUrls}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}