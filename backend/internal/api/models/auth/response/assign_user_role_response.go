package response

// AssignUserRoleResponse représente la réponse pour l'assignation d'un rôle à un utilisateur
type AssignUserRoleResponse struct {
	Assigned bool   `json:"assigned"`
	Message  string `json:"message"`
}
