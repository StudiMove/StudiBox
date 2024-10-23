package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type CategoryStore struct {
	db *gorm.DB
}

func NewCategoryStore(db *gorm.DB) *CategoryStore {
	return &CategoryStore{db: db}
}

// CreateCategory ajoute une nouvelle catégorie
func (s *CategoryStore) CreateCategory(category *models.Category) error {
	return s.db.Create(category).Error
}

// FindOrCreateCategory trouve ou crée une catégorie
func (s *CategoryStore) FindOrCreateCategory(categoryName string) (*models.Category, error) {
	var category models.Category
	err := s.db.FirstOrCreate(&category, models.Category{Name: categoryName}).Error
	return &category, err
}

// GetCategories récupère une liste de catégories
func (s *CategoryStore) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := s.db.Find(&categories).Error
	return categories, err
}
