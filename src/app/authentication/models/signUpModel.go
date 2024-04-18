package models

// SignUpRequest represents the data structure for a signup request
type SignUpRequest struct {
	Name          string `json:"name" validate:"required,alpha" example:"John"`
	Email         string `json:"email" validate:"required,email" example:"john@example.com"`
	PhoneNumber   int `json:"phoneNumber" validate:"required,len=10" example:"1234567890"`
	PancardNumber string `json:"pancardNumber" validate:"required,alphanum"  example:"ABCDE1234F"`
	Password      string `json:"password" validate:"required" example:"password123"`
}

//  name of the database table for the SignUpRequest model
func (SignUpRequest) TableName() string {
	return "users"
}
