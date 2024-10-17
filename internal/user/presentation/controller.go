package presentation

import (
	"net/http"
	"prontuario/internal/user/application"
	"prontuario/internal/user/application/interfaces"
	"prontuario/internal/user/domain"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service interfaces.UserServiceInterface
}

func NewUserController() *UserController {
	return &UserController{Service: application.NewUserService()}
}

func (controller *UserController) Create(ctx *gin.Context) {

	var request *domain.UserModel

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := controller.Service.Create(ctx, request)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user_id": user.ID,
	})
}
