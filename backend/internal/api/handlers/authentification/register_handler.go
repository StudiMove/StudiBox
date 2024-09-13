package authentification

import (
    "net/http"
    "encoding/json"
    "backend/internal/services/auth"
    "backend/internal/db/models"
    "log"
    "mime/multipart"
    "fmt"
    "time"
    "bytes"
    "io"
    "backend/pkg/httpclient"
)

type RegisterHandler struct {
    authService *auth.AuthService
    httpClient  *httpclient.APIClient
}

func NewRegisterHandler(authService *auth.AuthService, httpClient *httpclient.APIClient) *RegisterHandler {
    return &RegisterHandler{authService: authService, httpClient: httpClient}
}

// HandleRegisterUser gère l'inscription des utilisateurs normaux
func (h *RegisterHandler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
    err := r.ParseMultipartForm(10 << 20) // Limite de 10 Mo
    if err != nil {
        http.Error(w, "Invalid form data", http.StatusBadRequest)
        return
    }

    var user models.User
    if err := json.NewDecoder(bytes.NewReader([]byte(r.MultipartForm.Value["user"][0]))).Decode(&user); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Laisser le pseudo vide pour l'instant
    user.Pseudo = "" // Ou ne rien faire pour le laisser vide

    // Gestion de l'image de profil
    file, header, err := r.FormFile("profile_image")
    if err == nil {
        fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), header.Filename)
        url, err := h.uploadProfileImage(file, fileName)
        if err != nil {
            log.Printf("Failed to upload profile image: %v", err)
            http.Error(w, "Failed to upload profile image", http.StatusInternalServerError)
            return
        }
        user.ProfileImage = url
    }

    if err := h.authService.RegisterUser(&user); err != nil {
        http.Error(w, "Failed to register user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

// HandleRegisterBusinessUser gère l'inscription des utilisateurs entreprises
func (h *RegisterHandler) HandleRegisterBusinessUser(w http.ResponseWriter, r *http.Request) {
    err := r.ParseMultipartForm(10 << 20) // Limite de 10 Mo
    if err != nil {
        http.Error(w, "Invalid form data", http.StatusBadRequest)
        return
    }

    var businessUser models.BusinessUser
    if err := json.NewDecoder(bytes.NewReader([]byte(r.MultipartForm.Value["business_user"][0]))).Decode(&businessUser); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Gestion de l'image de profil
    file, header, err := r.FormFile("profile_image")
    if err == nil {
        fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), header.Filename)
        url, err := h.uploadProfileImage(file, fileName)
        if err != nil {
            log.Printf("Failed to upload profile image: %v", err)
            http.Error(w, "Failed to upload profile image", http.StatusInternalServerError)
            return
        }
        businessUser.User.ProfileImage = url
    }

    if err := h.authService.RegisterBusinessUser(&businessUser); err != nil {
        http.Error(w, "Failed to register business user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(businessUser)
}

// uploadProfileImage gère l'appel à la route d'upload pour télécharger l'image de profil
func (h *RegisterHandler) uploadProfileImage(file multipart.File, fileName string) (string, error) {
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

    // Utilisez l'URL de base correctement configurée
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
        return "", fmt.Errorf("failed to upload profile image: %s", resp.Status)
    }

    var result map[string]string
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return "", err
    }

    return result["url"], nil
}
