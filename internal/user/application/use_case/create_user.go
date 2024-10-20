package usecase

import (
	"context"
	"prontuario/internal/crypt"
	"prontuario/internal/user/application"
	"prontuario/internal/user/domain"
	"prontuario/internal/user/domain/dto"
	"time"
)

type CreateUserUseCase struct{}

func (CreateUserUseCase) Execute(createUserDtoRequest dto.CreateUserDtoRequest) (*domain.UserModel, error) {

	now := time.Now()
	crypt := crypt.NewCrypt()
	passwordCrypted, err := crypt.HashPassword(createUserDtoRequest.Password)

	if err != nil {
		return nil, err
	}

	user := &domain.UserModel{
		Name:        createUserDtoRequest.Name,
		Email:       createUserDtoRequest.Email,
		PhoneNumber: createUserDtoRequest.PhoneNumber,
		Password:    passwordCrypted,
		CreatedAt:   &now,
		UpdatedAt:   &now,
		DeletedAt:   nil,
	}

	userService := application.NewUserService()

	userDb, err := userService.Create(context.TODO(), user)

	if err != nil {
		return nil, err
	}

	return userDb, nil
}
