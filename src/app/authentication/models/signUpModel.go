package models
// SignUpRequest represents the data structure for a signup request
type SignUpRequest struct {
	Name          string `json:"name" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	PhoneNumber   string `json:"phoneNumber" validate:"required"`
	PancardNumber string `json:"pancardNumber" validate:"required"`
	Password      string `json:"password" validate:"required"`
}