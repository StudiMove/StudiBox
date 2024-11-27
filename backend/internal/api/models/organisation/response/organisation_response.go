package response

// OrganisationResponse représente les données d'une organisation.
type OrganisationResponse struct {
	UserID       uint   `json:"user_id"`
	Name         string `json:"name"`
	IsValidated  bool   `json:"is_validated"`
	IsActivated  bool   `json:"is_activated"`
	IsPending    bool   `json:"is_pending"`
	RoleName     string `json:"role_name"`
	ProfileImage string `json:"profile_image"`
	Status       string `json:"status"`
}

// OrganisationListResponse représente la liste des organisations.
type OrganisationListResponse struct {
    Organisations []OrganisationResponse `json:"organisations"`
    Success       bool                   `json:"success"`
    Message       string                 `json:"message,omitempty"`
}
