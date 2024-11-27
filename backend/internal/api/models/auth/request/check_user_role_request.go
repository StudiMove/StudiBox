package request

// CheckUserRoleRequest représente la requête pour vérifier les rôles de l'utilisateur
type CheckUserRoleRequest struct {
	UserID uint        `json:"userId"`
	Roles  interface{} `json:"roles"` // Peut être une string ou un tableau de strings
}
