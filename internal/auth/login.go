package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos inválidos"})
		return
	}

	if input.Username == "admin" && input.Password == "password" {
		c.JSON(http.StatusOK, gin.H{"message": "Login realizado com sucesso"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
	}
}
