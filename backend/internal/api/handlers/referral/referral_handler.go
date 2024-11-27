package referral

import (
	"encoding/json"
	"net/http"

	"backend/internal/services/referral"
)

// ReferralHandler gère les requêtes HTTP liées aux parrainages
type ReferralHandler struct {
	referralService *referral.ReferralService
}

// NewReferralHandler crée une nouvelle instance de ReferralHandler
func NewReferralHandler(referralService *referral.ReferralService) *ReferralHandler {
	return &ReferralHandler{referralService: referralService}
}

// HandleGetFilleulsByParrain gère la récupération des filleuls d'un parrain
func (h *ReferralHandler) HandleGetFilleulsByParrain(w http.ResponseWriter, r *http.Request) {
	// Décodez le corps de la requête pour obtenir le parrainId
	var request struct {
		ParrainID uint `json:"parrainId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Vérifiez si l'ID est valide
	if request.ParrainID == 0 {
		http.Error(w, "Missing or invalid parrainId", http.StatusBadRequest)
		return
	}

	// Appeler le service pour récupérer les filleuls
	filleuls, err := h.referralService.GetFilleulsByParrain(request.ParrainID)
	if err != nil {
		http.Error(w, "Failed to retrieve filleuls", http.StatusInternalServerError)
		return
	}

	// Répondre avec les données des filleuls
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filleuls)
}

// HandleCountFilleuls gère le comptage des filleuls d'un parrain
func (h *ReferralHandler) HandleCountFilleuls(w http.ResponseWriter, r *http.Request) {
	// Décodez le corps de la requête pour obtenir le parrainId
	var request struct {
		ParrainID uint `json:"parrainId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Vérifiez si l'ID est valide
	if request.ParrainID == 0 {
		http.Error(w, "Missing or invalid parrainId", http.StatusBadRequest)
		return
	}

	// Appeler le service pour compter les filleuls
	count, err := h.referralService.CountFilleuls(request.ParrainID)
	if err != nil {
		http.Error(w, "Failed to count filleuls", http.StatusInternalServerError)
		return
	}

	// Répondre avec le nombre de filleuls
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int64{"count": count})
}

// HandleGetFilleulIDsByParrain gère la récupération des IDs des filleuls d'un parrain
func (h *ReferralHandler) HandleGetFilleulIDsByParrain(w http.ResponseWriter, r *http.Request) {
	// Décoder le corps de la requête pour obtenir le parrainId
	var request struct {
		ParrainID uint `json:"parrainId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Vérifier si l'ID est valide
	if request.ParrainID == 0 {
		http.Error(w, "Missing or invalid parrainId", http.StatusBadRequest)
		return
	}

	// Appeler le service pour récupérer les IDs des filleuls
	filleulIDs, err := h.referralService.GetFilleulIDsByParrain(request.ParrainID)
	if err != nil {
		http.Error(w, "Failed to retrieve filleul IDs", http.StatusInternalServerError)
		return
	}

	// Répondre avec les IDs des filleuls
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filleulIDs)
}
