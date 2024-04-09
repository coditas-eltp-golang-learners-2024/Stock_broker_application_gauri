package constants

import "errors"

var (
    //db errors
    ErrDatabaseConnection = errors.New("failed to connect to database")
    ErrDatabasePing       = errors.New("failed to ping database")
    ErrDatabaseQuery     = errors.New("database query failed")
    ErrDatabaseInsert    = errors.New("failed to insert into database")

    //duplicte entry in db
    ErrEmailExists       = errors.New("email already exists")
    ErrPhoneNumberExists = errors.New("phone number already exists")
    ErrPancardExists     = errors.New("pancard number already exists")

    //validation error
    ErrInvalidName          = errors.New("name should contain only alphabetic characters")
    ErrInvalidEmail         = errors.New("invalid email address")
    ErrInvalidPhoneNumber   = errors.New("phone number should be 10 digits")
    ErrInvalidPancardNumber = errors.New("pancard number should contain only alphanumeric characters")
    ErrMissingPassword      = errors.New("password is required")

    ErrUserNotFound         = errors.New("User not found")
    ErrInvalidCredentials   =errors.New("Invalid credentals")
)