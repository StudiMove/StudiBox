package mails

import (
	"fmt"
	"net/smtp"
)

// Fonction pour envoyer un email avec le code de réinitialisation
func sendEmail(to string, resetCode int) error {
	// Configuration de MailHog
	smtpHost := "localhost"
	smtpPort := "1025"

	// Détails de l'email
	from := "no-reply@yourdomain.com"
	subject := "Code de réinitialisation de mot de passe"
	body := fmt.Sprintf("Votre code de réinitialisation est : %d", resetCode)
	message := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body))

	// Envoi de l'email
	auth := smtp.PlainAuth("", "", "", smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi de l'email :", err)
		return err
	}
	fmt.Println("Email envoyé avec succès à", to)
	return nil
}
