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
func (s *SignUpService) SignUp(signUpRequest models.SignUpRequest) error {
    // Validate the SignUpRequest struct
    if err :=utils.ValidateSignUpRequest(signUpRequest); err != nil {
        return err
    }

    // Check if email already exists
    if s.UserRepository.IsEmailExists(signUpRequest.Email) {
        return constants.ErrEmailExists
    }

    // Check if phone number already exists
    if s.UserRepository.IsPhoneNumberExists(signUpRequest.PhoneNumber) {
        return constants.ErrPhoneNumberExists
    }

    // Check if pancard number already exists
    if s.UserRepository.IsPancardNumberExists(signUpRequest.PancardNumber) {
        return constants.ErrPancardExists
    }

    // Insert user into the database
    if err := s.UserRepository.InsertUser(signUpRequest); err != nil {
        return err
    }

    return nil
}
