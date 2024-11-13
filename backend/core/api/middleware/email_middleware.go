package middleware

import (
	email "backend/core/services/email"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EmailMiddleware(emailService *email.EmailServiceType) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Détecter le type d'événement pour l’email
		eventType := c.GetString("email_event_type")
		to := c.GetStringSlice("email_recipients")
		data := c.GetStringMapString("email_data")

		if eventType == "" || len(to) == 0 {
			c.Next() // Si aucun email n'est requis, continuer
			return
		}

		// Envoyer l’email en utilisant le service d’email
		if err := emailService.SendEmailWithTemplate(email.EmailEventType(eventType), to, data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email", "details": err.Error()})
			c.Abort()
			return
		}

		c.Next() // Continuer vers le handler suivant si l'email a été envoyé
	}
}
