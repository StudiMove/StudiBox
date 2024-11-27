package response

// RoleIDResponse représente la réponse pour récupérer l'ID d'un rôle
type RoleIDResponse struct {
	RoleID  uint   `json:"roleId"`
	Message string `json:"message"`
}
