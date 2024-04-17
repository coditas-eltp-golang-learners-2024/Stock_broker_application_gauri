package repo

import (
	"Stock_broker_application/constants"
	"Stock_broker_application/models"
	"time"

	"gorm.io/gorm"
)

// UserRepository defines methods for interacting with user data in the database
type UserRepository interface {
	IsEmailExists(email string) bool
	IsPhoneNumberExists(phoneNumber string) bool
	IsPancardNumberExists(pancardNumber string) bool
	InsertUser(user models.SignUpRequest) error
	GetUserByEmail(email string) *models.SignInRequest
	SaveOTP(email string, newOTP string) error
	GetOTPByEmail(email string) (string, error)
	//UpdateOTP(email string, newOTP string) error
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
func (repo *UserRepositoryImpl) SaveOTP(email string, newOTP string) error {
	// Generate the new OTP creation time
	otpCreationTime := time.Now().Add(time.Minute * 1)
	otpCreationTime = otpCreationTime.Truncate(time.Second)

	// Update the OTP column for the given email
	if err := repo.db.Model(&models.OTPRequest{}).Where("email = ?", email).Update("otp", newOTP).Error; err != nil {
		return err
	}

	// Update the OTP creation time column for the given email
	if err := repo.db.Model(&models.OTPRequest{}).Where("email = ?", email).Update("otp_creation_time", otpCreationTime).Error; err != nil {
		return err
	}

	return nil
}

func (repo *UserRepositoryImpl) GetOTPByEmail(email string) (string, error) {
	var otp string
	if err := repo.db.Table("users").Where("email = ?", email).Select("otp").Scan(&otp).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", constants.ErrUserNotFound
		}
		return "", err
	}

	return otp, nil
}
// func (repo *UserRepositoryImpl) GetOTPByEmail(email string) (string, time.Time, error) {
//     var otp models.OTPRequest
//     if err := repo.db.Table("users").
//         Select("otp, otp_creation_time").
//         Where("email = ?", email).
//         First(&otp).Error; err != nil {
//         if err == gorm.ErrRecordNotFound {
//             return "", time.Time{}, constants.ErrUserNotFound
//         }
//         return "", time.Time{}, err
//     }

//     return otp.OTP, otp.OTPCreationTime, nil
// }
// // UpdateOTP updates the OTP for the user with the given email
// func (repo *UserRepositoryImpl) UpdateOTP(email string, newOTP string) error {
//     if err := repo.db.Model(&models.OTPRequest{}).Where("email = ?", email).Update("otp", newOTP).Error; err != nil {
//         return err
//     }
//     return nil
// }
