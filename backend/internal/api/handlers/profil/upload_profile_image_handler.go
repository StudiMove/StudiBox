package profil

import (
	"backend/internal/api/models/upload/request"
	"backend/internal/api/models/upload/response"
	"backend/internal/db/models"
	"backend/internal/services/profilservice"
	"backend/internal/services/storage"
	"backend/internal/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type UploadProfilImageHandler struct {
    profilService *profilservice.ProfilService  // Utiliser profilservice.ProfilService
    jwtSecret     string
}

// Changez le type de l'argument de profilService pour correspondre
func NewUploadProfilImageHandler(profilService *profilservice.ProfilService, jwtSecret string) *UploadProfilImageHandler {
    return &UploadProfilImageHandler{
        profilService: profilService,
        jwtSecret:     jwtSecret,
    }
}



type ProfilService struct {
    storageService storage.StorageService
    db             *gorm.DB
}

func NewProfilService(storageService storage.StorageService, db *gorm.DB) *ProfilService {
    return &ProfilService{
        storageService: storageService,
        db:             db,
    }
}

// HandleUploadProfileImage gère l'upload d'image de profil pour l'utilisateur authentifié
func (h *UploadProfilImageHandler) HandleUploadProfileImage(w http.ResponseWriter, r *http.Request) {
    log.Println("Received request to upload profile image")

    // Extraire et valider le token
    tokenStr := r.Header.Get("Authorization")
    if tokenStr == "" {
        http.Error(w, "Authorization token is required", http.StatusUnauthorized)
        return
    }

    if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
        tokenStr = tokenStr[7:]
    }

    userID, err := utils.ExtractUserIDFromToken(tokenStr, h.jwtSecret)
    if err != nil {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }

    h.handleUpload(w, r, userID)
}

// HandleUploadProfileImageWithTargetID gère l'upload d'image de profil pour un utilisateur spécifique via targetId

func (h *UploadProfilImageHandler) HandleUploadProfileImageWithTargetID(w http.ResponseWriter, r *http.Request) {
    log.Println("Received request to upload profile image with targetId")

    // Extraire et valider le token
    tokenStr := r.Header.Get("Authorization")
    if tokenStr == "" {
        http.Error(w, "Authorization token is required", http.StatusUnauthorized)
        return
    }

    if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
        tokenStr = tokenStr[7:]
    }

    userID, err := utils.ExtractUserIDFromToken(tokenStr, h.jwtSecret)
    if err != nil {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }

    log.Printf("UserID from claims: %d\n", userID)
    log.Println("Reading request body for targetId and file")

    // Traiter le formulaire multi-part
    if err := r.ParseMultipartForm(10 << 20); err != nil {
        http.Error(w, "Invalid form data", http.StatusBadRequest)
        return
    }

    // Récupérer le targetId
    targetIDStr := r.FormValue("targetId")
    if targetIDStr == "" {
        http.Error(w, "targetId is required", http.StatusBadRequest)
        return
    }

    targetID, err := strconv.ParseUint(targetIDStr, 10, 32)
    if err != nil {
        log.Printf("Failed to parse targetId: %v", err)
        http.Error(w, "Invalid targetId", http.StatusBadRequest)
        return
    }
    log.Printf("Received targetId: %d\n", targetID)

    // Récupérer le fichier
    file, header, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Failed to get file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    log.Printf("Received file: %s\n", header.Filename)

    fileName := fmt.Sprintf("profile_image_%d_%s", time.Now().Unix(), header.Filename)

    // Créer la requête DTO
    uploadReq := request.ProfileImageUploadRequest{
        File:     file,
        FileName: fileName,
        UserID:   uint(targetID),
    }

    // Appeler le service pour l'upload
    url, err := h.profilService.UploadProfileImage(&uploadReq)
    if err != nil {
        http.Error(w, "Failed to upload profile image", http.StatusInternalServerError)
        return
    }

    // Mettre à jour l'URL de l'image dans la base de données
    if err := h.profilService.UpdateUserProfileImage(uint(targetID), url); err != nil {
        http.Error(w, "Failed to update user profile image", http.StatusInternalServerError)
        return
    }

    // Préparer et envoyer la réponse
    resp := response.ProfileImageUploadResponse{
        URL:     url,
        Success: true,
        Message: "Profile image uploaded successfully",
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

// handleUpload est une fonction générique pour gérer le téléchargement d'image de profil
func (h *UploadProfilImageHandler) handleUpload(w http.ResponseWriter, r *http.Request, userID uint) {
    // Traiter le formulaire multi-part
    if err := r.ParseMultipartForm(10 << 20); err != nil {
        http.Error(w, "Invalid form data", http.StatusBadRequest)
        return
    }

    // Récupérer le fichier
    file, header, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Failed to get file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    fileName := fmt.Sprintf("profile_image_%d_%s", time.Now().Unix(), header.Filename)

    // Créer la requête DTO
    uploadReq := request.ProfileImageUploadRequest{
        File:     file,
        FileName: fileName,
        UserID:   userID,
    }

    // Appeler le service pour l'upload
    url, err := h.profilService.UploadProfileImage(&uploadReq)
    if err != nil {
        http.Error(w, "Failed to upload profile image", http.StatusInternalServerError)
        return
    }

    // Mettre à jour l'URL de l'image dans la base de données
    if err := h.profilService.UpdateUserProfileImage(userID, url); err != nil {
        http.Error(w, "Failed to update user profile image", http.StatusInternalServerError)
        return
    }

    // Préparer et envoyer la réponse
    resp := response.ProfileImageUploadResponse{
        URL:     url,
        Success: true,
        Message: "Profile image uploaded successfully",
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

// UploadProfileImage gère le téléchargement du fichier de profil via le storage service
func (s *ProfilService) UploadProfileImage(req *request.ProfileImageUploadRequest) (string, error) {
    log.Printf("Starting file upload for file: %s", req.FileName)

    url, err := s.storageService.UploadFile(req.File, req.FileName)
    if err != nil {
        log.Printf("Error while uploading file to storage service: %v", err)
        return "", storage.ErrUploadFailed
    }

    log.Printf("File uploaded successfully. URL: %s", url)
    return url, nil
}

// UpdateUserProfileImage met à jour l'URL de l'image de profil dans la base de données
func (s *ProfilService) UpdateUserProfileImage(userID uint, imageURL string) error {
    var user models.User
    if err := s.db.Model(&user).Where("id = ?", userID).Update("profile_image", imageURL).Error; err != nil {
        return err
    }
    return nil
}
