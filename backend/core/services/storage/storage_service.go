package storage

import (
	"errors"
	"mime/multipart"
)

// Définition des erreurs de stockage
var (
	ErrFileNotFound = errors.New("file not found")
	ErrUploadFailed = errors.New("upload failed")
	ErrDeleteFailed = errors.New("delete failed")
)

// StorageService définit l'interface pour les opérations de stockage
type StorageService interface {
	UploadFile(file multipart.File, fileName string) (string, error) // Retourne l'URL de l'objet
	DeleteFile(fileName string) error                                // Supprime le fichier
	GetFileURL(fileName string) string                               // Retourne l'URL du fichier
}
