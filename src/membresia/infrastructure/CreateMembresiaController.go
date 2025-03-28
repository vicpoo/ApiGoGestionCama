// CreateMembresiaController.go
package infrastructure

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/application"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain/entities"
)

type CreateMembresiaController struct {
	createUseCase *application.CreateMembresiaUseCase
}

func NewCreateMembresiaController(createUseCase *application.CreateMembresiaUseCase) *CreateMembresiaController {
	return &CreateMembresiaController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateMembresiaController) Run(c *gin.Context) {
	var request struct {
		UsuarioID   int32  `json:"usuarioId" binding:"required"`
		FechaInicio string `json:"fechaInicio" binding:"required"`
		FechaFin    string `json:"fechaFin" binding:"required"`
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

	// Valor por defecto para estado si no se especifica
	if request.Estado == 0 {
		request.Estado = 1 // 1 = activo por defecto
	}

	membresia := entities.NewMembresia(
		request.UsuarioID,
		fechaInicio,
		fechaFin,
		request.Estado,
	)

	createdMembresia, err := ctrl.createUseCase.Run(membresia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear la membresía",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdMembresia)
}
