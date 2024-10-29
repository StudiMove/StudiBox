package utils

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

// ConvertRequestToModel aide à convertir les requêtes en modèles de données, tout en gérant les images si nécessaire.
func ConvertRequestToModel(c *gin.Context, model interface{}, req interface{}, imageKey string) error {
	// Si l'imageKey est spécifiée, on tente de récupérer les URLs d'image depuis le contexte
	if imageKey != "" {
		imageURLs, exists := c.Get(imageKey)
		if exists {
			if imageURLsStr, ok := imageURLs.(string); ok {
				if err := setImageURLsField(model, imageURLsStr); err != nil {
					return fmt.Errorf("échec de l'attribution des URLs d'image: %w", err)
				}
			}
		}
	}

	// Marshal et Unmarshal pour lier les champs de la requête au modèle
	reqData, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("échec du marshalling de la requête : %w", err)
	}
	if err := json.Unmarshal(reqData, model); err != nil {
		return fmt.Errorf("échec de l'unmarshalling de la requête : %w", err)
	}

	return nil
}

// setImageURLsField tente d'assigner les URLs d'image à un champ "ImageURLs" s'il existe dans le modèle.
func setImageURLsField(model interface{}, imageURLs string) error {
	v := reflect.ValueOf(model).Elem()

	// Vérifier si le champ ImageURLs existe et est de type string
	field := v.FieldByName("ImageURLs")
	if !field.IsValid() {
		return fmt.Errorf("le champ ImageURLs n'existe pas dans le modèle")
	}
	if field.Kind() != reflect.String {
		return fmt.Errorf("le champ ImageURLs n'est pas de type string")
	}

	// Assigner les URLs d'image au champ
	field.SetString(imageURLs)
	return nil
}
