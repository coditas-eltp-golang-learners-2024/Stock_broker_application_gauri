package service

import (
	"Stock_broker_application/constants"
	"Stock_broker_application/models"
	"Stock_broker_application/repo"
	"Stock_broker_application/utils"
	"errors"
)

type PasswordService struct {
	UserRepository repo.UserRepository
}

func NewPasswordService(userRepository repo.UserRepository) *PasswordService {
	return &PasswordService{
		UserRepository: userRepository,
	}
}

// ChangePasswordService method in the service layer
func (service *PasswordService) ChangePasswordService(userInput *models.ChangePassword) error {
	if userInput.Password == userInput.NewPassword {
		return errors.New("new password cannot be the same as the old password")
	}

	if err := utils.ValidateChangePasswordRequest(*userInput); err != nil {
		return err
	}

	if !service.UserRepository.CheckOldPassword(userInput.Email, userInput.Password) {
		return constants.ErrChangingPassword
	}

	if err := service.UserRepository.UpdatePassword(userInput.Email, userInput.NewPassword); err != nil {
		return err
	}

	return nil
}
