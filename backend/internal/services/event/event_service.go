package event

import (
	"backend/internal/api/models/event/request"
	"backend/internal/api/models/event/response"
	"backend/internal/db/models"
	"backend/internal/services/storage" // Importer votre package de stockage
	"encoding/json"
	"log"
	"mime/multipart" // Ajoutez cette ligne

	"gorm.io/gorm"
)

type EventService struct {
	db             *gorm.DB
	StorageService storage.StorageService // Ajout du StorageService
}

func NewEventService(db *gorm.DB, storageService storage.StorageService) *EventService {
	return &EventService{
		db:             db,
		StorageService: storageService, // Initialisation du StorageService
	}
}

func (s *EventService) CreateEvent(req request.CreateEventRequest) (uint, error) {
	// Créer le modèle Event à partir de la requête
	event := &models.Event{
		UserID:      req.UserID,
		Title:       req.Title,
		Subtitle:    req.Subtitle,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		IsOnline:    req.IsOnline,
		IsVisible:   req.IsVisible,
		UseStudibox: req.UseStudibox,
		TicketPrice: req.TicketPrice,
		TicketStock: req.TicketStock,
		Address:     req.Location.Address,
		City:        req.Location.City,
		Postcode:    req.Location.Postcode,
		Region:      req.Location.Region,
		Country:     req.Location.Country,
		VideoURL:    req.VideoURL,
		ImageURLs:   req.Images,
	}

	// Sauvegarder l'événement principal
	if err := s.db.Create(event).Error; err != nil {
		return 0, err
	}

	// Sauvegarder les descriptions
	if len(req.Descriptions) > 0 {
		descriptions := make([]models.EventDescription, len(req.Descriptions))
		for i, desc := range req.Descriptions {
			descriptions[i] = models.EventDescription{
				EventID:     event.ID,
				Title:       desc.Title,
				Description: desc.Description,
			}
		}
		if err := s.db.Create(&descriptions).Error; err != nil {
			return 0, err
		}
	}

	// Sauvegarder les options
	if len(req.Options) > 0 {
		options := make([]models.EventOption, len(req.Options))
		for i, opt := range req.Options {
			options[i] = models.EventOption{
				EventID:     event.ID,
				Title:       opt.Title,
				Description: opt.Description,
				Price:       opt.Price,
				Stock:       opt.Stock,
			}
		}
		if err := s.db.Create(&options).Error; err != nil {
			return 0, err
		}
	}

	// Sauvegarder les tarifs
	if len(req.Tarifs) > 0 {
		tarifs := make([]models.EventTarif, len(req.Tarifs))
		for i, tarif := range req.Tarifs {
			tarifs[i] = models.EventTarif{
				EventID:     event.ID,
				Title:       tarif.Title,
				Description: tarif.Description,
				Price:       tarif.Price,
				Stock:       tarif.Stock,
			}
		}
		if err := s.db.Create(&tarifs).Error; err != nil {
			return 0, err
		}
	}

	// Associer les catégories
	if err := s.UpdateEventCategories(event.ID, req.Categories); err != nil {
		log.Printf("Erreur lors de la mise à jour des catégories pour l'événement ID %d: %v", event.ID, err)
		return 0, err
	}

	// Associer les tags
	if err := s.UpdateEventTags(event.ID, req.Tags); err != nil {
		log.Printf("Erreur lors de la mise à jour des tags pour l'événement ID %d: %v", event.ID, err)
		return 0, err
	}

	return event.ID, nil
}

// extractTagNames extrait les noms des tags
func (s *EventService) extractTagNames(tags []models.EventTag) []string {
	names := make([]string, len(tags))
	for i, tag := range tags {
		names[i] = tag.Name
	}
	return names
}

// extractOptions extrait les options en réponse
func (s *EventService) extractOptions(options []models.EventOption) []response.EventOptionResponse {
	responses := make([]response.EventOptionResponse, len(options))
	for i, option := range options {
		responses[i] = response.EventOptionResponse{
			ID:          option.ID,
			Title:       option.Title,
			Description: option.Description,
			Price:       option.Price,
		}
	}
	return responses
}

// extractTarifs extrait les tarifs en réponse
func (s *EventService) extractTarifs(tarifs []models.EventTarif) []response.EventTarifResponse {
	responses := make([]response.EventTarifResponse, len(tarifs))
	for i, tarif := range tarifs {
		responses[i] = response.EventTarifResponse{
			ID:          tarif.ID,
			Title:       tarif.Title,
			Description: tarif.Description,
			Price:       tarif.Price,
			Stock:       tarif.Stock,
		}
	}
	return responses
}

func (s *EventService) extractCategoryNames(categories []models.EventCategory) []string {
	names := make([]string, len(categories))
	for i, category := range categories {
		names[i] = category.Name
	}
	return names
}

func (s *EventService) GetEventsByUserID(userID uint) ([]response.EventResponse, error) {
	var events []models.Event
	if err := s.db.Where("user_id= ?", userID).Preload("Categories").Preload("Tags").Find(&events).Error; err != nil {
		log.Printf("Error fetching events for user ID %d: %v", userID, err)
		return nil, err
	}

	eventResponses := make([]response.EventResponse, len(events))
	for i, event := range events {
		eventResponses[i] = s.transformEventToResponse(event)
	}
	return eventResponses, nil
}

func (s *EventService) GetAllEvents() ([]response.EventResponse, error) {
	var events []models.Event
	if err := s.db.Preload("Categories").Preload("Tags").Find(&events).Error; err != nil {
		log.Printf("Error fetching all events: %v", err)
		return nil, err
	}

	eventResponses := make([]response.EventResponse, len(events))
	for i, event := range events {
		eventResponses[i] = s.transformEventToResponse(event)
	}
	return eventResponses, nil
}

func ExtractDescriptionIDs(descriptions []models.EventDescription) []uint {
	ids := make([]uint, len(descriptions))
	for i, description := range descriptions {
		ids[i] = description.ID
	}
	return ids
}

// Méthode de transformation pour convertir un modèle Event en EventResponse

func (s *EventService) transformEventToResponse(event models.Event) response.EventResponse {

	// Désérialiser la chaîne JSON `ImageURLs` en tableau `[]string`

	return response.EventResponse{
		ID:          event.ID,
		Title:       event.Title,
		Subtitle:    event.Subtitle,
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
		IsOnline:    event.IsOnline,
		TicketPrice: event.TicketPrice,
		TicketStock: event.TicketStock,
		Address:     event.Address,
		City:        event.City,
		Postcode:    event.Postcode,
		Region:      event.Region,
		Country:     event.Country,
		Categories:  s.extractCategoryNames(event.Categories),
		Tags:        s.extractTagNames(event.Tags),
		ImageURLs:   event.ImageURLs, // Mappé directement sans transformation
	}
}

func (s *EventService) transformEventToDetailResponse(event models.Event) *response.EventDetailResponse {

	return &response.EventDetailResponse{
		ID:                 event.ID,
		Title:              event.Title,
		Subtitle:           event.Subtitle,
		StartDate:          event.StartDate,
		EndDate:            event.EndDate,
		StartTime:          event.StartTime,
		EndTime:            event.EndTime,
		IsOnline:           event.IsOnline,
		TicketPrice:        event.TicketPrice,
		TicketStock:        event.TicketStock,
		Address:            event.Address,
		City:               event.City,
		Postcode:           event.Postcode,
		Region:             event.Region,
		Country:            event.Country,
		Tags:               s.extractTagNames(event.Tags),
		Categories:         s.extractCategoryNames(event.Categories),
		Options:            s.transformOptions(event.Options),
		Tarifs:             s.transformTarifs(event.Tarifs),
		Descriptions:       s.transformDescriptions(event.Descriptions),
		ImageURLs:          event.ImageURLs,
		IsValidatedByAdmin: event.IsValidatedByAdmin,
		UseStudibox:        event.UseStudibox,
		VideoURL:           event.VideoURL,
	}
}

func (s *EventService) transformOptions(options []models.EventOption) []response.EventOptionResponse {
	responses := make([]response.EventOptionResponse, len(options))
	for i, option := range options {
		responses[i] = response.EventOptionResponse{
			ID:          option.ID,
			Title:       option.Title,
			Description: option.Description,
			Price:       option.Price,
			Stock:       option.Stock,
		}
	}
	return responses
}

func (s *EventService) transformTarifs(tarifs []models.EventTarif) []response.EventTarifResponse {
	responses := make([]response.EventTarifResponse, len(tarifs))
	for i, tarif := range tarifs {
		responses[i] = response.EventTarifResponse{
			ID:          tarif.ID,
			Title:       tarif.Title,
			Description: tarif.Description,
			Price:       tarif.Price,
			Stock:       tarif.Stock,
		}
	}
	return responses
}
func (s *EventService) transformDescriptions(descriptions []models.EventDescription) []response.EventDescriptionResponse {
	responses := make([]response.EventDescriptionResponse, len(descriptions))
	for i, description := range descriptions {
		responses[i] = response.EventDescriptionResponse{
			Title:       description.Title,
			Description: description.Description,
		}
	}
	return responses
}

func (s *EventService) GetEventByID(eventID uint) (*response.EventDetailResponse, error) {
	var event models.Event
	if err := s.db.
		Preload("Categories").
		Preload("Tags").
		Preload("Options").
		Preload("Tarifs").
		Preload("Descriptions").
		Where("id = ?", eventID).
		First(&event).Error; err != nil {
		return nil, err
	}

	return s.transformEventToDetailResponse(event), nil
}

func (s *EventService) UpdateEvent(eventID uint, updatedEvent *models.Event, categoryNames []string, tagNames []string) error {
	var event models.Event
	if err := s.db.First(&event, eventID).Error; err != nil {
		log.Printf("Error finding event by ID: %d, error: %v", eventID, err)
		return err
	}

	// Mettre à jour les champs de l'événement
	event.Title = updatedEvent.Title
	event.Subtitle = updatedEvent.Subtitle
	event.StartDate = updatedEvent.StartDate
	event.EndDate = updatedEvent.EndDate
	event.StartTime = updatedEvent.StartTime
	event.EndTime = updatedEvent.EndTime
	event.IsOnline = updatedEvent.IsOnline
	event.TicketPrice = updatedEvent.TicketPrice
	event.TicketStock = updatedEvent.TicketStock
	event.Address = updatedEvent.Address
	event.City = updatedEvent.City
	event.Postcode = updatedEvent.Postcode
	event.Region = updatedEvent.Region
	event.Country = updatedEvent.Country
	event.UseStudibox = updatedEvent.UseStudibox
	event.VideoURL = updatedEvent.VideoURL
	event.ImageURLs = updatedEvent.ImageURLs

	// Mettre à jour les descriptions, options et tarifs
	if err := s.UpdateEventDescriptions(eventID, updatedEvent.Descriptions); err != nil {
		return err
	}
	if err := s.UpdateEventOptions(eventID, updatedEvent.Options); err != nil {
		return err
	}
	if err := s.UpdateEventTarifs(eventID, updatedEvent.Tarifs); err != nil {
		return err
	}

	// Mettre à jour les catégories et tags
	if err := s.UpdateEventCategories(eventID, categoryNames); err != nil {
		return err
	}
	if err := s.UpdateEventTags(eventID, tagNames); err != nil {
		return err
	}

	return s.db.Save(&event).Error
}

func joinImageURLs(imageURLs []string) string {
	jsonData, err := json.Marshal(imageURLs)
	if err != nil {
		return "[]" // Retourne un tableau vide en cas d'erreur
	}
	return string(jsonData)
}

func (s *EventService) UpdateEventOptions(eventID uint, options []models.EventOption) error {
	// Récupérer les options actuelles pour cet événement par leur `Title`
	var existingOptions []models.EventOption
	if err := s.db.Where("event_id = ?", eventID).Find(&existingOptions).Error; err != nil {
		log.Printf("Erreur lors de la récupération des options existantes pour l'événement ID %d: %v", eventID, err)
		return err
	}

	// Créer une map des options existantes pour un accès rapide par `Title`
	existingOptionsMap := make(map[string]models.EventOption)
	for _, option := range existingOptions {
		existingOptionsMap[option.Title] = option
	}

	// Créer un ensemble des titres des nouvelles options pour vérifier les suppressions
	newOptionsTitles := make(map[string]struct{})
	for _, newOption := range options {
		newOption.EventID = eventID                    // Associe chaque option à cet événement
		newOptionsTitles[newOption.Title] = struct{}{} // Ajoute le titre de l'option à l'ensemble

		if existingOption, exists := existingOptionsMap[newOption.Title]; exists {
			// Si l'option existe, mise à jour
			newOption.ID = existingOption.ID // Préserve l'ID existant pour éviter de créer un doublon
			if err := s.db.Model(&models.EventOption{}).Where("id = ?", existingOption.ID).Updates(newOption).Error; err != nil {
				log.Printf("Erreur lors de la mise à jour de l'option %s pour l'événement ID %d: %v", newOption.Title, eventID, err)
				return err
			}
		} else {
			// Sinon, crée une nouvelle option
			if err := s.db.Create(&newOption).Error; err != nil {
				log.Printf("Erreur lors de la création de l'option %s pour l'événement ID %d: %v", newOption.Title, eventID, err)
				return err
			}
		}
	}

	// Identifier et supprimer les options manquantes
	for title, existingOption := range existingOptionsMap {
		if _, found := newOptionsTitles[title]; !found {
			// Supprime les options qui ne sont pas dans la nouvelle liste
			if err := s.db.Delete(&existingOption).Error; err != nil {
				log.Printf("Erreur lors de la suppression de l'option %s pour l'événement ID %d: %v", existingOption.Title, eventID, err)
				return err
			}
			log.Printf("Option supprimée : %s pour l'événement ID %d", existingOption.Title, eventID)
		}
	}

	return nil
}
func (s *EventService) UpdateEventTarifs(eventID uint, tarifs []models.EventTarif) error {
	// Récupérer les tarifs actuels pour cet événement par leur `Title`
	var existingTarifs []models.EventTarif
	if err := s.db.Where("event_id = ?", eventID).Find(&existingTarifs).Error; err != nil {
		log.Printf("Erreur lors de la récupération des tarifs existants pour l'événement ID %d: %v", eventID, err)
		return err
	}

	// Créer une map des tarifs existants pour un accès rapide par `Title`
	existingTarifsMap := make(map[string]models.EventTarif)
	for _, tarif := range existingTarifs {
		existingTarifsMap[tarif.Title] = tarif
	}

	// Créer un ensemble des titres des nouveaux tarifs pour vérifier les suppressions
	newTarifsTitles := make(map[string]struct{})
	for _, newTarif := range tarifs {
		newTarif.EventID = eventID                   // Associe chaque tarif à cet événement
		newTarifsTitles[newTarif.Title] = struct{}{} // Ajoute le titre du tarif à l'ensemble

		if existingTarif, exists := existingTarifsMap[newTarif.Title]; exists {
			// Si le tarif existe, mise à jour
			newTarif.ID = existingTarif.ID // Préserve l'ID existant pour éviter de créer un doublon
			if err := s.db.Model(&models.EventTarif{}).Where("id = ?", existingTarif.ID).Updates(newTarif).Error; err != nil {
				log.Printf("Erreur lors de la mise à jour du tarif %s pour l'événement ID %d: %v", newTarif.Title, eventID, err)
				return err
			}
		} else {
			// Sinon, crée un nouveau tarif
			if err := s.db.Create(&newTarif).Error; err != nil {
				log.Printf("Erreur lors de la création du tarif %s pour l'événement ID %d: %v", newTarif.Title, eventID, err)
				return err
			}
		}
	}

	// Identifier et supprimer les tarifs manquants
	for title, existingTarif := range existingTarifsMap {
		if _, found := newTarifsTitles[title]; !found {
			// Supprime les tarifs qui ne sont pas dans la nouvelle liste
			if err := s.db.Delete(&existingTarif).Error; err != nil {
				log.Printf("Erreur lors de la suppression du tarif %s pour l'événement ID %d: %v", existingTarif.Title, eventID, err)
				return err
			}
			log.Printf("Tarif supprimé : %s pour l'événement ID %d", existingTarif.Title, eventID)
		}
	}

	return nil
}

func (s *EventService) UpdateEventDescriptions(eventID uint, descriptions []models.EventDescription) error {
	// Récupérer les descriptions actuelles pour cet événement par leur `Title`
	var existingDescriptions []models.EventDescription
	if err := s.db.Where("event_id = ?", eventID).Find(&existingDescriptions).Error; err != nil {
		log.Printf("Erreur lors de la récupération des descriptions existantes pour l'événement ID %d: %v", eventID, err)
		return err
	}

	// Créer une map des descriptions existantes pour un accès rapide par `Title`
	existingDescriptionsMap := make(map[string]models.EventDescription)
	for _, description := range existingDescriptions {
		existingDescriptionsMap[description.Title] = description
	}

	// Créer un ensemble des titres des nouvelles descriptions pour vérifier les suppressions
	newDescriptionsTitles := make(map[string]struct{})
	for _, newDescription := range descriptions {
		newDescription.EventID = eventID                         // Associe chaque description à cet événement
		newDescriptionsTitles[newDescription.Title] = struct{}{} // Ajoute le titre de la description à l'ensemble

		if existingDescription, exists := existingDescriptionsMap[newDescription.Title]; exists {
			// Si la description existe, mise à jour
			newDescription.ID = existingDescription.ID // Préserve l'ID existant pour éviter de créer un doublon
			if err := s.db.Model(&models.EventDescription{}).Where("id = ?", existingDescription.ID).Updates(newDescription).Error; err != nil {
				log.Printf("Erreur lors de la mise à jour de la description %s pour l'événement ID %d: %v", newDescription.Title, eventID, err)
				return err
			}
		} else {
			// Sinon, crée une nouvelle description
			if err := s.db.Create(&newDescription).Error; err != nil {
				log.Printf("Erreur lors de la création de la description %s pour l'événement ID %d: %v", newDescription.Title, eventID, err)
				return err
			}
		}
	}

	// Identifier et supprimer les descriptions manquantes
	for title, existingDescription := range existingDescriptionsMap {
		if _, found := newDescriptionsTitles[title]; !found {
			// Supprime les descriptions qui ne sont pas dans la nouvelle liste
			if err := s.db.Delete(&existingDescription).Error; err != nil {
				log.Printf("Erreur lors de la suppression de la description %s pour l'événement ID %d: %v", existingDescription.Title, eventID, err)
				return err
			}
			log.Printf("Description supprimée : %s pour l'événement ID %d", existingDescription.Title, eventID)
		}
	}

	return nil
}

func (s *EventService) deleteEventAssociations(eventID uint, hardDelete bool) error {
	db := s.db
	if hardDelete {
		db = db.Unscoped()
	}

	associations := []interface{}{
		&models.EventOption{},
		&models.EventTarif{},
		&models.Ticket{},
	}

	for _, assoc := range associations {
		if err := db.Where("event_id = ?", eventID).Delete(assoc).Error; err != nil {
			return err
		}
	}
	return nil
}

func (s *EventService) DeleteEvent(eventID uint) error {
	if err := s.deleteEventAssociations(eventID, false); err != nil {
		return err
	}

	if err := s.db.Delete(&models.Event{}, eventID).Error; err != nil {
		return err
	}
	return nil
}

func (s *EventService) HardDeleteEvent(eventID uint) error {
	if err := s.deleteEventAssociations(eventID, true); err != nil {
		return err
	}

	if err := s.db.Unscoped().Delete(&models.Event{}, eventID).Error; err != nil {
		return err
	}
	return nil
}

// Fonction utilitaire pour extraire les IDs des descriptions
func extractDescriptionIDs(descriptions []models.EventDescription) []uint {
	ids := make([]uint, len(descriptions))
	for i, description := range descriptions {
		ids[i] = description.ID
	}
	return ids
}

func (s *EventService) GetOnlineEventsByUserID(userID uint) ([]response.EventSummaryResponse, error) {
	var events []models.Event
	if err := s.db.Where("host_id = ? AND is_online = ?", userID, true).Find(&events).Error; err != nil {
		return nil, err
	}

	return s.transformToEventSummaryResponses(events), nil
}

func (s *EventService) GetAllOnlineEvents() ([]response.EventSummaryResponse, error) {
	var events []models.Event
	if err := s.db.Where("is_online = ?", true).Find(&events).Error; err != nil {
		return nil, err
	}

	return s.transformToEventSummaryResponses(events), nil
}

func (s *EventService) transformToEventSummaryResponses(events []models.Event) []response.EventSummaryResponse {
	responses := make([]response.EventSummaryResponse, len(events))
	for i, event := range events {
		responses[i] = response.EventSummaryResponse{
			ID:        event.ID,
			Title:     event.Title,
			StartDate: event.StartDate,
			StartTime: event.StartTime,
			IsOnline:  event.IsOnline,
		}
	}
	return responses
}

// UploadEventImage téléverse une image pour un événement
func (s *EventService) UploadEventImage(file multipart.File, fileName string) (response.UploadEventImageResponse, error) {
	url, err := s.StorageService.UploadFile(file, fileName)
	if err != nil {
		log.Printf("Error uploading file: %v", err)
		return response.UploadEventImageResponse{}, err
	}
	// Renvoyer un tableau contenant une seule URL pour correspondre au type attendu
	return response.UploadEventImageResponse{URLs: []string{url}}, nil
}

func (s *EventService) GetOrCreateCategories(categoryNames []string) ([]models.EventCategory, error) {
	var categories []models.EventCategory
	if err := s.db.Where("name IN ?", categoryNames).Find(&categories).Error; err != nil {
		return nil, err
	}

	existingCategoryNames := make(map[string]struct{})
	for _, category := range categories {
		existingCategoryNames[category.Name] = struct{}{}
	}

	for _, name := range categoryNames {
		if _, exists := existingCategoryNames[name]; !exists {
			newCategory := models.EventCategory{Name: name}
			if err := s.db.Create(&newCategory).Error; err != nil {
				return nil, err
			}
			categories = append(categories, newCategory)
		}
	}

	return categories, nil
}

// GetOrCreateTags vérifie ou crée les tags en utilisant le modèle `EventTag`
func (s *EventService) GetOrCreateTags(tagNames []string) ([]models.EventTag, error) {
	var tags []models.EventTag
	if err := s.db.Where("name IN ?", tagNames).Find(&tags).Error; err != nil {
		return nil, err
	}

	existingTagNames := make(map[string]struct{})
	for _, tag := range tags {
		existingTagNames[tag.Name] = struct{}{}
	}

	for _, name := range tagNames {
		if _, exists := existingTagNames[name]; !exists {
			newTag := models.EventTag{Name: name}
			if err := s.db.Create(&newTag).Error; err != nil {
				return nil, err
			}
			tags = append(tags, newTag)
		}
	}

	return tags, nil // Retourner `[]models.EventTag`
}

// GetCategoryByID renvoie une catégorie par ID
func (s *EventService) GetCategoryByID(categoryID uint) (*response.CategoryResponse, error) {
	var category models.EventCategory
	if err := s.db.First(&category, categoryID).Error; err != nil {
		return nil, err
	}
	return &response.CategoryResponse{Name: category.Name}, nil
}

// Récupérer tous les tags
func (s *EventService) GetAllTags() ([]string, error) {
	var tags []struct {
		Name string `json:"name"`
	}
	if err := s.db.Table("event_tags").Select("name").Find(&tags).Error; err != nil {
		return nil, err
	}

	names := make([]string, len(tags))
	for i, tag := range tags {
		names[i] = tag.Name
	}
	return names, nil
}

// GetAllCategories retourne toutes les catégories sous forme de liste de noms
func (s *EventService) GetAllCategories() ([]response.CategoryResponse, error) {
	var categories []models.EventCategory
	if err := s.db.Select("id", "name").Find(&categories).Error; err != nil {
		return nil, err
	}
	return s.transformCategoriesToResponses(categories), nil
}

func (s *EventService) UpdateEventCategories(eventID uint, categoryNames []string) error {
	// Obtenir ou créer les catégories en utilisant leur nom
	categories, err := s.GetOrCreateCategories(categoryNames) // Retourne un []models.EventCategory
	if err != nil {
		log.Printf("Erreur lors de la récupération ou création des catégories: %v", err)
		return err
	}

	// Récupérer l'événement correspondant
	var event models.Event
	if err := s.db.First(&event, eventID).Error; err != nil {
		log.Printf("Erreur lors de la récupération de l'événement avec ID: %d, erreur: %v", eventID, err)
		return err
	}

	// Mettre à jour les catégories pour cet événement avec le type correct
	if err := s.db.Model(&event).Association("Categories").Replace(categories); err != nil {
		log.Printf("Erreur lors de la mise à jour des catégories pour l'événement: %v", err)
		return err
	}

	return nil
}

func (s *EventService) UpdateEventTags(eventID uint, tagNames []string) error {
	// Utiliser `GetOrCreateTags` pour obtenir des modèles `EventTag`
	tags, err := s.GetOrCreateTags(tagNames)
	if err != nil {
		log.Printf("Erreur lors de la récupération ou création des tags: %v", err)
		return err
	}

	var event models.Event
	if err := s.db.First(&event, eventID).Error; err != nil {
		log.Printf("Erreur lors de la récupération de l'événement avec ID: %d, erreur: %v", eventID, err)
		return err
	}

	// Mettre à jour la relation avec `models.EventTag`
	if err := s.db.Model(&event).Association("Tags").Replace(tags); err != nil {
		log.Printf("Erreur lors de la mise à jour des tags pour l'événement: %v", err)
		return err
	}

	return nil
}

// Méthode pour transformer les catégories en ODT de réponse
func (s *EventService) transformCategoriesToResponses(categories []models.EventCategory) []response.CategoryResponse {
	responses := make([]response.CategoryResponse, len(categories))
	for i, category := range categories {
		responses[i] = response.CategoryResponse{
			Name: category.Name,
		}
	}
	return responses
}

// Méthode pour transformer les tags en ODT de réponse
func (s *EventService) transformTagsToResponses(tags []models.EventTag) []response.TagResponse {
	responses := make([]response.TagResponse, len(tags))
	for i, tag := range tags {
		responses[i] = response.TagResponse{
			ID:   tag.ID,
			Name: tag.Name,
		}
	}
	return responses
}

type UploadResult struct {
	URL string `json:"url"`
}

func (s *EventService) UploadEventImages(files []multipart.File, fileNames []string) ([]UploadResult, error) {
	var results []UploadResult
	for i, file := range files {
		// Processus de téléchargement pour chaque fichier
		url, err := s.StorageService.UploadFile(file, fileNames[i])
		if err != nil {
			return nil, err
		}
		results = append(results, UploadResult{URL: url})
	}
	return results, nil
}

// SERVICE MOBILE

// APP MOBILE
func (s *EventService) GetAllEventsMobile() ([]models.Event, error) {
	var events []models.Event
	if err := s.db.
		Preload("Categories").
		Preload("Tags").
		Preload("Options").
		Preload("Tarifs").
		Preload("Descriptions").
		Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}
