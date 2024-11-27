package response

// CheckUserRoleResponse représente la réponse pour vérifier les rôles de l'utilisateur
type CheckUserRoleResponse struct {
	HasRole bool   `json:"hasRole"`
	Message string `json:"message"`
}