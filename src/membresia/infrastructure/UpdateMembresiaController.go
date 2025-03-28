// UpdateMembresiaController.go
package infrastructure

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/application"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain/entities"
)

type UpdateMembresiaController struct {
	updateUseCase *application.UpdateMembresiaUseCase
}

func NewUpdateMembresiaController(updateUseCase *application.UpdateMembresiaUseCase) *UpdateMembresiaController {
	return &UpdateMembresiaController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateMembresiaController) Run(c *gin.Context) {
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
		UsuarioID   int32  `json:"usuarioId"`
		FechaInicio string `json:"fechaInicio"`
		FechaFin    string `json:"fechaFin"`
		Estado      int    `json:"estado"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Parsear fechas
	fechaInicio, err := time.Parse("2006-01-02", request.FechaInicio)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Formato de fechaInicio inválido (use YYYY-MM-DD)",
		})
		return
	}

	fechaFin, err := time.Parse("2006-01-02", request.FechaFin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Formato de fechaFin inválido (use YYYY-MM-DD)",
		})
		return
	}

	membresia := entities.Membresia{
		ID:          int32(id),
		UsuarioID:   request.UsuarioID,
		FechaInicio: fechaInicio,
		FechaFin:    fechaFin,
		Estado:      request.Estado,
	}

	updatedMembresia, err := ctrl.updateUseCase.Run(&membresia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar la membresía",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedMembresia)
}
