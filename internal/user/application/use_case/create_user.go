package usecase

import (
	"context"
	auth "prontuario/internal/auth/application"
	"prontuario/internal/user/application"
	"prontuario/internal/user/domain"
	"prontuario/internal/user/domain/dto"
	"time"
)

type CreateUserUseCase struct{}

func (useCase CreateUserUseCase) Execute(createUserDtoRequest dto.CreateUserDtoRequest) (*domain.UserModel, error) {

	now := time.Now()
	userService := application.NewUserService()
	authService := auth.NewAuthService()
	passwordCrypted, err := authService.HashPassword(createUserDtoRequest.Password)

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

	userDb, err := userService.Create(context.TODO(), user)

	if err != nil {
		return nil, err
	}

	return userDb, nil
}
