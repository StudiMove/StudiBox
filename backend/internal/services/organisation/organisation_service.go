package organisation

import (
	"backend/internal/api/models/organisation/response"
	"errors"
	"strings"

	"gorm.io/gorm"
)

// OrganisationService représente le service pour gérer les organisations.
type OrganisationService struct {
	db *gorm.DB
}

// NewOrganisationService crée une nouvelle instance de OrganisationService.
func NewOrganisationService(db *gorm.DB) *OrganisationService {
	return &OrganisationService{db: db}
}

// GetAllOrganisations récupère les données de BusinessUser, SchoolUser et AssociationUser

func (s *OrganisationService) GetAllOrganisations() (*response.OrganisationListResponse, error) {
	var organisations []response.OrganisationResponse

	// Récupération des Business Users
	err := s.db.Table("business_users").
		Select("business_users.user_id, business_users.company_name AS name, business_users.is_validated, business_users.is_activated, business_users.is_pending, business_users.status, users.profile_image").
		Joins("inner join users on users.id = business_users.user_id").
		Scan(&organisations).Error
	if err != nil {
		return nil, errors.New("failed to fetch organisations from business_users")
	}

	// Récupération des autres utilisateurs (School, Association)
	var schoolUsers []response.OrganisationResponse
	err = s.db.Table("school_users").
		Select("school_users.user_id, school_users.school_name AS name, school_users.is_validated, school_users.is_activated, school_users.is_pending, school_users.status, users.profile_image").
		Joins("inner join users on users.id = school_users.user_id").
		Scan(&schoolUsers).Error
	if err != nil {
		return nil, errors.New("failed to fetch organisations from school_users")
	}

	var associationUsers []response.OrganisationResponse
	err = s.db.Table("association_users").
		Select("association_users.user_id, association_users.association_name AS name, association_users.is_validated, association_users.is_activated, association_users.is_pending, association_users.status, users.profile_image").
		Joins("inner join users on users.id = association_users.user_id").
		Scan(&associationUsers).Error
	if err != nil {
		return nil, errors.New("failed to fetch organisations from association_users")
	}

	// Concaténer les résultats
	organisations = append(organisations, schoolUsers...)
	organisations = append(organisations, associationUsers...)

	// Récupérer les rôles pour chaque organisation
	for i, organisation := range organisations {
		var roles []string
		s.db.Table("roles").
			Joins("inner join user_roles on user_roles.role_id = roles.id").
			Where("user_roles.user_id = ?", organisation.UserID).
			Pluck("name", &roles)
		organisations[i].RoleName = strings.Join(roles, ", ")
	}

	return &response.OrganisationListResponse{Organisations: organisations}, nil
}

// GetActiveOrganisations récupère les organisations actives.
func (s *OrganisationService) GetActiveOrganisations() (*response.OrganisationListResponse, error) {
	var activeOrganisations []response.OrganisationResponse

	// Récupérer les Business Users actifs
	err := s.db.Table("business_users").
		Select("business_users.user_id, business_users.company_name AS name, business_users.is_validated, business_users.is_activated, business_users.is_pending, business_users.status, users.profile_image").
		Joins("inner join users on users.id = business_users.user_id").
		Where("business_users.is_validated = ? AND business_users.is_activated = ? AND business_users.is_pending = ?", true, true, false).
		Scan(&activeOrganisations).Error

	if err != nil {
		return nil, errors.New("failed to fetch active organisations from business_users")
	}

	// Récupérer les School Users actifs
	var activeSchoolUsers []response.OrganisationResponse
	err = s.db.Table("school_users").
		Select("school_users.user_id, school_users.school_name AS name, school_users.is_validated, school_users.is_activated, school_users.is_pending, school_users.status, users.profile_image").
		Joins("inner join users on users.id = school_users.user_id").
		Where("school_users.is_validated = ? AND school_users.is_activated = ? AND school_users.is_pending = ?", true, true, false).
		Scan(&activeSchoolUsers).Error

	if err != nil {
		return nil, errors.New("failed to fetch active organisations from school_users")
	}

	// Récupérer les Association Users actifs
	var activeAssociationUsers []response.OrganisationResponse
	err = s.db.Table("association_users").
		Select("association_users.user_id, association_users.association_name AS name, association_users.is_validated, association_users.is_activated, association_users.is_pending, association_users.status, users.profile_image").
		Joins("inner join users on users.id = association_users.user_id").
		Where("association_users.is_validated = ? AND association_users.is_activated = ? AND association_users.is_pending = ?", true, true, false).
		Scan(&activeAssociationUsers).Error

	if err != nil {
		return nil, errors.New("failed to fetch active organisations from association_users")
	}

	// Concaténer les résultats
	activeOrganisations = append(activeOrganisations, activeSchoolUsers...)
	activeOrganisations = append(activeOrganisations, activeAssociationUsers...)

	// Récupérer les rôles pour chaque organisation active
	for i, organisation := range activeOrganisations {
		var roles []string
		s.db.Table("roles").
			Joins("inner join user_roles on user_roles.role_id = roles.id").
			Where("user_roles.user_id = ?", organisation.UserID).
			Pluck("name", &roles)

		// Concaténer les rôles
		activeOrganisations[i].RoleName = strings.Join(roles, ", ")
	}

	// Retourner la réponse encapsulée dans OrganisationListResponse
	return &response.OrganisationListResponse{
		Organisations: activeOrganisations,
		Success:       true,
		Message:       "Active organisations retrieved successfully",
	}, nil
}

// GetInactiveOrganisations récupère les organisations inactives (status == isValidatedFalse)
func (s *OrganisationService) GetInactiveOrganisations() (*response.OrganisationListResponse, error) {
	var inactiveOrganisations []response.OrganisationResponse

	// Récupérer les Business Users inactifs
	err := s.db.Table("business_users").
		Select("business_users.user_id, business_users.company_name AS name, business_users.is_validated, business_users.is_activated, business_users.is_pending, business_users.status, users.profile_image").
		Joins("inner join users on users.id = business_users.user_id").
		Where("business_users.is_validated = ? AND business_users.is_activated = ? AND business_users.is_pending = ?", false, true, false).
		Scan(&inactiveOrganisations).Error

	if err != nil {
		return nil, errors.New("failed to fetch inactive organisations from business_users")
	}

	// Récupérer les School Users inactifs
	var inactiveSchoolUsers []response.OrganisationResponse
	err = s.db.Table("school_users").
		Select("school_users.user_id, school_users.school_name AS name, school_users.is_validated, school_users.is_activated, school_users.is_pending, school_users.status, users.profile_image").
		Joins("inner join users on users.id = school_users.user_id").
		Where("school_users.is_validated = ? AND school_users.is_activated = ? AND school_users.is_pending = ?", false, true, false).
		Scan(&inactiveSchoolUsers).Error

	if err != nil {
		return nil, errors.New("failed to fetch inactive organisations from school_users")
	}

	// Récupérer les Association Users inactifs
	var inactiveAssociationUsers []response.OrganisationResponse
	err = s.db.Table("association_users").
		Select("association_users.user_id, association_users.association_name AS name, association_users.is_validated, association_users.is_activated, association_users.is_pending, association_users.status, users.profile_image").
		Joins("inner join users on users.id = association_users.user_id").
		Where("association_users.is_validated = ? AND association_users.is_activated = ? AND association_users.is_pending = ?", false, true, false).
		Scan(&inactiveAssociationUsers).Error

	if err != nil {
		return nil, errors.New("failed to fetch inactive organisations from association_users")
	}

	// Concaténer les résultats
	inactiveOrganisations = append(inactiveOrganisations, inactiveSchoolUsers...)
	inactiveOrganisations = append(inactiveOrganisations, inactiveAssociationUsers...)

	// Récupérer les rôles pour chaque organisation inactive
	for i, organisation := range inactiveOrganisations {
		var roles []string
		s.db.Table("roles").
			Joins("inner join user_roles on user_roles.role_id = roles.id").
			Where("user_roles.user_id = ?", organisation.UserID).
			Pluck("name", &roles)

		// Concaténer les rôles
		inactiveOrganisations[i].RoleName = strings.Join(roles, ", ")
	}

	// Retourner la réponse encapsulée
	return &response.OrganisationListResponse{
		Organisations: inactiveOrganisations,
		Success:       true,
		Message:       "Inactive organisations retrieved successfully",
	}, nil
}

// GetAllPendingOrganisations récupère les organisations en attente de validation (status == isPending)
func (s *OrganisationService) GetAllPendingOrganisations() (*response.OrganisationListResponse, error) {
	var pendingOrganisations []response.OrganisationResponse

	// Récupérer les Business Users en attente
	err := s.db.Table("business_users").
		Select("business_users.user_id, business_users.company_name AS name, business_users.is_validated, business_users.is_activated, business_users.is_pending, business_users.status, users.profile_image").
		Joins("inner join users on users.id = business_users.user_id").
		Where("business_users.is_validated = ? AND business_users.is_pending = ? AND business_users.is_activated = ?", false, true, false).
		Scan(&pendingOrganisations).Error

	if err != nil {
		return nil, errors.New("failed to fetch pending organisations from business_users")
	}

	// Récupérer les School Users en attente
	var pendingSchoolUsers []response.OrganisationResponse
	err = s.db.Table("school_users").
		Select("school_users.user_id, school_users.school_name AS name, school_users.is_validated, school_users.is_activated, school_users.is_pending, school_users.status, users.profile_image").
		Joins("inner join users on users.id = school_users.user_id").
		Where("school_users.is_validated = ? AND school_users.is_pending = ? AND school_users.is_activated = ?", false, true, false).
		Scan(&pendingSchoolUsers).Error

	if err != nil {
		return nil, errors.New("failed to fetch pending organisations from school_users")
	}

	// Récupérer les Association Users en attente
	var pendingAssociationUsers []response.OrganisationResponse
	err = s.db.Table("association_users").
		Select("association_users.user_id, association_users.association_name AS name, association_users.is_validated, association_users.is_activated, association_users.is_pending, association_users.status, users.profile_image").
		Joins("inner join users on users.id = association_users.user_id").
		Where("association_users.is_validated = ? AND association_users.is_pending = ? AND association_users.is_activated = ?", false, true, false).
		Scan(&pendingAssociationUsers).Error

	if err != nil {
		return nil, errors.New("failed to fetch pending organisations from association_users")
	}

	// Concaténer les résultats
	pendingOrganisations = append(pendingOrganisations, pendingSchoolUsers...)
	pendingOrganisations = append(pendingOrganisations, pendingAssociationUsers...)

	// Récupérer les rôles pour chaque organisation en attente
	for i, organisation := range pendingOrganisations {
		var roles []string
		s.db.Table("roles").
			Joins("inner join user_roles on user_roles.role_id = roles.id").
			Where("user_roles.user_id = ?", organisation.UserID).
			Pluck("name", &roles)

		// Concaténer les rôles
		pendingOrganisations[i].RoleName = strings.Join(roles, ", ")
	}

	// Retourner la réponse encapsulée
	return &response.OrganisationListResponse{
		Organisations: pendingOrganisations,
		Success:       true,
		Message:       "Pending organisations retrieved successfully",
	}, nil
}

// GetSuspendedOrganisations récupère les organisations validées mais inactives (status == Suspendu)
func (s *OrganisationService) GetSuspendedOrganisations() (*response.OrganisationListResponse, error) {
	var suspendedOrganisations []response.OrganisationResponse

	// Récupérer les Business Users suspendus
	err := s.db.Table("business_users").
		Select("business_users.user_id, business_users.company_name AS name, business_users.is_validated, business_users.is_activated, business_users.is_pending, business_users.status, users.profile_image").
		Joins("inner join users on users.id = business_users.user_id").
		Where("business_users.is_validated = ? AND business_users.is_activated = ? AND business_users.is_pending = ?", false, false, false).
		Scan(&suspendedOrganisations).Error

	if err != nil {
		return nil, errors.New("failed to fetch suspended organisations from business_users")
	}

	// Récupérer les School Users suspendus
	var suspendedSchoolUsers []response.OrganisationResponse
	err = s.db.Table("school_users").
		Select("school_users.user_id, school_users.school_name AS name, school_users.is_validated, school_users.is_activated, school_users.is_pending, school_users.status, users.profile_image").
		Joins("inner join users on users.id = school_users.user_id").
		Where("school_users.is_validated = ? AND school_users.is_activated = ? AND school_users.is_pending = ?", false, false, false).
		Scan(&suspendedSchoolUsers).Error

	if err != nil {
		return nil, errors.New("failed to fetch suspended organisations from school_users")
	}

	// Récupérer les Association Users suspendus
	var suspendedAssociationUsers []response.OrganisationResponse
	err = s.db.Table("association_users").
		Select("association_users.user_id, association_users.association_name AS name, association_users.is_validated, association_users.is_activated, association_users.is_pending, association_users.status, users.profile_image").
		Joins("inner join users on users.id = association_users.user_id").
		Where("association_users.is_validated = ? AND association_users.is_activated = ? AND association_users.is_pending = ?", false, false, false).
		Scan(&suspendedAssociationUsers).Error

	if err != nil {
		return nil, errors.New("failed to fetch suspended organisations from association_users")
	}

	// Concaténer les résultats
	suspendedOrganisations = append(suspendedOrganisations, suspendedSchoolUsers...)
	suspendedOrganisations = append(suspendedOrganisations, suspendedAssociationUsers...)

	// Récupérer les rôles pour chaque organisation suspendue
	for i, organisation := range suspendedOrganisations {
		var roles []string
		s.db.Table("roles").
			Joins("inner join user_roles on user_roles.role_id = roles.id").
			Where("user_roles.user_id = ?", organisation.UserID).
			Pluck("name", &roles)

		// Concaténer les rôles
		suspendedOrganisations[i].RoleName = strings.Join(roles, ", ")
	}

	// Retourner la réponse encapsulée
	return &response.OrganisationListResponse{
		Organisations: suspendedOrganisations,
		Success:       true,
		Message:       "Suspended organisations retrieved successfully",
	}, nil
}
