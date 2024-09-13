package authentification

import (
    "net/http"
    "encoding/json"
    "backend/internal/services/auth"
)

type AuthHandler struct {
    authService *auth.AuthService
}

func NewAuthHandler(authService *auth.AuthService) *AuthHandler {
    return &AuthHandler{authService: authService}
}

// HandleLogin gère la connexion des utilisateurs
func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
    var loginData struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    token, err := h.authService.Login(loginData.Email, loginData.Password)
    if err != nil {
        http.Error(w, "Failed to login", http.StatusUnauthorized)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
