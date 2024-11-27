package request

// AssignUserRoleRequest représente la requête pour assigner un rôle à un utilisateur
type AssignUserRoleRequest struct {
	UserID uint `json:"userId"`
	RoleID uint `json:"roleId"`
}
