package events

import (
    "net/http"
    "encoding/json"
    "backend/internal/services/event"
    "backend/internal/db/models"
    "log"
    "mime/multipart"
    "fmt"
    "time"
    "bytes"
    "io"
    "backend/pkg/httpclient"
)

type CreateEventHandler struct {
    eventService *event.EventService
    httpClient   *httpclient.APIClient
}

func NewCreateEventHandler(eventService *event.EventService, httpClient *httpclient.APIClient) *CreateEventHandler {
    return &CreateEventHandler{eventService: eventService, httpClient: httpClient}
}

func (h *CreateEventHandler) HandleCreateEvent(w http.ResponseWriter, r *http.Request) {
    err := r.ParseMultipartForm(10 << 20) // Limite de 10 Mo
    if err != nil {
        http.Error(w, "Invalid form data", http.StatusBadRequest)
        return
    }

    var event models.Event
    if err := json.NewDecoder(bytes.NewReader([]byte(r.MultipartForm.Value["event"][0]))).Decode(&event); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Gestion des images
    imageURLs := []string{}
    for i := 0; i < 4; i++ {
        fileKey := fmt.Sprintf("image%d", i+1)
        file, header, err := r.FormFile(fileKey)
        if err == nil {
            fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), header.Filename)
            url, err := h.uploadEventImage(file, fileName)
            if err != nil {
                log.Printf("Failed to upload event image: %v", err)
                http.Error(w, "Failed to upload event image", http.StatusInternalServerError)
                return
            }
            imageURLs = append(imageURLs, url)
        }
    }
    if len(imageURLs) > 0 {
        // Convertir le []byte en string
        imageURLsJSON, _ := json.Marshal(imageURLs)
        event.ImageURLs = string(imageURLsJSON)
    }

    if err := h.eventService.CreateEvent(&event); err != nil {
        http.Error(w, "Failed to create event", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(event)
}

func (h *CreateEventHandler) uploadEventImage(file multipart.File, fileName string) (string, error) {
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile("file", fileName)
    if err != nil {
        return "", err
    }
    if _, err := io.Copy(part, file); err != nil {
        return "", err
    }
    writer.Close()

    uploadURL := fmt.Sprintf("%s/upload", h.httpClient.BaseURL)

    req, err := http.NewRequest("POST", uploadURL, body)
    if err != nil {
        return "", err
    }
    req.Header.Set("Content-Type", writer.FormDataContentType())

    resp, err := h.httpClient.DoRequest(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("failed to upload event image: %s", resp.Status)
    }

    var result map[string]string
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return "", err
    }

    return result["url"], nil
}
