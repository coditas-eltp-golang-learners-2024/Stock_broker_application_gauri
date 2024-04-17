package models

import "time"

type OTPRequest struct {
	Email           string    `json:"email"`
	OTP             string    `json:"otp"`
	OTPCreationTime time.Time `json:"otp_created_at"`
}

func (OTPRequest) TableName() string {
	return "users"
}
