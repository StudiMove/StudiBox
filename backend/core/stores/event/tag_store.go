package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type TagStore struct {
	db *gorm.DB
}

func NewTagStore(db *gorm.DB) *TagStore {
	return &TagStore{db: db}
}

// CreateTag ajoute un nouveau tag
func (s *TagStore) CreateTag(tag *models.Tag) error {
	return s.db.Create(tag).Error
}

// FindOrCreateTag trouve ou crée un tag
func (s *TagStore) FindOrCreateTag(tagName string) (*models.Tag, error) {
	var tag models.Tag
	err := s.db.FirstOrCreate(&tag, models.Tag{Name: tagName}).Error
	return &tag, err
}

// GetTags récupère une liste de tags
func (s *TagStore) GetTags() ([]models.Tag, error) {
	var tags []models.Tag
	err := s.db.Find(&tags).Error
	return tags, err
}
