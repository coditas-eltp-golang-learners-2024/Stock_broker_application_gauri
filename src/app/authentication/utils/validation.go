// service/validation.go

package utils

import (
	"Stock_broker_application/constants"
	"Stock_broker_application/models"

	"github.com/go-playground/validator/v10"
)

// ValidateSignUpRequest validates the SignUpRequest struct
func ValidateSignUpRequest(req models.SignUpRequest) error {
	validate := validator.New()

	// Validate the SignUpRequest struct
	if err := validate.Struct(req); err != nil {
		// If validation fails, return the corresponding error message
		switch err.(type) {
		case validator.ValidationErrors:
			// Handle validation errors
			for _, err := range err.(validator.ValidationErrors) {
				switch err.Field() {
				case "Name":
					if req.Name == "" {
						return constants.ErrEmptyName
					}
					return constants.ErrInvalidName
				case "Email":
					if req.Email == "" {
						return constants.ErrEmptyEmail
					}
					return constants.ErrInvalidEmail
				case "PhoneNumber":
					if len(string(req.PhoneNumber)) == 0 {
						return constants.ErrEmptyPhoneNumber
					}
					return constants.ErrInvalidPhoneNumber
				case "PancardNumber":
					if req.PancardNumber == "" {
						return constants.ErrEmptyPancardNumber
					}
					return constants.ErrInvalidPancardNumber
				case "Password":
					if req.Password == "" {
						return constants.ErrEmptyPassword
					}
					return constants.ErrMissingPassword
				}
			}
		default:
			// Return a generic error if validation fails
			return constants.ErrValidationFailed
		}
	}

	return nil
}
// ValidateChangePasswordRequest validates the ChangePassword struct
func ValidateChangePasswordRequest(cp models.ChangePassword) error {
    validate := validator.New()

    // Validate the ChangePassword struct
    if err := validate.Struct(cp); err != nil {
        // If validation fails, return the corresponding error message
        switch err.(type) {
        case validator.ValidationErrors:
            // Handle validation errors
            for _, err := range err.(validator.ValidationErrors) {
                switch err.Field() {
                case "Email":
                    if cp.Email == "" {
                        return constants.ErrEmptyEmail
                    }
                    return constants.ErrInvalidEmail
                case "Password":
                    if cp.Password == "" {
                        return constants.ErrEmptyPassword
                    }
                    return constants.ErrInvalidPassword
                case "NewPassword":
                    if cp.NewPassword == "" {
                        return constants.ErrEmptyNewPassword
                    }
                    return constants.ErrInvalidNewPassword
                }
            }
        default:
            // Return a generic error if validation fails
            return constants.ErrValidationFailed
        }
    }

    return nil
}