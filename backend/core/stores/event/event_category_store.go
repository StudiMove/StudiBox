package event

import (
	"backend/database/models"

	"gorm.io/gorm"
)

type EventCategoryStoreType struct {
	db *gorm.DB
}

func EventCategoryStore(db *gorm.DB) *EventCategoryStoreType {
	return &EventCategoryStoreType{db: db}
}

// CreateCategory ajoute une nouvelle catégorie
func (s *EventCategoryStoreType) CreateCategory(category *models.EventCategory) error {
	return s.db.Create(category).Error
}

// FindOrCreateCategory trouve ou crée une catégorie
func (s *EventCategoryStoreType) FindOrCreateCategory(categoryName string) (*models.EventCategory, error) {
	var category models.EventCategory
	err := s.db.FirstOrCreate(&category, models.EventCategory{Name: categoryName}).Error
	return &category, err
}

// GetCategories récupère toutes les catégories
func (s *EventCategoryStoreType) GetCategories() ([]models.EventCategory, error) {
	var categories []models.EventCategory
	err := s.db.Find(&categories).Error
	return categories, err
}

func (s *EventCategoryStoreType) GetCategoriesByIDs(ids []int64) ([]models.EventCategory, error) {
	var categories []models.EventCategory
	err := s.db.Where("id IN ?", ids).Find(&categories).Error
	return categories, err
}

func (s *EventCategoryStoreType) GetCategoryNamesByIDs(categoryIDs []int64) ([]string, error) {
	var categories []models.EventCategory
	if err := s.db.Where("id IN ?", categoryIDs).Find(&categories).Error; err != nil {
		return nil, err
	}

	// Extraire les noms
	names := make([]string, len(categories))
	for i, category := range categories {
		names[i] = category.Name
	}
	return names, nil
}
