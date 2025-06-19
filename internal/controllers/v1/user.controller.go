package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	baseController "github.com/txzy2/simple-api/internal/controllers"
	"github.com/txzy2/simple-api/internal/services"
)

const CLASS_NAME string = "USER_CONTROLLER"

type (
	UserController struct {
		*baseController.Controller
		services *services.Provider
	}

	IUserController interface {
		GetUserById(ctx *gin.Context)
		CreateNewUser(ctx *gin.Context)
	}
)

func NewUserController(services *services.Provider) *UserController {
	return &UserController{
		Controller: &baseController.Controller{},
		services:   services,
	}
}

func (u *UserController) GetUserById(ctx *gin.Context) {
	fmt.Printf("[DEBUG] [%s] GetUserById REQUEST: id=%s\n", CLASS_NAME, ctx.Param("id"))
	idParam := ctx.Param("id")

	input := services.UserInput{Id: idParam}
	userOutput, err := u.services.UserService.GetUserByID(input)

	if err != nil {
		u.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	u.SuccessResponse(ctx, "Пользователь успешно получен", userOutput)
}

func (u *UserController) CreateNewUser(ctx *gin.Context) {
	var input services.CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := u.services.UserService.CreateUser(input)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		u.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	u.SuccessResponse(ctx, "User is create", gin.H{"id": id})
}
