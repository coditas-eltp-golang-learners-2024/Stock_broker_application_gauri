package service

import (
	"Stock_broker_application/constants"
	"Stock_broker_application/models"
	"Stock_broker_application/repo"
	"Stock_broker_application/utils"
)

type PasswordService struct {
	userRepository repo.UserRepository
}

func NewPasswordService(userRepository repo.UserRepository) *PasswordService {
	return &PasswordService{
		userRepository: userRepository,
	}
}

// ChangePasswordService method in the service layer
func (service *PasswordService) ChangePasswordService(userInput *models.ChangePassword) error {
	if userInput.Password == userInput.NewPassword {
		return constants.ErrNewPasswordSameAsOld
	}

	if err := utils.ValidateChangePasswordRequest(*userInput); err != nil {
		return constants.ErrInvalidChangeRequest
	}

	if !service.userRepository.CheckOldPassword(userInput.Email, userInput.Password) {
		return constants.ErrChangingPassword
	}

	if err := service.userRepository.UpdatePassword(userInput.Email, userInput.NewPassword); err != nil {
		return constants.ErrChangingPassword
	}

	return nil
}
