package server

import (
	auth "prontuario/internal/auth/presentation"

	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/login", auth.Login)
	return r
}
