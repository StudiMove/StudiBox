package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

// ConvertMJMLToHTML utilise le CLI MJML pour convertir un fichier MJML en HTML avec des données dynamiques
func ConvertMJMLToHTML(templatePath string, data map[string]string) (string, error) {
	// Vérifier si le fichier MJML existe
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return "", fmt.Errorf("le fichier MJML n'existe pas : %s", templatePath)
	}

	// Lire le contenu du fichier MJML
	content, err := os.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la lecture du fichier MJML : %v", err)
	}

	// Interpoler les données dynamiques dans le contenu du template
	tmpl, err := template.New("email").Parse(string(content))
	if err != nil {
		return "", fmt.Errorf("erreur lors de la création du template : %v", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("erreur lors de l'exécution du template : %v", err)
	}

	// Créer un fichier temporaire pour le fichier MJML avec les données interpolées
	tmpMJMLFile, err := os.CreateTemp("", "temp-*.mjml")
	if err != nil {
		return "", fmt.Errorf("impossible de créer le fichier temporaire MJML : %v", err)
	}
	defer os.Remove(tmpMJMLFile.Name())

	// Écrire le contenu dans le fichier temporaire MJML
	if _, err := tmpMJMLFile.Write(buf.Bytes()); err != nil {
		return "", fmt.Errorf("erreur lors de l'écriture dans le fichier temporaire MJML : %v", err)
	}
	tmpMJMLFile.Close()

	// Créer un fichier temporaire pour le HTML converti
	tmpHTMLFile, err := os.CreateTemp("", "temp-*.html")
	if err != nil {
		return "", fmt.Errorf("impossible de créer le fichier temporaire HTML : %v", err)
	}
	defer os.Remove(tmpHTMLFile.Name())

	// Exécuter la commande MJML pour convertir le fichier MJML en HTML
	cmd := exec.Command("mjml", tmpMJMLFile.Name(), "-o", tmpHTMLFile.Name())
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
