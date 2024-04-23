package service

import (
	"Stock_broker_application/repo"
	"errors"
	"math/rand"
	"time"
)

type OTPService struct {
	userRepository repo.UserRepository
}

func NewOTPService(userRepository repo.UserRepository) *OTPService {
	return &OTPService{
		userRepository: userRepository,
	}
}

// GenerateAndSaveOTP generates OTP, swith its creation time in the db

func (otpservice *OTPService) GenerateAndSaveOTP(email string) error {
	// Generate random OTP
	otp := rand.Intn(8999) + 1000

	if err := otpservice.userRepository.SaveOTP(email, otp); err != nil {
		return err
	}

	return nil
}

// ValidateOTP validates the provided OTP 
func (otpservice *OTPService) ValidateOTP(email string, providedOTP int) error {
	
	savedOTP, otpCreationTime, err := otpservice.userRepository.GetOTPByEmail(email)
	if err != nil {
		return err
	}

	if providedOTP != savedOTP {
		return errors.New("invalid OTP")
	}

	if time.Since(otpCreationTime) > 10*time.Minute {
		return errors.New("OTP has expired")
	}

	return nil
}
