package utils

import (
	"encoding/csv"
	"os"
)

// ExportToCSV crée un fichier CSV temporaire et écrit des données dedans
func ExportToCSV(headers []string, rows [][]string, delimiter rune) (string, error) {
	// Créer un fichier CSV temporaire
	file, err := os.CreateTemp("", "export-*.csv")
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Initialiser un écrivain CSV avec le délimiteur choisi
	writer := csv.NewWriter(file)
	writer.Comma = delimiter
	defer writer.Flush()

	// Écrire les en-têtes CSV
	if err := writer.Write(headers); err != nil {
		return "", err
	}

	// Écrire les données des lignes
	for _, row := range rows {
		if err := writer.Write(row); err != nil {
			return "", err
		}
	}

	// Finaliser l'écriture
	writer.Flush()

	// Vérifier les erreurs potentielles
	if err := writer.Error(); err != nil {
		return "", err
	}

	// Retourner le nom du fichier temporaire
	return file.Name(), nil
}
