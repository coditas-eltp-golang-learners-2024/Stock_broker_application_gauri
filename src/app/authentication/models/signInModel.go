package models

// SignInRequest represents the data structure for a sign-in request
// swagger:model
type SignInRequest struct {
	Email    string `json:"email" validate:"required,email" example:"johndoe@.com"`
	Password string `json:"password" validate:"required" example:"password123"`
}

//  name of the database table for the SignInRequest model
func (SignInRequest) TableName() string {
	return "users"
}
