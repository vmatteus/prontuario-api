package dto

import "github.com/go-playground/validator/v10"

type CreateUserDtoRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (req CreateUserDtoRequest) Validate() error {
	validator := validator.New()
	return validator.Struct(req)
}
