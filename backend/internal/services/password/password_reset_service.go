package password

import (
	"backend/internal/api/models/password/response"
	"backend/internal/db/models"
	"backend/internal/utils"
	"errors"
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"time"

	"gorm.io/gorm"
)

type PasswordResetService struct {
	db *gorm.DB
}

// NewPasswordResetService crée une nouvelle instance de PasswordResetService
func NewPasswordResetService(db *gorm.DB) *PasswordResetService {
	return &PasswordResetService{db: db}
}

// Fonction pour envoyer un email sans authentification
func sendEmailReset(to string, resetCode int) error {
	smtpHost := os.Getenv("SMTP_SERVER") // Utiliser SMTP_SERVER de l'environnement
	smtpPort := os.Getenv("SMTP_PORT")   // Utiliser SMTP_PORT de l'environnement

	if smtpHost == "" || smtpPort == "" {
		return fmt.Errorf("configuration SMTP manquante")
	}

	from := "no-reply@yourdomain.com"
	subject := "Code de réinitialisation de mot de passe"
	body := fmt.Sprintf("Votre code de réinitialisation est : %d", resetCode)
	message := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body))

	// Envoyer l'email sans authentification
	err := smtp.SendMail(smtpHost+":"+smtpPort, nil, from, []string{to}, message)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi de l'email :", err)
		return err
	}
	fmt.Println("Email envoyé avec succès à", to)
	return nil
}

// SendResetCode génère un code de réinitialisation et envoie un email avec ce code
func (s *PasswordResetService) SendResetCode(email string, userID uint) (int, error) {
	resetCode := generateSixDigitCode()

	var passwordReset models.PasswordReset
	err := s.db.Where("user_id = ?", userID).First(&passwordReset).Error

	if err == nil {
		// Mise à jour du code existant si l'entrée est déjà présente
		passwordReset.ResetCode = resetCode
		passwordReset.Expiration = time.Now().Add(5 * time.Minute)
		if updateErr := s.db.Save(&passwordReset).Error; updateErr != nil {
			return 0, updateErr
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		// Création d'un nouvel enregistrement si aucun code précédent n'existe
		passwordReset = models.PasswordReset{
			UserID:     userID,
			ResetCode:  resetCode,
			Expiration: time.Now().Add(5 * time.Minute),
		}
		if createErr := s.db.Create(&passwordReset).Error; createErr != nil {
			return 0, createErr
		}
	} else {
		return 0, err
	}

	// Envoi de l'email en utilisant mails.sendEmail
	if emailErr := sendEmailReset(email, resetCode); emailErr != nil {
		return 0, emailErr
	}

	return resetCode, nil
}

// generateSixDigitCode génère un code aléatoire à 6 chiffres
func generateSixDigitCode() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000000) // Génère un nombre aléatoire entre 0 et 999999
}

// GetResetCode récupère le code de réinitialisation pour un email donné
func (s *PasswordResetService) GetResetCode(email string) (int, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return 0, errors.New("user not found")
	}

	var passwordReset models.PasswordReset
	if err := s.db.Where("user_id = ?", user.ID).First(&passwordReset).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("reset code not found")
		}
		return 0, err
	}

	// Vérifie l'expiration du code
	if time.Now().After(passwordReset.Expiration) {
		return 0, errors.New("reset code has expired")
	}

	return passwordReset.ResetCode, nil
}

// sendEmail envoie un e-mail avec le code de réinitialisation
func sendEmail(email string, resetCode int) error {
	// Implémentation de l'envoi d'email, utilisant `resetCode` pour le contenu
	// Exemple de code utilisant une bibliothèque d'envoi d'emails
	return nil
}

// UpdatePassword met à jour le mot de passe d'un utilisateur avec l'email fourni
func (s *PasswordResetService) UpdatePassword(email string, newPassword string) (*response.UpdatePasswordResponse, error) {
	var user models.User

	// Récupère l'utilisateur par email
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &response.UpdatePasswordResponse{
				Success: false,
				Message: "User not found",
			}, nil
		}
		return nil, err
	}

	// Hachage du nouveau mot de passe
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return &response.UpdatePasswordResponse{
			Success: false,
			Message: "Failed to hash the password",
		}, err
	}

	// Mise à jour du mot de passe
	user.Password = hashedPassword
	if saveErr := s.db.Save(&user).Error; saveErr != nil {
		return &response.UpdatePasswordResponse{
			Success: false,
			Message: "Failed to update password",
		}, saveErr
	}

	return &response.UpdatePasswordResponse{
		Success: true,
		Message: "Password updated successfully",
	}, nil
}
