package presentation

import (
	"net/http"
	"prontuario/internal/user/application"
	"prontuario/internal/user/application/interfaces"
	usecase "prontuario/internal/user/application/use_case"
	"prontuario/internal/user/domain/dto"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service interfaces.UserServiceInterface
}

func NewUserController() *UserController {
	return &UserController{Service: application.NewUserService()}
}

func (controller *UserController) Create(ctx *gin.Context) {

	var request *dto.CreateUserDtoRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := usecase.CreateUserUseCase{}.Execute(*request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user_id": user.ID,
	})
}
