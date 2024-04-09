// service/validation.go

package utils

import (
	"Stock_broker_application/constants"
	"Stock_broker_application/models"
	"github.com/go-playground/validator/v10"
)

// validateSignUpRequest validates the SignUpRequest struct with custom validation rules
func ValidateSignUpRequest(req models.SignUpRequest) error {
    validate := validator.New()
    validate.RegisterValidation("isAlpha", func(fl validator.FieldLevel) bool {
        name := fl.Field().String()
        // Check if name contains only alphabetic characters
        return IsAlpha(name)
    })
    validate.RegisterValidation("isAlphaNumeric", func(fl validator.FieldLevel) bool {
        value := fl.Field().String()
        // Check if value contains only alphanumeric characters
        return IsAlphaNumeric(value)
    })

    // Custom validation rules for each field
    if err := validate.Var(req.Name, "required,isAlpha"); err != nil {
        return constants.ErrInvalidName
    }
    if err := validate.Var(req.Email, "required,email"); err != nil {
        return constants.ErrInvalidEmail
    }
    if err := validate.Var(req.PhoneNumber, "required,len=10"); err != nil {
        return constants.ErrInvalidPhoneNumber
    }
    if err := validate.Var(req.PancardNumber, "required,isAlphaNumeric"); err != nil {
        return constants.ErrInvalidPancardNumber
    }
    if err := validate.Var(req.Password, "required"); err != nil {
        return constants.ErrMissingPassword
    }

    return nil
}

// IsAlphaNumeric checks if a string contains only alphanumeric characters
func IsAlphaNumeric(str string) bool {
    for _, char := range str {
        if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < '0' || char > '9') {
            return false
        }
    }
    return true
}

// IsAlpha checks if a string contains only alphabetic characters
func IsAlpha(str string) bool {
    for _, char := range str {
        if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
            return false
        }
    }
    return true
}
