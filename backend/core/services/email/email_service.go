package email

import (
	"backend/core/utils"
	"fmt"
	"path/filepath"
)

// EmailEventType définit les types d’événements pour l'envoi d'emails
type EmailEventType string

const (
	EventRegistration  EmailEventType = "registration"
	EventPasswordReset EmailEventType = "password_reset"
	EventCreateEvent   EmailEventType = "create_event"
)

// EmailServiceType contient la configuration pour le service email
type EmailServiceType struct{}

// NewEmailService crée une instance de EmailServiceType
func EmailService() *EmailServiceType {
	return &EmailServiceType{}
}

// SendEmailWithTemplate envoie un email en fonction du type d'événement et du template associé
func (e *EmailServiceType) SendEmailWithTemplate(eventType EmailEventType, to []string, data map[string]string) error {
	// Définir le sujet et le template en fonction de l’événement
	var subject, templateFile string
	switch eventType {
	case EventRegistration:
		subject = "Bienvenue sur StudiMove !"
		templateFile = "registration.mjml"
	case EventPasswordReset:
		subject = "Réinitialisation de votre mot de passe"
		templateFile = "password_reset.mjml"
	case EventCreateEvent:
		subject = data["subject"]
		templateFile = "create_event.mjml"
	default:
		return fmt.Errorf("unhandled email event type: %s", eventType)
	}

	// Convertir le template MJML en HTML
	htmlContent, err := utils.ConvertMJMLToHTML(filepath.Join("templates", templateFile))
	if err != nil {
		return fmt.Errorf("failed to convert template to HTML: %w", err)
	}

	// Envoyer l’email avec le contenu généré
	return utils.SendEmail(to, subject, htmlContent)
}
