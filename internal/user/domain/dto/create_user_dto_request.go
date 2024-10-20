package dto

import "github.com/go-playground/validator/v10"

type CreateUserDtoRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func (req CreateUserDtoRequest) Validate() error {
	validator := validator.New()
	return validator.Struct(req)
}
