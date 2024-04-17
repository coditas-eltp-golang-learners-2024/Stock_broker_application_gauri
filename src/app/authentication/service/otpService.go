package service

import (
	"Stock_broker_application/repo"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"
	// "time"
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
func (s *OTPService) GenerateAndSaveOTP(email string) error {
    // Generate random OTP
    otp, err := generateRandomOTP()
    if err != nil {
        return err
    }

  
	
    if err := s.UserRepository.SaveOTP(email, otp); err != nil {
        return err
    }

    return nil
}

// generateRandomOTP generates a random 6-digit OTP
func generateRandomOTP() (string, error) {
    // Generate 3 bytes of random data
    bytes := make([]byte, 3)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }

    // Encode the bytes to base64
    otp := base64.StdEncoding.EncodeToString(bytes)

    // Strip any special characters from the OTP
    otp = otp[:4]

    return otp, nil
}


// ValidateOTP validates the provided OTP for the given email and checks if it has expired after 5 minutes.
func (s *OTPService) ValidateOTP(email, providedOTP string) error {
	// Retrieve the OTP and its creation time from the database based on the provided email
	savedOTP, otpCreationTime, err := s.UserRepository.GetOTPByEmail(email)
	if err != nil {
		return err
	}

	// Check if the provided OTP matches the saved OTP
	if providedOTP != savedOTP {
		return errors.New("invalid OTP")
	}

	// Check if the OTP has expired (e.g., 5 minutes after creation time)
	if time.Since(otpCreationTime) > 5*time.Minute {
		return errors.New("OTP has expired")
	}

	return nil
}