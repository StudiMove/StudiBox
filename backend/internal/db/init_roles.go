package db

import (
	"backend/internal/db/models" // Assurez-vous que ce chemin est correct

	"gorm.io/gorm"
)

func InitRoles(db *gorm.DB) {
	// Créer les rôles
	roles := []models.Role{
		{Name: "admin"},
		{Name: "business"},
		{Name: "school"},
		{Name: "association"},
		{Name: "user"},
	}

	// Vérifier si les rôles existent déjà
	for _, role := range roles {
		var existingRole models.Role
		// Vérifiez si le rôle existe déjà
		if err := db.Where("name = ?", role.Name).First(&existingRole).Error; err != nil {
			// Si le rôle n'existe pas, l'ajouter
			db.Create(&role)
		}
	}

	// Initialiser les tags
	tags := []models.EventTag{
		{Name: "TikTok"},
		{Name: "Hype"},
		{Name: "Fun"},
		{Name: "Drole"},
	}
	// Vérifier si les tags existent déjà
	for _, tag := range tags {
		var existingTag models.EventTag
		if err := db.Where("name = ?", tag.Name).First(&existingTag).Error; err != nil {
			// Si le tag n'existe pas, l'ajouter
			db.Create(&tag)
		}
	}

	// Initialiser les catégories
	categories := []models.EventCategory{
		{Name: "Musique"},
		{Name: "Soirée"},
		{Name: "Voyage"},
		{Name: "Bar"},
	}
	// Vérifier si les catégories existent déjà
	for _, category := range categories {
		var existingCategory models.EventCategory
		if err := db.Where("name = ?", category.Name).First(&existingCategory).Error; err != nil {
			// Si la catégorie n'existe pas, l'ajouter
			db.Create(&category)
		}
	}

}
