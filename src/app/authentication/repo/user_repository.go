package repo

import "Stock_broker_application/models"

// UserRepository defines methods for interacting with user data in the database
type UserRepository interface {
    // IsEmailExists checks if the email already exists in the database
    // @param email The email to check for existence
    // @return bool Returns true if the email exists, false otherwise
    IsEmailExists(email string) bool

    // IsPhoneNumberExists checks if the phone number already exists in the database
    // @param phoneNumber The phone number to check for existence
    // @return bool Returns true if the phone number exists, false otherwise
    IsPhoneNumberExists(phoneNumber string) bool

    // IsPancardNumberExists checks if the pancard number already exists in the database
    // @param pancardNumber The pancard number to check for existence
    // @return bool Returns true if the pancard number exists, false otherwise
    IsPancardNumberExists(pancardNumber string) bool

    // InsertUser inserts a new user into the database
    // @param user The user details to insert into the database
    // @return error Returns an error if the insertion fails, nil otherwise
    InsertUser(user models.SignUpRequest) error

    // GetUserByEmail retrieves a user from the database by email
    // @param email The email of the user to retrieve
    // @return *models.SignInRequest Returns the user information if found, nil otherwise
    GetUserByEmail(email string) *models.SignInRequest
}
