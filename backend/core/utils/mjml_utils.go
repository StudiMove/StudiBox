package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// ConvertMJMLToHTML utilise le CLI MJML pour convertir un fichier MJML en HTML
func ConvertMJMLToHTML(templatePath string) (string, error) {
	// Vérifier si le fichier MJML existe
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return "", fmt.Errorf("le fichier MJML n'existe pas : %s", templatePath)
	}

	// Créer un fichier temporaire pour le HTML converti
	tmpHTMLFile, err := os.CreateTemp("", "temp-*.html")
	if err != nil {
		return "", fmt.Errorf("impossible de créer le fichier temporaire HTML : %v", err)
	}
	defer os.Remove(tmpHTMLFile.Name())

	// Exécuter la commande MJML pour convertir le contenu MJML en HTML
	cmd := exec.Command("mjml", templatePath, "-o", tmpHTMLFile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("échec de la conversion MJML : %v - sortie : %s", err, string(output))
	}

	// Lire le contenu HTML généré
	htmlContent, err := os.ReadFile(tmpHTMLFile.Name())
	if err != nil {
		return "", fmt.Errorf("impossible de lire le fichier HTML : %v", err)
	}

	return string(htmlContent), nil
}
