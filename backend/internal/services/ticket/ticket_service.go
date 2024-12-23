package ticket

import (
	"backend/internal/db/models"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TicketService représente le service pour gérer les tickets.
type TicketService struct {
	db *gorm.DB
}

// NewTicketService crée une nouvelle instance de TicketService.
func NewTicketService(db *gorm.DB) *TicketService {
	return &TicketService{db: db}
}

// CreateTicket crée un nouveau ticket pour un utilisateur et un événement donnés.
func (s *TicketService) CreateTicketWithDetails(userID uint, eventID uint, tarifIDs []uint, optionIDs []uint) (*models.Ticket, error) {
	// Générer UUID et numéro de ticket
	uuid := uuid.New().String()
	timestamp := time.Now().Format("20060102")
	ticketNumber := s.generateReadableTicketNumber(timestamp)

	// Créer une instance de ticket
	ticket := &models.Ticket{
		UUID:                 uuid,
		UserID:               userID,
		EventID:              eventID,
		TicketNumberReadable: ticketNumber,
		Status:               "valid",
	}

	// Ajouter le ticket à la base de données
	if err := s.db.Create(ticket).Error; err != nil {
		return nil, err
	}

	// Associer les tarifs
	for _, tarifID := range tarifIDs {
		ticketTarif := &models.TicketEventTarif{
			TicketID: ticket.ID,
			TarifID:  tarifID,
		}
		if err := s.db.Create(ticketTarif).Error; err != nil {
			return nil, err
		}
	}

	// Associer les options
	for _, optionID := range optionIDs {
		ticketOption := &models.TicketEventOption{
			TicketID: ticket.ID,
			OptionID: optionID,
		}
		if err := s.db.Create(ticketOption).Error; err != nil {
			return nil, err
		}
	}

	return ticket, nil
}

// GetTicketByID récupère un ticket par son ID unique avec les relations User et Event.
func (s *TicketService) GetTicketByID(ticketID uint) (*models.Ticket, error) {
	var ticket models.Ticket

	// Charger le ticket avec les relations User et Event
	if err := s.db.Preload("User").Preload("Event").First(&ticket, ticketID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("ticket not found")
		}
		return nil, err
	}

	return &ticket, nil
}

// GetTicketsByUserID récupère tous les tickets d'un utilisateur donné avec leurs événements associés.
func (s *TicketService) GetTicketsByUserID(userID uint) ([]models.Ticket, error) {
	var tickets []models.Ticket

	// Rechercher les tickets associés à l'utilisateur
	if err := s.db.Preload("Event").Where("user_id = ?", userID).Find(&tickets).Error; err != nil {
		return nil, err
	}

	return tickets, nil
}

// GetTicketByUUID récupère un ticket par son UUID unique avec les relations User et Event.
// GetTicketByUUID récupère un ticket par son UUID unique avec les relations User, Event, Tarifs et Options associés.
func (s *TicketService) GetTicketByUUID(ticketUUID string) (*models.Ticket, error) {
	var ticket models.Ticket

	fmt.Printf("Searching for ticket with UUID: %s\n", ticketUUID) // Log pour vérifier l'UUID

	// Charger le ticket avec les relations User, Event, Tarifs et Options
	if err := s.db.
		Preload("User").
		Preload("Event").
		Preload("Tarifs.Tarif").   // Charger les détails des tarifs associés au ticket
		Preload("Options.Option"). // Charger les détails des options associées au ticket
		Where("uuid = ?", ticketUUID).
		First(&ticket).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("Ticket not found for UUID: %s\n", ticketUUID) // Log si non trouvé
			return nil, errors.New("ticket not found")
		}
		fmt.Printf("Database error: %s\n", err.Error()) // Log en cas d'erreur DB
		return nil, err
	}

	return &ticket, nil
}

// CancelTicket met à jour le statut d'un ticket pour indiquer qu'il est annulé.
func (s *TicketService) CancelTicket(ticketID uint) error {
	// Mettre à jour le statut du ticket
	if err := s.db.Model(&models.Ticket{}).Where("id = ?", ticketID).Update("status", "cancelled").Error; err != nil {
		return err
	}

	return nil
}

// MarkTicketAsUsed marque un ticket comme utilisé.
func (s *TicketService) MarkTicketAsUsed(ticketID uint) error {
	// Mettre à jour le statut du ticket
	if err := s.db.Model(&models.Ticket{}).Where("id = ?", ticketID).Update("status", "used").Error; err != nil {
		return err
	}

	return nil
}

// generateReadableTicketNumber génère un numéro lisible unique pour le ticket.
func (s *TicketService) generateReadableTicketNumber(timestamp string) string {
	var lastTicket models.Ticket

	// Trouver le dernier ticket du même jour pour incrémenter le compteur
	if err := s.db.Where("ticket_number_readable LIKE ?", "TICKET-"+timestamp+"%").
		Order("id desc").First(&lastTicket).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Aucun ticket trouvé pour ce jour, commencer à 1
			return "TICKET-" + timestamp + "-00001"
		}
	}

	// Extraire le dernier numéro et incrémenter
	lastNumber := lastTicket.TicketNumberReadable[len(lastTicket.TicketNumberReadable)-5:] // Les 5 derniers chiffres
	newNumber := 1
	fmt.Sscanf(lastNumber, "%d", &newNumber)
	newNumber++

	return fmt.Sprintf("TICKET-%s-%05d", timestamp, newNumber)
}

// GetTicketsWithDetailsByUserID récupère les tickets d'un utilisateur avec les détails utilisateur, événement, options et tarifs.
func (s *TicketService) GetTicketsWithDetailsByUserID(userID uint) ([]TicketWithDetails, error) {
	var tickets []models.Ticket

	// Charger les tickets avec les relations User, Event, Options et Tarifs
	if err := s.db.Preload("User").Preload("Event.Options").Preload("Event.Tarifs").Where("user_id = ?", userID).Find(&tickets).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no tickets found for this user")
		}
		return nil, err
	}

	// Construire la réponse avec les informations détaillées
	var detailedTickets []TicketWithDetails
	for _, ticket := range tickets {
		detailedTickets = append(detailedTickets, TicketWithDetails{
			TicketNumberReadable: ticket.TicketNumberReadable,
			Status:               ticket.Status,
			UserFirstName:        ticket.User.FirstName,
			UserLastName:         ticket.User.LastName,
			EventTitle:           ticket.Event.Title,
			EventOptions:         ticket.Event.Options,
			EventTarifs:          ticket.Event.Tarifs,
		})
	}

	return detailedTickets, nil
}

// TicketWithDetails représente un ticket avec des informations détaillées, incluant les options et tarifs de l'événement.
type TicketWithDetails struct {
	TicketNumberReadable string               `json:"ticket_number_readable"`
	Status               string               `json:"status"`
	UserFirstName        string               `json:"user_first_name"`
	UserLastName         string               `json:"user_last_name"`
	EventTitle           string               `json:"event_title"`
	EventOptions         []models.EventOption `json:"event_options"`
	EventTarifs          []models.EventTarif  `json:"event_tarifs"`
}

func (s *TicketService) GetEventIDsAndUUIDsByUserID(userID uint) ([]map[string]interface{}, error) {
	tickets, err := s.GetTicketsByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Construire une liste des données nécessaires.
	var results []map[string]interface{}
	for _, ticket := range tickets {
		results = append(results, map[string]interface{}{
			"eventID": ticket.EventID,
			"uuid":    ticket.UUID,
		})
	}

	return results, nil
}
