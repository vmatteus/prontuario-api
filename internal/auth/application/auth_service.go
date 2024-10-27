package application

import (
	"context"
	"errors"
	"prontuario/configs"
	"prontuario/internal/user/domain"
	"prontuario/internal/user/infrastructure"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo domain.UserRepositoryInterface
}

func NewAuthService() *AuthService {
	return &AuthService{repo: infrastructure.GetUserRepository()}
}

func (service *AuthService) Login(ctx context.Context, username, password string) (*domain.UserModel, string, error) {
	user, err := service.repo.FindUserByEmail(ctx, username)

	if err != nil {
		return nil, "", err
	}

	comparasion := service.CompareHashAndPassword(user.Password, password)

	token, errToken := service.createToken(username)

	if errToken != nil {
		return nil, "", errors.New("error creating token")
	}

	if comparasion == nil {
		return user, token, nil
	}

	return nil, "", errors.New("invalid password")
}

func (AuthService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (AuthService) CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (AuthService) ValidateToken(token string) (*domain.UserModel, error) {

	config := configs.Get()

	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Jwt.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, errors.New("invalid token")
	}

	return nil, nil
}

func (AuthService) createToken(username string) (string, error) {

	config := configs.Get()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(config.Jwt.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
