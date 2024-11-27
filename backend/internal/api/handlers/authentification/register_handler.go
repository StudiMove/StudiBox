package authentification

import (
	"backend/internal/api/models/auth/request"
	"backend/internal/services/auth"
	"backend/pkg/httpclient"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type RegisterHandler struct {
	authService *auth.AuthService
	httpClient  *httpclient.APIClient
}

func NewRegisterHandler(authService *auth.AuthService, httpClient *httpclient.APIClient) *RegisterHandler {
	return &RegisterHandler{authService: authService, httpClient: httpClient}
}
func (h *RegisterHandler) uploadProfileImages(files []multipart.File, fileNames []string) (string, error) {
	var urls []string

	for i, file := range files {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", fileNames[i])
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

		urls = append(urls, result["url"])
	}

	// Convertir le tableau d'URL en JSON string
	jsonUrls, err := json.Marshal(urls)
	if err != nil {
		return "", err
	}

	return string(jsonUrls), nil
}

// HandleRegisterUser gère l'inscription des utilisateurs normaux
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

func (h *RegisterHandler) HandleRegisterOrganisationUser(w http.ResponseWriter, r *http.Request) {
	var registerReq request.RegisterOrganisationUserRequest

	// Décoder la requête JSON
	if err := json.NewDecoder(r.Body).Decode(&registerReq); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validation des champs requis
	if registerReq.Email == "" || registerReq.Password == "" || registerReq.OrganisationName == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Enregistrer l'organisation
	registerResp, err := h.authService.RegisterOrganisationUser(&registerReq)
	if err != nil {
		http.Error(w, "Failed to register organisation", http.StatusInternalServerError)
		return
	}

	// Créez la requête pour obtenir l'ID du rôle
	roleIDReq := &request.RoleIDRequest{RoleName: registerReq.OrganisationType}
	roleIDResp, err := h.authService.GetRoleIDByName(roleIDReq)
	if err != nil {
		http.Error(w, "Failed to get role ID", http.StatusInternalServerError)
		return
	}

	// Créez la requête pour assigner le rôle
	assignRoleReq := &request.AssignUserRoleRequest{
		UserID: registerResp.UserID,
		RoleID: roleIDResp.RoleID,
	}

	// Assigner le rôle
	if _, err := h.authService.AssignUserRole(assignRoleReq); err != nil {
		http.Error(w, "Failed to assign user role", http.StatusInternalServerError)
		return
	}

	// Réponse de succès
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(registerResp)
}

func (h *RegisterHandler) HandleRegisterNormalUser(w http.ResponseWriter, r *http.Request) {
	var registerReq request.RegisterNormalUserRequest

	// Décoder la requête JSON
	if err := json.NewDecoder(r.Body).Decode(&registerReq); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validation des champs requis
	if registerReq.FirstName == "" || registerReq.LastName == "" || registerReq.Email == "" || registerReq.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Enregistrer l'utilisateur normal
	registerResp, err := h.authService.RegisterNormalUser(&registerReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to register user: %v", err), http.StatusInternalServerError)
		return
	}

	// Réponse de succès
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(registerResp)
}
