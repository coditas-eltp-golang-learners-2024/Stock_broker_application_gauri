package repo
import (
	"Stock_broker_application/constants"
	"Stock_broker_application/models"

	"gorm.io/gorm"
)
// UserRepositoryImpl is the implementation of UserRepository
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepositoryImpl creates a new instance of UserRepositoryImpl
func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

// IsEmailExists checks if the email already exists in the database
// @Summary Check if the email exists
// @Description Checks if the provided email already exists in the database
// @Param email query string true "Email address to check"
// @Success 200 {string} string "true if email exists, false otherwise"

func (r *UserRepositoryImpl) IsEmailExists(email string) bool {
	var count int64
	r.db.Model(&models.SignUpRequest{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneNumberExists checks if the phone number already exists in the database
// @Summary Check if the phone number exists
// @Description Checks if the provided phone number already exists in the database
// @Param phoneNumber query string true "Phone number to check"
// @Success 200 {string} string "true if phone number exists, false otherwise"

func (r *UserRepositoryImpl) IsPhoneNumberExists(phoneNumber string) bool {
	var count int64
	r.db.Model(&models.SignUpRequest{}).Where("phone_number = ?", phoneNumber).Count(&count)
	return count > 0
}

// IsPancardNumberExists checks if the pancard number already exists in the database
// @Summary Check if the pancard number exists
// @Description Checks if the provided pancard number already exists in the database
// @Param pancardNumber query string true "Pancard number to check"
// @Success 200 {string} string "true if pancard number exists, false otherwise"

func (r *UserRepositoryImpl) IsPancardNumberExists(pancardNumber string) bool {
	var count int64
	r.db.Model(&models.SignUpRequest{}).Where("pancard_number = ?", pancardNumber).Count(&count)
	return count > 0
}


// InsertUser inserts a new user into the database
// @Summary Add a new user
// @Description Adds a new user to the database
// @Param user body models.SignUpRequest true "User object to insert"
// @Success 200 {string} string "User inserted successfully"
// @Failure 400 {string} string "Invalid input"

func (r *UserRepositoryImpl) InsertUser(user models.SignUpRequest) error {
	err := r.db.Create(&models.SignUpRequest{
		Name:          user.Name,
		Email:         user.Email,
		PhoneNumber:   user.PhoneNumber,
		PancardNumber: user.PancardNumber,
		Password:      user.Password,
	}).Error

	if err != nil {
		return constants.ErrDatabaseInsert
	}
	return nil
}


// GetUserByEmail retrieves a user by email from the database
// @Summary Get user by email
// @Description Retrieves a user by their email address
// @Param email query string true "Email address of the user"
// @Success 200 {object} models.SignInRequest "User object"
// @Failure 404 {string} string "User not found"

func (r *UserRepositoryImpl) GetUserByEmail(email string) *models.SignInRequest {
	var user models.SignInRequest
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		// Handle other database errors if any
		return nil
	}
	return &user
}
