package response

// APIResponse représente une réponse d'API standardisée
type APIResponse struct {
	Status  string      `json:"status"`  // success or error
	Message string      `json:"message"` // message court
	Data    interface{} `json:"data"`    // contenu
}

// SuccessResponse crée une réponse de succès avec des données
func SuccessResponse(message string, data interface{}) *APIResponse {
	return &APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

// ErrorResponse crée une réponse d'erreur avec un message d'erreur
func ErrorResponse(message string, err error) *APIResponse {
	if err != nil {
		return &APIResponse{
			Status:  "error",
			Message: message,
			Data:    map[string]string{"error": err.Error()},
		}
	}
	return &APIResponse{
		Status:  "error",
		Message: message,
		Data:    nil,
	}
}
