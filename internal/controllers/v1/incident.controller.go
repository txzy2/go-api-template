package v1

import (
	"github.com/gin-gonic/gin"
	baseController "github.com/txzy2/simple-api/internal/controllers"
	"github.com/txzy2/simple-api/internal/services"
)

type (
	IncidentController struct {
		*baseController.Controller
		services *services.Provider
	}

	IIncidentController interface {
		New(ctx *gin.Context)
	}
)

func NewIncidentController(services *services.Provider) *IncidentController {
	return &IncidentController{
		Controller: &baseController.Controller{},
		services:   services,
	}
}

func (inc *IncidentController) New(ctx *gin.Context) {

}
