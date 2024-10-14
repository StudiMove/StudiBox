package authentification

import (
    "net/http"
    "encoding/json"
    "backend/internal/services/auth"
)

type AuthHandler struct {
    authService *auth.AuthService
}

// NewAuthHandler crée une nouvelle instance de AuthHandler
func NewAuthHandler(authService *auth.AuthService) *AuthHandler {
    return &AuthHandler{authService: authService}
}

// HandleLogin gère la connexion des utilisateurs
func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
    // Structure pour les données de connexion
    var loginData struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    // Décode les données JSON reçues
    if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Appelle le service d'authentification pour se connecter
    token, err := h.authService.Login(loginData.Email, loginData.Password)
    if err != nil {
        http.Error(w, "Failed to login", http.StatusUnauthorized)
        return
    }

    // Répond avec le token en cas de succès
    w.WriteHeader(http.StatusOK)
    // Répond avec le token et le statut d'authentification
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "token": token,
        "isAuthenticated": true, // Indique que l'utilisateur est authentifié
    })

}
