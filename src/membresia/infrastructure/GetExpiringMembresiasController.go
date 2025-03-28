// GetExpiringMembresiasController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/application"
)

type GetExpiringMembresiasController struct {
	getExpiringUseCase *application.GetExpiringMembresiasUseCase
}

func NewGetExpiringMembresiasController(getExpiringUseCase *application.GetExpiringMembresiasUseCase) *GetExpiringMembresiasController {
	return &GetExpiringMembresiasController{
		getExpiringUseCase: getExpiringUseCase,
	}
}

func (ctrl *GetExpiringMembresiasController) Run(c *gin.Context) {
	daysParam := c.DefaultQuery("days", "7")
	days, err := strconv.Atoi(daysParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Parámetro 'days' inválido",
			"error":   err.Error(),
		})
		return
	}

	membresias, err := ctrl.getExpiringUseCase.Run(days)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las membresías por expirar",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, membresias)
}
