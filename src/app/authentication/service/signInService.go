package service

import (
	"Stock_broker_application/constants"
	"Stock_broker_application/models"
	"Stock_broker_application/repo"
)

// SignInService handles the sign-in logic
type SignInService struct {
	userRepository repo.UserRepository
}

// NewSignInService creates a new SignInService with the provided repository
func NewSignInService(userRepository repo.UserRepository) *SignInService {
	return &SignInService{
		userRepository: userRepository,
	}
}

func (service *SignInService) SignIn(signInRequest models.SignInRequest) error {
	// Retrieve user from the database based on the provided email
	user := service.userRepository.GetUserByEmail(signInRequest.Email)

	// Check if user exists
	if user == nil {
		return constants.ErrUserNotFound
	}
	return nil
}
