package response

type UserProfileResponse struct {
	UserID       uint                   `json:"user_id"`
	Email        string                 `json:"email"`
	Phone        string                 `json:"phone"`
	ProfileImage string                 `json:"profile_image"`
	RoleNames    []string               `json:"roles"`
	Organisation interface{}            `json:"organisation"`
	
}
