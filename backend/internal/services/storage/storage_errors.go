package storage

import "errors"

// Définition des erreurs de stockage
var (
    ErrFileNotFound = errors.New("file not found")
    ErrUploadFailed = errors.New("upload failed")
    ErrDeleteFailed = errors.New("delete failed")
)
