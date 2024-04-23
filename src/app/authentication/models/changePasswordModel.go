package models

type ChangePassword struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}

func (ChangePassword) TableName() string {
	return "users"
}
