// RenewMembresiaController.go
package infrastructure

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/application"
)

type RenewMembresiaController struct {
	renewUseCase *application.RenewMembresiaUseCase
}

func NewRenewMembresiaController(renewUseCase *application.RenewMembresiaUseCase) *RenewMembresiaController {
	return &RenewMembresiaController{
		renewUseCase: renewUseCase,
	}
}

func (ctrl *RenewMembresiaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var request struct {
		NewEndDate time.Time `json:"newEndDate" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fecha de fin inválida",
			"error":   err.Error(),
		})
		return
	}

	err = ctrl.renewUseCase.Run(int32(id), request.NewEndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo renovar la membresía",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Membresía renovada exitosamente",
	})
}
