// GetActiveMembresiaByUsuarioIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/application"
)

type GetActiveMembresiaByUsuarioIDController struct {
	getActiveUseCase *application.GetActiveMembresiaByUsuarioIDUseCase
}

func NewGetActiveMembresiaByUsuarioIDController(getActiveUseCase *application.GetActiveMembresiaByUsuarioIDUseCase) *GetActiveMembresiaByUsuarioIDController {
	return &GetActiveMembresiaByUsuarioIDController{
		getActiveUseCase: getActiveUseCase,
	}
}

func (ctrl *GetActiveMembresiaByUsuarioIDController) Run(c *gin.Context) {
	usuarioIDParam := c.Param("usuarioId")
	usuarioID, err := strconv.Atoi(usuarioIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	membresia, err := ctrl.getActiveUseCase.Run(int32(usuarioID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener la membresía activa del usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, membresia)
}
