package auth

import (
	"log"
	"net/http"
	"prontuario/internal/auth/application"
	"prontuario/internal/auth/application/interfaces"
	"prontuario/internal/auth/domain/dto"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service interfaces.AuthServiceInterface
}

type ResposeLogin struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"PhoneNumber"`
	Token       string `json:"token"`
}

func NewAuthController() *AuthController {
	return &AuthController{Service: application.NewAuthService()}
}

func (controller *AuthController) Login(ctx *gin.Context) {

	var request *dto.LoginDtoRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := controller.Service.Login(ctx, request.Email, request.Password)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "login error"})
		return
	}

	ResposeLogin := ResposeLogin{
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Token:       token,
	}

	ctx.JSON(http.StatusOK, gin.H{"data": ResposeLogin})
}

func (controller *AuthController) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
			ctx.Abort()
			return
		}

		_, err := controller.Service.ValidateToken(token)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token invalid"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
