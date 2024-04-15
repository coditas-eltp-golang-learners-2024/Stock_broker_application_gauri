// service/SignUpService.go
package service

import (
	"Stock_broker_application/constants"
	"Stock_broker_application/models"
	"Stock_broker_application/repo"
	"Stock_broker_application/utils"
)

// SignUpService handles the signup logic
type SignUpService struct {
	userRepository repo.UserRepository
}

// NewSignUpService creates a new SignUpService with the provided repository
func NewSignUpService(userRepository repo.UserRepository) *SignUpService {
	return &SignUpService{
		userRepository: userRepository,
	}
}

func (service *SignUpService) SignUp(signUpRequest models.SignUpRequest) error {
	// Validate the SignUpRequest struct
	if err := utils.ValidateSignUpRequest(signUpRequest); err != nil {
		return err
	}

	// Check if email already exists
	if service.userRepository.IsEmailExists(signUpRequest.Email) {
		return constants.ErrEmailExists
	}

	// Check if phone number already exists
	if service.userRepository.IsPhoneNumberExists(signUpRequest.PhoneNumber) {
		return constants.ErrPhoneNumberExists
	}

	// Check if pancard number already exists
	if service.userRepository.IsPancardNumberExists(signUpRequest.PancardNumber) {
		return constants.ErrPancardExists
	}

	// Insert user into the database
	if err := service.userRepository.InsertUser(signUpRequest); err != nil {
		return err
	}

	return nil
}
