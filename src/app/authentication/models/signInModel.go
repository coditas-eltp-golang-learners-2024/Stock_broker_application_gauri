package models
// SignUpRequest represents the data structure for a signup request
type SignInRequest struct {
	Email         string `json:"email" validate:"required,email"`
	Password      string `json:"password" validate:"required"`
}

func (SignInRequest) TableName() string {
    return "users"
}