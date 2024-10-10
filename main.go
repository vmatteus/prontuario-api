package main

import (
	"prontuario/internal/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", auth.Login)
	r.Run(":8080")
}
