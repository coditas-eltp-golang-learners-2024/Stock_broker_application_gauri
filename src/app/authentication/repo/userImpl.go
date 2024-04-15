package repo

import (
	"Stock_broker_application/constants"
	"Stock_broker_application/models"

	"gorm.io/gorm"
)
// UserRepository defines methods for interacting with user data in the database
type UserRepository interface {
	IsEmailExists(email string) bool
	IsPhoneNumberExists(phoneNumber string) bool
	IsPancardNumberExists(pancardNumber string) bool
	InsertUser(user models.SignUpRequest) error
	GetUserByEmail(email string) *models.SignInRequest
}
// UserRepositoryImpl is the implementation of UserRepository
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepositoryImpl creates a new instance of UserRepositoryImpl
func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (repo *UserRepositoryImpl) IsEmailExists(email string) bool {
	var count int64
	repo.db.Model(&models.SignUpRequest{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func (repo *UserRepositoryImpl) IsPhoneNumberExists(phoneNumber string) bool {
	var count int64
	repo.db.Model(&models.SignUpRequest{}).Where("phone_number = ?", phoneNumber).Count(&count)
	return count > 0
}

func (repo *UserRepositoryImpl) IsPancardNumberExists(pancardNumber string) bool {
	var count int64
	repo.db.Model(&models.SignUpRequest{}).Where("pancard_number = ?", pancardNumber).Count(&count)
	return count > 0
}

func (repo *UserRepositoryImpl) InsertUser(user models.SignUpRequest) error {
	err := repo.db.Create(&models.SignUpRequest{
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

func (repo *UserRepositoryImpl) GetUserByEmail(email string) *models.SignInRequest {
	var user models.SignInRequest
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}

		return nil
	}
	return &user
}
