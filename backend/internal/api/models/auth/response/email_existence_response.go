package response

// EmailExistenceResponse représente la réponse pour vérifier si un email existe
type EmailExistenceResponse struct {
	Exists  bool   `json:"exists"`
	Message string `json:"message"`
}
