// service/SignUpService.go
package service

import (
    "Stock_broker_application/models"
    "Stock_broker_application/repo"
    "errors"
    "github.com/go-playground/validator/v10"
)

// SignUpService handles the signup logic
type SignUpService struct {
    UserRepository repo.UserRepository
    Validator      *validator.Validate
}

// NewSignUpService creates a new SignUpService with the provided repository
func NewSignUpService(userRepository repo.UserRepository) *SignUpService {
    return &SignUpService{
        UserRepository: userRepository,
        Validator:      validator.New(),
    }
}

// SignUp processes the signup request
func (s *SignUpService) SignUp(signUpRequest models.SignUpRequest) error {
    // Validate the SignUpRequest struct
    if err := s.Validator.Struct(signUpRequest); err != nil {
        return err
    }

    // Check if email already exists
    if s.UserRepository.IsEmailExists(signUpRequest.Email) {
        return errors.New("email already exists")
    }

    // Check if phone number already exists
    if s.UserRepository.IsPhoneNumberExists(signUpRequest.PhoneNumber) {
        return errors.New("phone number already exists")
    }

    // Check if pancard number already exists
    if s.UserRepository.IsPancardNumberExists(signUpRequest.PancardNumber) {
        return errors.New("pancard number already exists")
    }

    // Insert user into the database
    if err := s.UserRepository.InsertUser(signUpRequest); err != nil {
        return err
    }

    return nil
}
