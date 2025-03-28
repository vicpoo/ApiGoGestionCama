// GetAllMembresiasController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/application"
)

type GetAllMembresiasController struct {
	getAllUseCase *application.GetAllMembresiasUseCase
}

func NewGetAllMembresiasController(getAllUseCase *application.GetAllMembresiasUseCase) *GetAllMembresiasController {
	return &GetAllMembresiasController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllMembresiasController) Run(c *gin.Context) {
	membresias, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las membresías",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, membresias)
}
