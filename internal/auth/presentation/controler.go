package auth

import (
	"net/http"
	"prontuario/internal/auth/application"
	"prontuario/internal/auth/application/interfaces"
	"prontuario/internal/auth/domain/dto"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service interfaces.AuthServiceInterface
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

	_, err := controller.Service.Login(ctx, request.Email, request.Password)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "login error"})
	}

	// //var input LoginInput

	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Campos inválidos"})
	// 	return
	// }

	// if input.Username == "admin" && input.Password == "password" {
	// 	c.JSON(http.StatusOK, gin.H{"message": "Login realizado com sucesso"})
	// } else {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
	// }
}
