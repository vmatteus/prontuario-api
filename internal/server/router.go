package server

import (
	authController "prontuario/internal/auth/presentation"
	userController "prontuario/internal/user/presentation"

	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	r := gin.Default()

	authController := authController.NewAuthController()
	userController := userController.NewUserController()

	r.POST("/login", authController.Login)
	r.POST("/register", userController.Create)
	return r
}
