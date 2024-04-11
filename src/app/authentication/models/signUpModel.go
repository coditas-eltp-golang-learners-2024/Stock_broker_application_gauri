package models

// SignUpRequest represents the data structure for a signup request
type SignUpRequest struct {
	Name          string `json:"name" validate:"required=Name is required,alpha=Name should contain only alphabetic characters"`
	Email         string `json:"email" validate:"required,email"`
	PhoneNumber   string `json:"phoneNumber" validate:"required,len=10"`
	PancardNumber string `json:"pancardNumber" validate:"required,alphanum"`
	Password      string `json:"password" validate:"required"`
}
func (SignUpRequest) TableName() string {
    return "users"
}