package response

type ActionUpdateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ActionRemoveResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
