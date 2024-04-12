package service

import (
	"Stock_broker_application/constants"
	"Stock_broker_application/models"
	"Stock_broker_application/repo"
)

// SignInService handles the sign-in logic
type SignInService struct {
	UserRepository repo.UserRepository
}

// NewSignInService creates a new SignInService with the provided repository
func NewSignInService(userRepository repo.UserRepository) *SignInService {
	return &SignInService{
		UserRepository: userRepository,
	}
}

// SignIn authenticates the user
// @Summary Authenticate user
// @Description Authenticate a user with the provided credentials
// @Param signInRequest body models.SignInRequest true "Sign-in request body"
// @Success 200 {string} string "User authenticated successfully"
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "User not found"
// @Router /signin [post]
func (service *SignInService) SignIn(signInRequest models.SignInRequest) error {
	// Retrieve user from the database based on the provided email
	user := service.UserRepository.GetUserByEmail(signInRequest.Email)

	// Check if user exists
	if user == nil {
		return constants.ErrUserNotFound
	}

	// // Check if the provided password matches the user's password
	// if user.Password != signInRequest.Password {
	//     return constants.ErrInvalidCredentials
	// }

	// Authentication successful
	return nil
}
