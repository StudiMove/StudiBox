package request

import "mime/multipart"

// ProfileImageUploadRequest représente la requête pour télécharger une image de profil.
type ProfileImageUploadRequest struct {
    File     multipart.File `json:"-"`
    FileName string         `json:"fileName"`
    UserID   uint           `json:"userId"`
}