package organisation

import (
	"backend/internal/api/models/organisation/response"
	"backend/internal/services/organisation"
	"encoding/json"
	"net/http"
)

// GetAllOrganisationsHandler gère les requêtes pour récupérer les organisations
type GetAllOrganisationsHandler struct {
	organisationService *organisation.OrganisationService // Service pour interagir avec les données d'organisation
}

// NewGetAllOrganisationsHandler crée une nouvelle instance de GetAllOrganisationsHandler
func NewGetAllOrganisationsHandler(orgService *organisation.OrganisationService) *GetAllOrganisationsHandler {
	return &GetAllOrganisationsHandler{organisationService: orgService}
}

// HandleGetAllOrganisations gère la requête pour récupérer toutes les organisations
func (h *GetAllOrganisationsHandler) HandleGetAllOrganisations(w http.ResponseWriter, r *http.Request) {
	// Appel au service pour obtenir toutes les organisations
	orgResp, err := h.organisationService.GetAllOrganisations()
	if err != nil {
		http.Error(w, "Failed to fetch organisations", http.StatusInternalServerError) // Gérer l'erreur en cas d'échec
		return
	}

	// Préparer la réponse pour la liste des organisations
	response := &response.OrganisationListResponse{
		Success:       true,
		Message:       "Organisations retrieved successfully",
		Organisations: orgResp.Organisations,
	}

	// Définir le type de contenu de la réponse à JSON et envoyer la réponse
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response) // Encoder et envoyer la réponse
}

// HandleGetActiveOrganisations gère la requête pour récupérer les organisations actives
func (h *GetAllOrganisationsHandler) HandleGetActiveOrganisations(w http.ResponseWriter, r *http.Request) {
	// Appel au service pour obtenir les organisations actives
	orgResp, err := h.organisationService.GetActiveOrganisations()
	if err != nil {
		http.Error(w, "Failed to fetch active organisations", http.StatusInternalServerError) // Gérer l'erreur en cas d'échec
		return
	}

	// Préparer la réponse pour la liste des organisations actives
	response := &response.OrganisationListResponse{
		Success:       true,
		Message:       "Active organisations retrieved successfully",
		Organisations: orgResp.Organisations,
	}

	// Définir le type de contenu de la réponse à JSON et envoyer la réponse
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response) // Encoder et envoyer la réponse
}

// HandleGetInactiveOrganisations gère la requête pour récupérer les organisations inactives
func (h *GetAllOrganisationsHandler) HandleGetInactiveOrganisations(w http.ResponseWriter, r *http.Request) {
	// Appel au service pour obtenir les organisations inactives
	organisations, err := h.organisationService.GetInactiveOrganisations()
	if err != nil {
		http.Error(w, "Failed to fetch inactive organisations", http.StatusInternalServerError) // Gérer l'erreur en cas d'échec
		return
	}

	// Préparer la réponse pour la liste des organisations inactives
	response := &response.OrganisationListResponse{
		Success:       true,
		Message:       "Inactive organisations retrieved successfully",
		Organisations: organisations.Organisations,
	}

	// Définir le type de contenu de la réponse à JSON et envoyer la réponse
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response) // Encoder et envoyer la réponse
}

// HandleGetPendingOrganisations gère la requête pour récupérer les organisations en attente
func (h *GetAllOrganisationsHandler) HandleGetPendingOrganisations(w http.ResponseWriter, r *http.Request) {
	// Appel au service pour obtenir les organisations en attente
	organisations, err := h.organisationService.GetAllPendingOrganisations()
	if err != nil {
		http.Error(w, "Failed to fetch pending organisations", http.StatusInternalServerError) // Gérer l'erreur en cas d'échec
		return
	}

	// Préparer la réponse pour la liste des organisations en attente
	response := &response.OrganisationListResponse{
		Success:       true,
		Message:       "Pending organisations retrieved successfully",
		Organisations: organisations.Organisations,
	}

	// Définir le type de contenu de la réponse à JSON et envoyer la réponse
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response) // Encoder et envoyer la réponse
}

// HandleGetSuspendedOrganisations gère la requête pour récupérer les organisations suspendues
func (h *GetAllOrganisationsHandler) HandleGetSuspendedOrganisations(w http.ResponseWriter, r *http.Request) {
	// Appel au service pour obtenir les organisations suspendues
	organisations, err := h.organisationService.GetSuspendedOrganisations()
	if err != nil {
		http.Error(w, "Failed to fetch suspended organisations", http.StatusInternalServerError) // Gérer l'erreur en cas d'échec
		return
	}

	// Préparer la réponse pour la liste des organisations suspendues
	response := &response.OrganisationListResponse{
		Success:       true,
		Message:       "Suspended organisations retrieved successfully",
		Organisations: organisations.Organisations,
	}

	// Définir le type de contenu de la réponse à JSON et envoyer la réponse
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response) // Encoder et envoyer la réponse
}
