package storage

import "mime/multipart"

// StorageService définit l'interface pour les opérations de stockage
type StorageService interface {
    UploadFile(file multipart.File, fileName string) (string, error) // Retourne l'URL de l'objet
    DeleteFile(fileName string) error                                // Supprime le fichier
    GetFileURL(fileName string) string                               // Retourne l'URL du fichier
}
