package auth

import (
	"prontuario/internal/auth/application"
	"prontuario/internal/auth/application/interfaces"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service interfaces.AuthServiceInterface
}

func NewAuthController() *AuthController {
	return &AuthController{Service: application.NewAuthService()}
}

func (controller *AuthController) Login(ctx *gin.Context) {
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
