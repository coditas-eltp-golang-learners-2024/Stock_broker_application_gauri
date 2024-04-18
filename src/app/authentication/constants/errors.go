package constants

import "errors"

var (
	// Database errors
	ErrDatabaseConnection = errors.New("failed to connect to database")
	ErrDatabasePing       = errors.New("failed to ping database")
	ErrDatabaseQuery      = errors.New("database query failed")
	ErrDatabaseInsert     = errors.New("failed to insert into database")

	// Duplicate entry in database
	ErrEmailExists       = errors.New("email already exists")
	ErrPhoneNumberExists = errors.New("phone number already exists")
	ErrPancardExists     = errors.New("pancard number already exists")

	// Validation errors
	ErrInvalidName          = errors.New("name should contain only alphabetic characters")
	ErrInvalidEmail         = errors.New("invalid email address")
	ErrInvalidPhoneNumber   = errors.New("phone number should be 10 digits")
	ErrInvalidPancardNumber = errors.New("pancard number should contain only alphanumeric characters")
	ErrMissingPassword      = errors.New("password is required")

	// Customized error messages for empty fields
	ErrEmptyName          = errors.New("name field cannot be empty")
	ErrEmptyEmail         = errors.New("email field cannot be empty")
	ErrEmptyPhoneNumber   = errors.New("phone number should jhave 10 digits")
	ErrEmptyPancardNumber = errors.New("pancard number field cannot be empty")
	ErrEmptyPassword      = errors.New("password field cannot be empty")

	// Other errors
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmptyField         = errors.New("all required fields should be present")
	ErrValidationFailed   = errors.New("validation failed")
)
