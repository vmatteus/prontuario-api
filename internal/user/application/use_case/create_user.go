package usecase

import (
	"context"
	"prontuario/internal/user/application"
	"prontuario/internal/user/domain"
	"prontuario/internal/user/domain/dto"
	"time"
)

type CreateUserUseCase struct{}

func (CreateUserUseCase) Execute(createUserDtoRequest dto.CreateUserDtoRequest) (*domain.UserModel, error) {

	now := time.Now()

	user := &domain.UserModel{
		Name:      createUserDtoRequest.Name,
		Username:  createUserDtoRequest.Username,
		Password:  createUserDtoRequest.Password,
		CreatedAt: &now,
		UpdatedAt: &now,
		DeletedAt: nil,
	}

	userService := application.NewUserService()

	user, err := userService.Create(context.TODO(), user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
