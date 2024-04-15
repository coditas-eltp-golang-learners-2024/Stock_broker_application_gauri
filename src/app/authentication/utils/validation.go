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
                    if req.PhoneNumber == "" {
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
