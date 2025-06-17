package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	baseController "github.com/txzy2/simple-api/internal/http"
	"github.com/txzy2/simple-api/internal/services/user"
)

type UserController struct {
	*baseController.Controller
	UserService user.IUserService
}

func NewUserController(userService user.IUserService) *UserController {
	return &UserController{
		Controller:  &baseController.Controller{},
		UserService: userService,
	}
}

func (u *UserController) GetUserById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	input := user.UserInput{Id: idParam}
	userOutput, err := u.UserService.GetUserByID(input)

	if err != nil {
		u.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	u.SuccessResponse(ctx, "Пользователь успешно получен", map[string]any{
		"id":   userOutput.Id,
		"name": userOutput.Name,
		"age":  userOutput.Age,
	})
}
