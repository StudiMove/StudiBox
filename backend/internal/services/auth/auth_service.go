package auth

import (
	"backend/config"
	"backend/internal/db/models"
	"backend/internal/utils"
	"errors"
	"fmt"
	"log"

	"backend/internal/api/models/auth/request"
	"backend/internal/api/models/auth/response"

	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

// NewAuthService crée une nouvelle instance de AuthService
func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

func (s *AuthService) CheckUserRole(req *request.CheckUserRoleRequest) (*response.CheckUserRoleResponse, error) {
	var user models.User

	// Charger l'utilisateur avec ses rôles
	if err := s.DB.Preload("Roles").First(&user, req.UserID).Error; err != nil {
		return &response.CheckUserRoleResponse{
			HasRole: false,
			Message: "User not found",
		}, err
	}

	// Vérification des rôles
	switch v := req.Roles.(type) {
	case string:
		for _, userRole := range user.Roles {
			if userRole.Name == v {
				isValid, _ := s.isUserValidated(&request.UserValidationRequest{UserID: req.UserID, Role: userRole.Name})
				return &response.CheckUserRoleResponse{HasRole: isValid, Message: "Role found"}, nil
			}
		}
	case []string:
		for _, userRole := range user.Roles {
			for _, requiredRole := range v {
				if userRole.Name == requiredRole {
					isValid, _ := s.isUserValidated(&request.UserValidationRequest{UserID: req.UserID, Role: userRole.Name})
					return &response.CheckUserRoleResponse{HasRole: isValid, Message: "Role found"}, nil
				}
			}
		}
	default:
		return &response.CheckUserRoleResponse{HasRole: false, Message: "Invalid roles format"}, errors.New("invalid type for roles parameter")
	}

	return &response.CheckUserRoleResponse{HasRole: false, Message: "No matching role found"}, nil
}

// / J UTILISE PENDINT ICIC FAUT CREE LE HANDLER
func (s *AuthService) isUserValidated(req *request.UserValidationRequest) (bool, error) {
	switch req.Role {
	case "admin":
		return true, nil
	case "business":
		var businessUser models.BusinessUser
		if err := s.DB.Where("user_id = ?", req.UserID).First(&businessUser).Error; err != nil {
			return false, err
		}
		return !businessUser.IsPending, nil
	case "school":
		var schoolUser models.SchoolUser
		if err := s.DB.Where("user_id = ?", req.UserID).First(&schoolUser).Error; err != nil {
			return false, err
		}
		return !schoolUser.IsPending, nil
	case "association":
		var associationUser models.AssociationUser
		if err := s.DB.Where("user_id = ?", req.UserID).First(&associationUser).Error; err != nil {
			return false, err
		}
		return !associationUser.IsPending, nil
	}
	return false, nil
}

// CheckIfEmailExists vérifie si un email existe déjà dans la base de données.
func (s *AuthService) CheckIfEmailExists(emailReq *request.EmailExistenceRequest) error {
	var existingUser models.User
	err := s.DB.Where("email = ?", emailReq.Email).First(&existingUser).Error
	if err == nil {
		log.Printf("Email %s already exists in the system", emailReq.Email)
		return errors.New("email already taken")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Email %s does not exist, proceeding with registration", emailReq.Email)
		return nil
	}
	log.Printf("Error checking email %s: %v", emailReq.Email, err)
	return err
}

// GetRoleIDByName récupère l'ID d'un rôle par son nom.
func (s *AuthService) GetRoleIDByName(req *request.RoleIDRequest) (*response.RoleIDResponse, error) {
	var role models.Role
	if err := s.DB.Where("name = ?", req.RoleName).First(&role).Error; err != nil {
		return &response.RoleIDResponse{RoleID: 0, Message: "Role not found"}, err
	}
	return &response.RoleIDResponse{RoleID: role.ID, Message: "Role found"}, nil
}

func (s *AuthService) AssignUserRole(req *request.AssignUserRoleRequest) (*response.AssignUserRoleResponse, error) {
	var userRole models.UserRole

	// Vérifie si l'association existe déjà
	if err := s.DB.Where("user_id = ? AND role_id = ?", req.UserID, req.RoleID).First(&userRole).Error; err == nil {
		return &response.AssignUserRoleResponse{Assigned: false, Message: "User already has this role"}, nil
	}

	// Crée une nouvelle entrée
	newUserRole := models.UserRole{
		UserID: req.UserID,
		RoleID: req.RoleID,
	}

	if err := s.DB.Create(&newUserRole).Error; err != nil {
		return &response.AssignUserRoleResponse{Assigned: false, Message: "Failed to assign role"}, err
	}
	return &response.AssignUserRoleResponse{Assigned: true, Message: "Role assigned successfully"}, nil
}

// Login gère la connexion d'un utilisateur et retourne une structure LoginResponse.
func (s *AuthService) Login(loginReq *request.LoginRequest) (*response.LoginResponse, error) {
	var user models.User

	// Vérifie si l'utilisateur existe en recherchant par email
	if err := s.DB.Where("email = ?", loginReq.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err // Retourne une erreur générique si une autre erreur est rencontrée
	}

	// Vérifie le mot de passe
	if err := utils.VerifyPassword(user.Password, loginReq.Password); err != nil {
		return nil, errors.New("invalid password")
	}

	// Génère un JWT pour l'utilisateur avec son ID
	token, err := utils.GenerateJWT(user.ID, config.AppConfig.JwtSecretAccessKey)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}
	// Génère un Refresh Token et le stocke en base de données
	refreshToken, err := utils.GenerateMobileRefreshToken(user.ID, config.AppConfig.JwtSecretRefreshKey)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	// Sauvegarder le refresh token dans la base de données
	user.RefreshToken = refreshToken
	if err := s.DB.Save(&user).Error; err != nil {
		return nil, errors.New("failed to save refresh token")
	}

	// Crée la réponse de connexion avec les détails de l’utilisateur
	loginResp := &response.LoginResponse{
		Token:           token,
		ProfileImage:    user.ProfileImage,
		IsAuthenticated: true,
	}

	return loginResp, nil
}

func (s *AuthService) RegisterUser(registerReq *request.RegisterUserRequest) (*response.RegisterUserResponse, error) {
	// Vérifie si l'email est déjà pris
	emailCheckReq := &request.EmailExistenceRequest{Email: registerReq.Email}
	if err := s.CheckIfEmailExists(emailCheckReq); err != nil {
		return nil, err
	}

	// Hasher le mot de passe
	hashedPassword, err := utils.HashPassword(registerReq.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
		return nil, errors.New("failed to hash password")
	}

	// Crée et enregistre l'utilisateur
	user := models.User{
		Email:        registerReq.Email,
		Password:     hashedPassword,
		FirstName:    registerReq.FirstName,
		LastName:     registerReq.LastName,
		Phone:        registerReq.Phone,
		ProfileImage: registerReq.ProfileImage,
	}
	if err := s.DB.Create(&user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	return &response.RegisterUserResponse{
		UserID:  user.ID,
		Message: "User successfully registered",
		Success: true,
	}, nil
}

// RegisterBusinessUser gère l'inscription d'un utilisateur de type business.
func (s *AuthService) RegisterBusinessUser(registerReq *request.RegisterBusinessUserRequest) (*response.RegisterBusinessUserResponse, error) {
	// Vérifie si l'email est déjà pris
	emailCheckReq := &request.EmailExistenceRequest{Email: registerReq.Email}
	if err := s.CheckIfEmailExists(emailCheckReq); err != nil {
		return nil, err
	}

	// Hasher le mot de passe
	hashedPassword, err := utils.HashPassword(registerReq.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Crée et enregistre l'utilisateur
	businessUser := models.BusinessUser{
		User: models.User{
			Email:    registerReq.Email,
			Password: hashedPassword,
			Phone:    registerReq.Phone,
		},
		CompanyName: registerReq.OrganisationName,
		Address:     registerReq.Address,
		Postcode:    registerReq.PostalCode,
		City:        registerReq.City,
		Country:     registerReq.Country,
		Description: registerReq.Description,
	}
	if err := s.DB.Create(&businessUser.User).Error; err != nil {
		return nil, errors.New("failed to create user")
	}
	businessUser.UserID = businessUser.User.ID
	if err := s.DB.Create(&businessUser).Error; err != nil {
		return nil, errors.New("failed to save business user")
	}

	// Récupération du rôle
	roleIDReq := &request.RoleIDRequest{RoleName: "business"}
	roleIDResp, err := s.GetRoleIDByName(roleIDReq)
	if err != nil {
		return nil, err
	}

	// Assignation du rôle
	assignReq := &request.AssignUserRoleRequest{UserID: businessUser.User.ID, RoleID: roleIDResp.RoleID}
	if _, err := s.AssignUserRole(assignReq); err != nil {
		return nil, err
	}

	return &response.RegisterBusinessUserResponse{
		UserID:  businessUser.User.ID,
		Message: "Business user successfully registered",
		Success: true,
	}, nil
}

// RegisterAssociationUser gère l'inscription d'un utilisateur de type association.
func (s *AuthService) RegisterAssociationUser(registerReq *request.RegisterAssociationUserRequest) (*response.RegisterAssociationUserResponse, error) {
	// Vérifie si l'email est déjà pris
	emailCheckReq := &request.EmailExistenceRequest{Email: registerReq.Email}
	if err := s.CheckIfEmailExists(emailCheckReq); err != nil {
		return nil, err
	}

	// Hasher le mot de passe
	hashedPassword, err := utils.HashPassword(registerReq.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Créer le modèle AssociationUser
	associationUser := models.AssociationUser{
		User: models.User{
			Email:    registerReq.Email,
			Password: hashedPassword,
			Phone:    registerReq.Phone,
		},
		AssociationName: registerReq.AssociationName,
		Address:         registerReq.Address,
		Postcode:        registerReq.PostalCode,
		City:            registerReq.City,
		Country:         registerReq.Country,
		Description:     registerReq.Description,
	}

	// Enregistrer l'utilisateur et l'information spécifique AssociationUser
	if err := s.DB.Create(&associationUser.User).Error; err != nil {
		return nil, errors.New("failed to create user")
	}
	associationUser.UserID = associationUser.User.ID
	if err := s.DB.Create(&associationUser).Error; err != nil {
		return nil, errors.New("failed to create association user")
	}

	// Attribuer le rôle "association" à l'utilisateur
	roleIDReq := &request.RoleIDRequest{RoleName: "association"}
	roleIDResp, err := s.GetRoleIDByName(roleIDReq)
	if err != nil {
		return nil, err
	}

	assignRoleReq := &request.AssignUserRoleRequest{
		UserID: associationUser.User.ID,
		RoleID: roleIDResp.RoleID,
	}
	if _, err := s.AssignUserRole(assignRoleReq); err != nil {
		return nil, err
	}

	return &response.RegisterAssociationUserResponse{
		UserID:  associationUser.User.ID,
		Message: "Association user successfully registered",
		Success: true,
	}, nil
}

// RegisterSchoolUser gère l'inscription d'un utilisateur de type school.
func (s *AuthService) RegisterSchoolUser(registerReq *request.RegisterSchoolUserRequest) (*response.RegisterSchoolUserResponse, error) {
	// Vérifie si l'email est déjà pris
	emailCheckReq := &request.EmailExistenceRequest{Email: registerReq.Email}
	if err := s.CheckIfEmailExists(emailCheckReq); err != nil {
		return nil, err
	}

	// Hasher le mot de passe
	hashedPassword, err := utils.HashPassword(registerReq.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Créer le modèle SchoolUser
	schoolUser := models.SchoolUser{
		User: models.User{
			Email:    registerReq.Email,
			Password: hashedPassword,
			Phone:    registerReq.Phone,
		},
		SchoolName:  registerReq.SchoolName,
		Address:     registerReq.Address,
		Postcode:    registerReq.PostalCode,
		City:        registerReq.City,
		Country:     registerReq.Country,
		Description: registerReq.Description,
	}

	// Enregistrer l'utilisateur et l'information spécifique SchoolUser
	if err := s.DB.Create(&schoolUser.User).Error; err != nil {
		return nil, errors.New("failed to create user")
	}
	schoolUser.UserID = schoolUser.User.ID
	if err := s.DB.Create(&schoolUser).Error; err != nil {
		return nil, errors.New("failed to create school user")
	}

	// Attribuer le rôle "school" à l'utilisateur
	roleIDReq := &request.RoleIDRequest{RoleName: "school"}
	roleIDResp, err := s.GetRoleIDByName(roleIDReq)
	if err != nil {
		return nil, err
	}

	assignRoleReq := &request.AssignUserRoleRequest{
		UserID: schoolUser.User.ID,
		RoleID: roleIDResp.RoleID,
	}
	if _, err := s.AssignUserRole(assignRoleReq); err != nil {
		return nil, err
	}

	return &response.RegisterSchoolUserResponse{
		UserID:  schoolUser.User.ID,
		Message: "School user successfully registered",
		Success: true,
	}, nil
}

// RegisterOrganisationUser gère l'inscription d'un utilisateur en fonction du type d'organisation.
func (s *AuthService) RegisterOrganisationUser(registerReq *request.RegisterOrganisationUserRequest) (*response.RegisterOrganisationUserResponse, error) {
	// Vérifie si l'email est déjà pris
	emailCheckReq := &request.EmailExistenceRequest{Email: registerReq.Email}
	if err := s.CheckIfEmailExists(emailCheckReq); err != nil {
		return nil, err
	}

	// Hasher le mot de passe
	hashedPassword, err := utils.HashPassword(registerReq.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	var userID uint
	// Détecte le type d'organisation et crée le modèle approprié
	switch registerReq.OrganisationType {
	case "business":
		businessUser := &models.BusinessUser{
			User: models.User{
				Email:    registerReq.Email,
				Password: hashedPassword,
				Phone:    registerReq.Phone,
			},
			CompanyName: registerReq.OrganisationName,
			Address:     registerReq.Address,
			Postcode:    registerReq.PostalCode,
			City:        registerReq.City,
			Country:     registerReq.Country,
			Description: registerReq.Description,
		}
		if err := s.DB.Create(&businessUser.User).Error; err != nil {
			return nil, errors.New("failed to create business user")
		}
		businessUser.UserID = businessUser.User.ID
		if err := s.DB.Create(businessUser).Error; err != nil {
			return nil, errors.New("failed to save business user")
		}
		userID = businessUser.User.ID

	case "school":
		schoolUser := &models.SchoolUser{
			User: models.User{
				Email:    registerReq.Email,
				Password: hashedPassword,
				Phone:    registerReq.Phone,
			},
			SchoolName:  registerReq.OrganisationName,
			Address:     registerReq.Address,
			Postcode:    registerReq.PostalCode,
			City:        registerReq.City,
			Country:     registerReq.Country,
			Description: registerReq.Description,
		}
		if err := s.DB.Create(&schoolUser.User).Error; err != nil {
			return nil, errors.New("failed to create school user")
		}
		schoolUser.UserID = schoolUser.User.ID
		if err := s.DB.Create(schoolUser).Error; err != nil {
			return nil, errors.New("failed to save school user")
		}
		userID = schoolUser.User.ID

	case "association":
		associationUser := &models.AssociationUser{
			User: models.User{
				Email:    registerReq.Email,
				Password: hashedPassword,
				Phone:    registerReq.Phone,
			},
			AssociationName: registerReq.OrganisationName,
			Address:         registerReq.Address,
			Postcode:        registerReq.PostalCode,
			City:            registerReq.City,
			Country:         registerReq.Country,
			Description:     registerReq.Description,
		}
		if err := s.DB.Create(&associationUser.User).Error; err != nil {
			return nil, errors.New("failed to create association user")
		}
		associationUser.UserID = associationUser.User.ID
		if err := s.DB.Create(associationUser).Error; err != nil {
			return nil, errors.New("failed to save association user")
		}
		userID = associationUser.User.ID

	default:
		return nil, errors.New("invalid organisation type")
	}

	return &response.RegisterOrganisationUserResponse{
		UserID:  userID, // Utilise l'ID généré par la création en BD
		Message: fmt.Sprintf("%s user successfully registered", registerReq.OrganisationType),
		Success: true,
	}, nil
}

// GetUserIDByEmail récupère l'ID de l'utilisateur à partir de l'email
func (s *AuthService) GetUserIDByEmail(email string) (*response.UserIDResponse, error) {
	var user models.User

	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &response.UserIDResponse{
				Success: false,
				Message: "User not found",
			}, nil
		}
		return nil, err
	}

	return &response.UserIDResponse{
		Success: true,
		UserID:  user.ID,
		Message: "User ID retrieved successfully",
	}, nil
}

// func (s *AuthService) RegisterNormalUser(registerReq *request.RegisterNormalUserRequest) (*response.RegisterResponse, error) {
// 	// Vérifie si l'email est déjà pris
// 	emailCheckReq := &request.EmailExistenceRequest{Email: registerReq.Email}
// 	if err := s.CheckIfEmailExists(emailCheckReq); err != nil {
// 		return nil, err
// 	}

// 	// Vérifie le code de parrainage s'il est fourni
// 	var parrainID uint
// 	if registerReq.ParrainageCode != "" {
// 		var parrain models.User
// 		if err := s.DB.Where("parrain_code = ?", registerReq.ParrainageCode).First(&parrain).Error; err != nil {
// 			return nil, errors.New("invalid parrainage code")
// 		}
// 		parrainID = parrain.ID
// 	}

// 	// Hasher le mot de passe
// 	hashedPassword, err := utils.HashPassword(registerReq.Password)
// 	if err != nil {
// 		return nil, errors.New("failed to hash password")
// 	}

// 	// Générer un code de parrain unique
// 	parrainCode := utils.GenerateParrainCode()

// 	// Crée l'utilisateur sans encore l'insérer
// 	user := models.User{
// 		Email:          registerReq.Email,
// 		Password:       hashedPassword,
// 		FirstName:      registerReq.FirstName,
// 		LastName:       registerReq.LastName,
// 		ParrainageCode: registerReq.ParrainageCode,
// 		ParrainCode:    parrainCode,
// 	}

// 	// Enregistrer l'utilisateur dans la base de données
// 	if err := s.DB.Create(&user).Error; err != nil {
// 		return nil, errors.New("failed to create user")
// 	}

// 	// Si le code de parrainage est valide, créer une relation de parrainage
// 	if parrainID != 0 {
// 		referral := models.Referral{
// 			ParrainID: parrainID,
// 			FilleulID: user.ID,
// 		}
// 		if err := s.DB.Create(&referral).Error; err != nil {
// 			// Supprimer l'utilisateur si l'insertion de la relation échoue
// 			s.DB.Delete(&user)
// 			return nil, errors.New("failed to create referral")
// 		}
// 	}

// 	// Récupérer l'ID du rôle "user"
// 	roleIDReq := &request.RoleIDRequest{RoleName: "user"}
// 	roleIDResp, err := s.GetRoleIDByName(roleIDReq)
// 	if err != nil {
// 		// Supprimer l'utilisateur si l'attribution du rôle échoue
// 		s.DB.Delete(&user)
// 		return nil, err
// 	}

// 	// Assigner le rôle "user" à l'utilisateur
// 	assignRoleReq := &request.AssignUserRoleRequest{
// 		UserID: user.ID,
// 		RoleID: roleIDResp.RoleID,
// 	}
// 	if _, err := s.AssignUserRole(assignRoleReq); err != nil {
// 		// Supprimer l'utilisateur si l'attribution du rôle échoue
// 		s.DB.Delete(&user)
// 		return nil, err
// 	}

//		// Retourner une réponse réussie
//		return &response.RegisterResponse{
//			UserID:  user.ID,
//			Message: "Normal user successfully registered",
//			Success: true,
//		}, nil
//	}

func (s *AuthService) RegisterNormalUser(registerReq *request.RegisterNormalUserRequest) (*response.RegisterUserResponse, error) {
	// Vérifie si l'email est déjà pris
	emailCheckReq := &request.EmailExistenceRequest{Email: registerReq.Email}
	if err := s.CheckIfEmailExists(emailCheckReq); err != nil {
		return nil, err
	}

	// Vérifie le code de parrainage s'il est fourni
	var parrainID uint
	if registerReq.ParrainageCode != "" {
		var parrain models.User
		if err := s.DB.Where("parrain_code = ?", registerReq.ParrainageCode).First(&parrain).Error; err != nil {
			return nil, errors.New("invalid parrainage code")
		}
		parrainID = parrain.ID
	}

	// Hasher le mot de passe
	hashedPassword, err := utils.HashPassword(registerReq.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Générer un code de parrain unique
	parrainCode := utils.GenerateParrainCode()

	// Crée l'utilisateur sans encore l'insérer
	user := models.User{
		Email:          registerReq.Email,
		Password:       hashedPassword,
		FirstName:      registerReq.FirstName,
		LastName:       registerReq.LastName,
		ParrainageCode: registerReq.ParrainageCode,
		ParrainCode:    parrainCode,
	}

	// Enregistrer l'utilisateur dans la base de données
	if err := s.DB.Create(&user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	// Si le code de parrainage est valide, créer une relation de parrainage
	if parrainID != 0 {
		referral := models.Referral{
			ParrainID: parrainID,
			FilleulID: user.ID,
		}
		if err := s.DB.Create(&referral).Error; err != nil {
			// Supprimer l'utilisateur si l'insertion de la relation échoue
			s.DB.Delete(&user)
			return nil, errors.New("failed to create referral")
		}
	}

	// Récupérer l'ID du rôle "user"
	roleIDReq := &request.RoleIDRequest{RoleName: "user"}
	roleIDResp, err := s.GetRoleIDByName(roleIDReq)
	if err != nil {
		// Supprimer l'utilisateur si l'attribution du rôle échoue
		s.DB.Delete(&user)
		return nil, err
	}

	// Assigner le rôle "user" à l'utilisateur
	assignRoleReq := &request.AssignUserRoleRequest{
		UserID: user.ID,
		RoleID: roleIDResp.RoleID,
	}
	if _, err := s.AssignUserRole(assignRoleReq); err != nil {
		// Supprimer l'utilisateur si l'attribution du rôle échoue
		s.DB.Delete(&user)
		return nil, err
	}

	// Générer un token JWT
	token, err := utils.GenerateJWT(user.ID, config.AppConfig.JwtSecretAccessKey)
	if err != nil {
		return nil, errors.New("failed to generate JWT")
	}

	// Générer un refresh token et le stocker dans la base de données
	refreshToken, err := utils.GenerateMobileRefreshToken(user.ID, config.AppConfig.JwtSecretRefreshKey)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}
	user.RefreshToken = refreshToken
	if err := s.DB.Save(&user).Error; err != nil {
		return nil, errors.New("failed to save refresh token")
	}

	// Retourner une réponse réussie avec le token
	return &response.RegisterUserResponse{
		UserID:  user.ID,
		Message: "User successfully registered",
		Success: true,
		Token:   token,
	}, nil
}

func (s *AuthService) AutoAuthenticate(token string) (*models.User, error) {
	// Extraire l'ID utilisateur du token JWT
	userID, err := utils.ExtractUserIDFromToken(token, config.AppConfig.JwtSecretAccessKey)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	// Vérifier si l'utilisateur existe dans la base de données
	var user models.User
	if err := s.DB.First(&user, userID).Error; err != nil {
		return nil, errors.New("user not found")
	}

	// Retourner l'utilisateur si trouvé
	return &user, nil
}

// RefreshAccessToken génère un nouveau access token pour un utilisateur à partir de son ID
func (s *AuthService) RefreshAccessToken(userID uint) (string, error) {
	// Générer un nouveau access token avec une expiration prolongée
	newAccessToken, err := utils.GenerateJWT(userID, config.AppConfig.JwtSecretAccessKey)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}
