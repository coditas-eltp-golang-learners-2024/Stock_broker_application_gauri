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
	UserRepository repo.UserRepository
}

// NewSignUpService creates a new SignUpService with the provided repository
func NewSignUpService(userRepository repo.UserRepository) *SignUpService {
	return &SignUpService{
		UserRepository: userRepository,
	}
}

// SignUp processes the signup request
// @Summary Process signup request
// @Description Process signup request and create a new user account
// @Param signUpRequest body models.SignUpRequest true "Sign-up request body"
// @Success 200 {string} string "User signed up successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 409 {object} string "Conflict: Email/Phone"
// @Failure 500 {object} string "Internal server error"
// @Tags SignUp
// @Router /signup [post]

func (service *SignUpService) SignUp(signUpRequest models.SignUpRequest) error {
	// Validate the SignUpRequest struct
	if err := utils.ValidateSignUpRequest(signUpRequest); err != nil {
		return err
	}

	// Check if email already exists
	if service.UserRepository.IsEmailExists(signUpRequest.Email) {
		return constants.ErrEmailExists
	}

	// Check if phone number already exists
	if service.UserRepository.IsPhoneNumberExists(signUpRequest.PhoneNumber) {
		return constants.ErrPhoneNumberExists
	}

	// Check if pancard number already exists
	if service.UserRepository.IsPancardNumberExists(signUpRequest.PancardNumber) {
		return constants.ErrPancardExists
	}

	// Insert user into the database
	if err := service.UserRepository.InsertUser(signUpRequest); err != nil {
		return err
	}

	return nil
}
