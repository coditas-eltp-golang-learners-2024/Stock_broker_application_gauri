package repo

import (
	"Stock_broker_application/constants"
	"Stock_broker_application/models"
	"database/sql"
)

// UserRepositoryImpl is the implementation of UserRepository
type UserRepositoryImpl struct {
	db *sql.DB
}

// NewUserRepositoryImpl creates a new instance of UserRepositoryImpl
func NewUserRepositoryImpl(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

// IsEmailExists checks if the email already exists in the database
func (r *UserRepositoryImpl) IsEmailExists(email string) bool {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
	if err != nil {
		// Handle error
		return false
	}
	return count > 0
}

// IsPhoneNumberExists checks if the phone number already exists in the database
func (r *UserRepositoryImpl) IsPhoneNumberExists(phoneNumber string) bool {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM users WHERE phone_number = ?", phoneNumber).Scan(&count)
	if err != nil {
		// Handle error
		return false
	}
	return count > 0
}

// IsPancardNumberExists checks if the pancard number already exists in the database
func (r *UserRepositoryImpl) IsPancardNumberExists(pancardNumber string) bool {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM users WHERE pancard_number = ?", pancardNumber).Scan(&count)
	if err != nil {
		// Handle error
		return false
	}
	return count > 0
}

// InsertUser inserts a new user into the database
func (r *UserRepositoryImpl) InsertUser(user models.SignUpRequest) error {
	_, err := r.db.Exec("INSERT INTO users (name, email, phone_number, pancard_number, password) VALUES (?, ?, ?, ?, ?)",
		user.Name, user.Email, user.PhoneNumber, user.PancardNumber, user.Password)
	if err != nil {
		// Handle error
		return constants.ErrDatabaseInsert
	}
	return nil
}
func (r *UserRepositoryImpl) GetUserByEmail(email string) *models.SignInRequest {
	// Declare a variable to hold the retrieved user
	var user models.SignInRequest

	// Query the database for a user with the provided email
	err := r.db.QueryRow("SELECT email,password FROM users WHERE email=?", email).Scan(&user.Email, &user.Password)
	if err != nil {
		// If no rows are found, return nil indicating no user was found for the provided email
		if err == sql.ErrNoRows {
			return nil
		}
		// Handle other database errors if any
		return nil
	}

	// Return a pointer to the retrieved user
	return &user
}
