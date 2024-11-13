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
	store *stores.UserPasswordStoreType
}

func UserPasswordService(store *stores.UserPasswordStoreType) *UserPasswordServiceType {
	return &UserPasswordServiceType{store: store}
}

// SendResetCode génère et stocke un code de réinitialisation pour un utilisateur
func (s *UserPasswordServiceType) SendResetCode(userID uint) (int, error) {
	resetCode := generateSixDigitCode()
	userPassword, err := s.store.GetByUserID(userID)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}

	if userPassword != nil {
		userPassword.ResetCode = resetCode
		userPassword.Expiration = time.Now().Add(5 * time.Minute)
		err = s.store.Update(userPassword)
	} else {
		userPassword = &models.PasswordReset{
			UserID:     userID,
			ResetCode:  resetCode,
			Expiration: time.Now().Add(5 * time.Minute),
		}
		err = s.store.Create(userPassword)
	}

	if err != nil {
		return 0, err
	}
	return resetCode, nil
}

// UpdatePassword met à jour le mot de passe d'un utilisateur par ID
func (s *UserPasswordServiceType) UpdatePassword(userID uint, newPassword string) error {
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}
	return s.store.UpdateUserPassword(userID, hashedPassword)
}

func generateSixDigitCode() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000000)
}
