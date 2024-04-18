package models

import "time"

type OTPRequest struct {
	Email           string    `json:"email"  validate:"required,email"`
	OTP             int    `json:"otp" validate:"required"`
	OTPCreationTime time.Time `json:"otp_created_at"`
}

func (OTPRequest) TableName() string {
	return "users"
}
