package server

import (
	authController "prontuario/internal/auth/presentation"
	userController "prontuario/internal/user/presentation"

	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	r := gin.Default()

	authController := authController.NewAuthController()
	r.POST("/login", authController.Login)

	userController := userController.NewUserController()
	r.POST("/register", userController.Create)

	authorized := r.Group("/api")
	authorized.Use(authController.AuthMiddleware())
	{
		authorized.GET("/user/:id", userController.Get)
	}

	return r
}
