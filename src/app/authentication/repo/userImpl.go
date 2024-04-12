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

func (r *UserRepositoryImpl) IsEmailExists(email string) bool {
	var count int64
	r.db.Model(&models.SignUpRequest{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func (r *UserRepositoryImpl) IsPhoneNumberExists(phoneNumber string) bool {
	var count int64
	r.db.Model(&models.SignUpRequest{}).Where("phone_number = ?", phoneNumber).Count(&count)
	return count > 0
}

func (r *UserRepositoryImpl) IsPancardNumberExists(pancardNumber string) bool {
	var count int64
	r.db.Model(&models.SignUpRequest{}).Where("pancard_number = ?", pancardNumber).Count(&count)
	return count > 0
}

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

func (r *UserRepositoryImpl) GetUserByEmail(email string) *models.SignInRequest {
	var user models.SignInRequest
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}

		return nil
	}
	return &user
}
