package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/application"
)

type GetAllCamasController struct {
	getAllUseCase *application.GetAllCamasUseCase
}

func NewGetAllCamasController(getAllUseCase *application.GetAllCamasUseCase) *GetAllCamasController {
	return &GetAllCamasController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllCamasController) Run(c *gin.Context) {
	camas, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las camas",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, camas)
}
