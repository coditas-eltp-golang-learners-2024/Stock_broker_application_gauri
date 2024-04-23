package repo

import (
	"Stock_broker_application/constants"
	"Stock_broker_application/models"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

// UserRepository defines methods for interacting with user data in the database
type UserRepository interface {
	IsEmailExists(email string) bool
	IsPhoneNumberExists(phoneNumber uint64) bool
	IsPancardNumberExists(pancardNumber string) bool
	InsertUser(user models.SignUpRequest) error
	GetUserByEmail(email string) *models.SignInRequest
	SaveOTP(email string, newOTP int) error
	GetOTPByEmail(email string) (int, time.Time, error)
	UpdatePassword(email string, newPassword string) error
	CheckOldPassword(email string, password string) bool
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

func (repo *UserRepositoryImpl) IsPhoneNumberExists(phoneNumber uint64) bool {
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
func (repo *UserRepositoryImpl) SaveOTP(email string, newOTP int) error {

	// Update the OTP column for the given email
	if err := repo.db.Model(&models.OTPRequest{}).Where("email = ?", email).Update("otp", newOTP).Error; err != nil {
		return err
	}
	// Generate the new OTP creation time
	otpCreationTime := time.Now()
	otpCreationTime = otpCreationTime.Truncate(time.Second)
	// Update the OTP creation time column for the given email
	if err := repo.db.Model(&models.OTPRequest{}).Where("email = ?", email).Update("otp_creation_time", otpCreationTime).Error; err != nil {
		return err
	}

	return nil
}

func (repo *UserRepositoryImpl) GetOTPByEmail(email string) (int, time.Time, error) {
	var otpData struct {
		OTP             int            `db:"otp"`
		OTPCreationTime mysql.NullTime `db:"otp_creation_time"`
	}

	err := repo.db.Table("users").
		Where("email = ?", email).
		Select("otp, otp_creation_time").
		Scan(&otpData).
		Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, time.Time{}, constants.ErrUserNotFound
		}
		return 0, time.Time{}, err
	}

	otpCreationTime := otpData.OTPCreationTime.Time

	return otpData.OTP, otpCreationTime, nil
}
func (repo *UserRepositoryImpl) UpdatePassword(email string, newPassword string) error {
	if err := repo.db.Model(&models.ChangePassword{}).Where("email = ?", email).Update("password", newPassword).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepositoryImpl) CheckOldPassword(email string, password string) bool {
	var count int64
	repo.db.Model(&models.ChangePassword{}).Where("email = ? AND password = ?", email, password).Count(&count)
	return count > 0
}
