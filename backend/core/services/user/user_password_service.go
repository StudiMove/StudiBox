package user

import (
	stores "backend/core/stores/user"
	"backend/core/utils"
	"backend/database/models"
	"errors"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type UserPasswordServiceType struct {
	passwordStore *stores.UserPasswordStoreType
	userStore     *stores.UserStoreType
}

func UserPasswordService(passwordStore *stores.UserPasswordStoreType, userStore *stores.UserStoreType) *UserPasswordServiceType {
	return &UserPasswordServiceType{
		passwordStore: passwordStore,
		userStore:     userStore,
	}
}

// Générer un code de réinitialisation et le stocker
func (s *UserPasswordServiceType) GenerateResetCode(email string) (int, error) {
	user, err := s.userStore.GetByEmail(email)
	if user == nil || err != nil {
		return 0, errors.New("utilisateur introuvable")
	}

	resetCode := generateSixDigitCode()
	userPassword, err := s.passwordStore.GetByUserID(user.ID)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}

	// Mettre à jour ou créer un enregistrement PasswordReset
	if userPassword != nil {
		userPassword.ResetCode = resetCode
		userPassword.Expiration = time.Now().Add(5 * time.Minute)
		err = s.passwordStore.Update(userPassword)
	} else {
		userPassword = &models.PasswordReset{
			UserID:     user.ID,
			ResetCode:  resetCode,
			Expiration: time.Now().Add(5 * time.Minute),
		}
		err = s.passwordStore.Create(userPassword)
	}
	if err != nil {
		return 0, err
	}

	return resetCode, nil
}

// Mettre à jour le mot de passe avec le code de réinitialisation
func (s *UserPasswordServiceType) UpdatePasswordWithCode(email string, code int, newPassword string) error {
	user, err := s.userStore.GetByEmail(email)
	if user == nil || err != nil {
		return errors.New("utilisateur introuvable")
	}

	userPassword, err := s.passwordStore.GetByUserID(user.ID)
	if err != nil || userPassword.ResetCode != code || userPassword.Expiration.Before(time.Now()) {
		return errors.New("code de réinitialisation invalide ou expiré")
	}

	// Hash du nouveau mot de passe
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// Mettre à jour le mot de passe
	return s.passwordStore.UpdateUserPassword(user.ID, hashedPassword)
}

func generateSixDigitCode() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000000)
}
