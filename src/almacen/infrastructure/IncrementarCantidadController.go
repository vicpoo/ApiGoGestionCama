// IncrementarCantidadController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/application"
)

type IncrementarCantidadController struct {
	incrementarUseCase *application.IncrementarCantidadUseCase
}

func NewIncrementarCantidadController(incrementarUseCase *application.IncrementarCantidadUseCase) *IncrementarCantidadController {
	return &IncrementarCantidadController{
		incrementarUseCase: incrementarUseCase,
	}
}

func (ctrl *IncrementarCantidadController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var incrementoRequest struct {
		Cantidad int32 `json:"cantidad"`
	}

	if err := c.ShouldBindJSON(&incrementoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	err = ctrl.incrementarUseCase.Run(int32(id), incrementoRequest.Cantidad)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo incrementar la cantidad",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Cantidad incrementada exitosamente",
	})
}
