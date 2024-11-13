package utils

import (
	"backend/config"
	"fmt"
	"net/smtp"
	"strings"
)

// SendEmail envoie un email avec le contenu HTML
func SendEmail(to []string, subject, htmlContent string) error {
	mailConfig := config.AppConfig

	// Construire le corps du message
	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n%s",
		mailConfig.EmailFrom, strings.Join(to, ","), subject, htmlContent)

	// Préparer l'adresse SMTP
	smtpAddress := fmt.Sprintf("%s:%s", mailConfig.SMTPHost, mailConfig.SMTPPort)

	// Si les informations d'authentification ne sont pas fournies, envoyer sans authentification
	if mailConfig.SMTPUser == "" || mailConfig.SMTPPass == "" {
		return smtp.SendMail(smtpAddress, nil, mailConfig.EmailFrom, to, []byte(message))
	}

	// Si l'authentification est nécessaire, configurer l'authentification
	auth := smtp.PlainAuth("", mailConfig.SMTPUser, mailConfig.SMTPPass, mailConfig.SMTPHost)
	err := smtp.SendMail(smtpAddress, auth, mailConfig.EmailFrom, to, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Println("Email envoyé avec succès à :", to)
	return nil
}
