package response

// RegisterResponse représente les données renvoyées après un enregistrement réussi.
type RegisterResponse struct {
    UserID      uint   `json:"userId"`
    Message     string `json:"message"`
    Success     bool   `json:"success"`
}

// RegisterUserResponse représente la réponse après un enregistrement réussi.
type RegisterUserResponse struct {
    UserID  uint   `json:"userId"`
    Message string `json:"message"`
    Success bool   `json:"success"`
}

// RegisterBusinessUserResponse représente la réponse après l'enregistrement d'un utilisateur business.
type RegisterBusinessUserResponse struct {
    UserID  uint   `json:"userId"`
    Message string `json:"message"`
    Success bool   `json:"success"`
}

// RegisterSchoolUserResponse représente la réponse après un enregistrement réussi.
type RegisterSchoolUserResponse struct {
    UserID  uint   `json:"userId"`
    Message string `json:"message"`
    Success bool   `json:"success"`
}

// RegisterAssociationUserResponse représente la réponse après un enregistrement réussi.
type RegisterAssociationUserResponse struct {
    UserID  uint   `json:"userId"`
    Message string `json:"message"`
    Success bool   `json:"success"`
}


type RegisterOrganisationUserResponse struct {
    UserID  uint   `json:"userId"`
    Message string `json:"message"`
    Success bool   `json:"success"`
}