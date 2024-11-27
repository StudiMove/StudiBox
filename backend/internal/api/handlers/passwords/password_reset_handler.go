package passwords

import (
	"encoding/json"
	"net/http"

	"backend/internal/api/models/password/request"
	"backend/internal/api/models/password/response"
	"backend/internal/db/models"
	"backend/internal/services/auth"
	"backend/internal/services/password"
	"time"

	"gorm.io/gorm"
)

type PasswordResetHandler struct {
	passwordResetService *password.PasswordResetService
	authService          *auth.AuthService
	db                   *gorm.DB
}

func NewPasswordResetHandler(passwordResetService *password.PasswordResetService, authService *auth.AuthService, db *gorm.DB) *PasswordResetHandler {
	return &PasswordResetHandler{
		passwordResetService: passwordResetService,
		authService:          authService,
		db:                   db,
	}
}

// HandleRequestPasswordReset gère la demande de réinitialisation de mot de passe
func (h *PasswordResetHandler) HandleRequestPasswordReset(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Appel de GetUserIDByEmail depuis le service auth
	userIDResponse, err := h.authService.GetUserIDByEmail(req.Email)
	if err != nil || !userIDResponse.Success {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	resetCode, err := h.passwordResetService.SendResetCode(req.Email, userIDResponse.UserID)
	if err != nil {
		resp := response.SendResetCodeResponse{
			Success: false,
			Message: "Failed to send reset code",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := response.SendResetCodeResponse{
		Success:   true,
		Message:   "Reset code sent successfully",
		ResetCode: resetCode,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// HandleUpdatePassword gère la mise à jour du mot de passe
func (h *PasswordResetHandler) HandleUpdatePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Email       string `json:"email"`
		NewPassword string `json:"new_password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if responseData, err := h.passwordResetService.UpdatePassword(req.Email, req.NewPassword); err != nil {
		// En cas d'erreur dans le service, on renvoie un message d'erreur
		http.Error(w, responseData.Message, http.StatusInternalServerError)
		return
	} else if !responseData.Success {
		// Si l'opération n'est pas réussie mais sans erreur (par exemple, user not found)
		http.Error(w, responseData.Message, http.StatusBadRequest)
		return
	}

	resp := response.UpdatePasswordResponse{
		Success: true,
		Message: "Password updated successfully",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// HandleGetResetCode gère la récupération du code de réinitialisation
func (h *PasswordResetHandler) HandleGetResetCode(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	code, err := h.passwordResetService.GetResetCode(req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := response.GetResetCodeResponse{
		Success:   true,
		Message:   "Reset code retrieved successfully",
		ResetCode: code,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// HandleVerifyResetCode gère la vérification du code de réinitialisation
func (h *PasswordResetHandler) HandleVerifyResetCode(w http.ResponseWriter, r *http.Request) {
	var req request.VerifyResetCodeRequest

	// Décoder le payload JSON dans VerifyResetCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Vérifier l'existence de l'utilisateur par email
	var user models.User
	if err := h.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Vérifier l'entrée de code de réinitialisation par user_id et reset_code
	var resetEntry models.PasswordReset
	if err := h.db.Where("user_id = ? AND reset_code = ?", user.ID, req.ResetCode).First(&resetEntry).Error; err != nil {
		http.Error(w, "Invalid reset code", http.StatusBadRequest)
		return
	}

	// Vérifier l'expiration du code de réinitialisation
	if time.Now().After(resetEntry.Expiration) {
		http.Error(w, "The reset code has expired", http.StatusUnauthorized)
		return
	}

	// Créer et encoder la réponse en utilisant VerifyResetCodeResponse
	resp := response.VerifyResetCodeResponse{
		Success: true,
		Message: "Reset code is valid",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
