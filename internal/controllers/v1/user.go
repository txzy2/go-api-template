package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	baseController "github.com/txzy2/simple-api/internal/controllers"
	"github.com/txzy2/simple-api/internal/services"
)

type UserController struct {
	*baseController.Controller
	services *services.Provider
}

func NewUserController(services *services.Provider) *UserController {
	return &UserController{
		Controller: &baseController.Controller{},
		services:   services,
	}
}

func (u *UserController) GetUserById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	input := services.UserInput{Id: idParam}
	userOutput, err := u.services.UserService.GetUserByID(input)

	if err != nil {
		u.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	u.SuccessResponse(ctx, "Пользователь успешно получен", userOutput)
}
