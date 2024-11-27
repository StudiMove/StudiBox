package request

// RoleIDRequest représente la requête pour récupérer un rôle par nom
type RoleIDRequest struct {
	RoleName string `json:"roleName"`
}
