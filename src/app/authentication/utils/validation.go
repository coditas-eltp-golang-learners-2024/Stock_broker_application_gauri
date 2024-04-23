package utils

import (
	"Stock_broker_application/constants"
	"Stock_broker_application/models"

	"github.com/go-playground/validator/v10"
)

// ValidateSignUpRequest validates the SignUpRequest struct
func ValidateSignUpRequest(request models.SignUpRequest) error {
	validate := validator.New()

	// Validate the SignUpRequest struct
	if err := validate.Struct(request); err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			for _, err := range err.(validator.ValidationErrors) {
				switch err.Field() {
				case "Name":
					if request.Name == "" {
						return constants.ErrEmptyName
					}
					return constants.ErrInvalidName
				case "Email":
					if request.Email == "" {
						return constants.ErrEmptyEmail
					}
					return constants.ErrInvalidEmail
				case "PhoneNumber":
					if request.PhoneNumber == 0 {
						return constants.ErrEmptyPhoneNumber
					}
					return constants.ErrInvalidPhoneNumber
				case "PancardNumber":
					if request.PancardNumber == "" {
						return constants.ErrEmptyPancardNumber
					}
					return constants.ErrInvalidPancardNumber
				case "Password":
					if request.Password == "" {
						return constants.ErrEmptyPassword
					}
					return constants.ErrMissingPassword
				}
			}
		default:
			return constants.ErrValidationFailed
		}
	}

	return nil
}

// ValidateChangePasswordRequest validates the ChangePassword struct
func ValidateChangePasswordRequest(changePassword models.ChangePassword) error {
	validate := validator.New()

	// Validate the ChangePassword struct
	if err := validate.Struct(changePassword); err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			for _, err := range err.(validator.ValidationErrors) {
				switch err.Field() {
				case "Email":
					if changePassword.Email == "" {
						return constants.ErrEmptyEmail
					}
					return constants.ErrInvalidEmail
				case "Password":
					if changePassword.Password == "" {
						return constants.ErrEmptyPassword
					}
					if len(changePassword.Password) < 8 {
						return constants.ErrMinLengthPassword
					}
					return constants.ErrInvalidPassword
				case "NewPassword":
					if changePassword.NewPassword == "" {
						return constants.ErrEmptyNewPassword
					}
					if len(changePassword.NewPassword) < 8 {
						return constants.ErrMinLengthPassword
					}
					return constants.ErrInvalidNewPassword
				}
			}
		default:
			return constants.ErrValidationFailed
		}
	}

	return nil
}
