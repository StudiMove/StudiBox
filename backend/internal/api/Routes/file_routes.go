package routes

import (
    "net/http"
    "backend/internal/services/storage"
    "backend/internal/api/handlers/file"
)

func RegisterFileRoutes(storageService storage.StorageService) {
    fileUploadHandler := file.NewUploadFileHandler(storageService)
    fileDeleteHandler := file.NewDeleteFileHandler(storageService)
    fileGetURLHandler := file.NewGetFileURLHandler(storageService)

    http.HandleFunc("/upload", fileUploadHandler.HandleUpload)
    http.HandleFunc("/files/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodDelete:
            fileDeleteHandler.HandleDelete(w, r)
        case http.MethodGet:
            fileGetURLHandler.HandleGetFileURL(w, r)
        default:
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        }
    })
}
