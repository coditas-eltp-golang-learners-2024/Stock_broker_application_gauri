package repo

import "Stock_broker_application/models"

// UserRepository defines methods for interacting with user data in the database
type UserRepository interface {
    IsEmailExists(email string) bool
    IsPhoneNumberExists(phoneNumber string) bool
    IsPancardNumberExists(pancardNumber string) bool
    InsertUser(user models.SignUpRequest) error
}
