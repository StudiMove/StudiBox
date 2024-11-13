package event

import (
	"backend/database/models"

	"gorm.io/gorm"
)

type EventTagStoreType struct {
	db *gorm.DB
}

func EventTagStore(db *gorm.DB) *EventTagStoreType {
	return &EventTagStoreType{db: db}
}

// CreateTag ajoute un nouveau tag
func (s *EventTagStoreType) CreateTag(tag *models.EventTag) error {
	return s.db.Create(tag).Error
}

// FindOrCreateTag trouve ou crée un tag
func (s *EventTagStoreType) FindOrCreateTag(tagName string) (*models.EventTag, error) {
	var tag models.EventTag
	err := s.db.FirstOrCreate(&tag, models.EventTag{Name: tagName}).Error
	return &tag, err
}

// GetTags récupère une liste de tous les tags
func (s *EventTagStoreType) GetTags() ([]models.EventTag, error) {
	var tags []models.EventTag
	err := s.db.Find(&tags).Error
	return tags, err
}

func (s *EventTagStoreType) GetTagsByIDs(ids []int64) ([]models.EventTag, error) {
	var tags []models.EventTag
	err := s.db.Where("id IN ?", ids).Find(&tags).Error
	return tags, err
}

// GetTagNamesByIDs retourne les noms des tags à partir d'une liste d'IDs
func (s *EventTagStoreType) GetTagNamesByIDs(tagIDs []int64) ([]string, error) {
	var tags []models.EventTag
	if err := s.db.Where("id IN ?", tagIDs).Find(&tags).Error; err != nil {
		return nil, err
	}

	// Extraire les noms
	names := make([]string, len(tags))
	for i, tag := range tags {
		names[i] = tag.Name
	}
	return names, nil
}
