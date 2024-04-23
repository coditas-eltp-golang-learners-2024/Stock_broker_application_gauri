package service

import (
	"Stock_broker_application/repo"
	"errors"
	"math/rand"
	"time"
)

type OTPService struct {
	UserRepository repo.UserRepository
}

func NewOTPService(userRepository repo.UserRepository) *OTPService {
	return &OTPService{
		UserRepository: userRepository,
	}
}

// GenerateAndSaveOTP generates a random OTP, saves it along with its creation time in the database
// for the given email address
func (otpservice *OTPService) GenerateAndSaveOTP(email string) error {
	// Generate random OTP
	otp := rand.Intn(8999) + 1000

	if err := otpservice.UserRepository.SaveOTP(email, otp); err != nil {
		return err
	}

	return nil
}

// ValidateOTP validates the provided OTP for the given email and checks if it has expired after 5 minutes.
func (otpservice *OTPService) ValidateOTP(email string, providedOTP int) error {
	// Retrieve the OTP and its creation time from the database based on the provided email
	savedOTP, otpCreationTime, err := otpservice.UserRepository.GetOTPByEmail(email)
	if err != nil {
		return err
	}

	// Check if the provided OTP matches the saved OTP
	if providedOTP != savedOTP {
		return errors.New("invalid OTP")
	}

	// Check if the OTP has expired (e.g., 5 minutes after creation time)
	if time.Since(otpCreationTime) > 24*time.Hour {
		return errors.New("OTP has expired")
	}

	return nil
}
