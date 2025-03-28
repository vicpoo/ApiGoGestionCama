// GetMembresiasByUsuarioIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/application"
)

type GetMembresiasByUsuarioIDController struct {
	getByUsuarioIDUseCase *application.GetMembresiasByUsuarioIDUseCase
}

func NewGetMembresiasByUsuarioIDController(getByUsuarioIDUseCase *application.GetMembresiasByUsuarioIDUseCase) *GetMembresiasByUsuarioIDController {
	return &GetMembresiasByUsuarioIDController{
		getByUsuarioIDUseCase: getByUsuarioIDUseCase,
	}
}

func (ctrl *GetMembresiasByUsuarioIDController) Run(c *gin.Context) {
	usuarioIDParam := c.Param("usuarioId")
	usuarioID, err := strconv.Atoi(usuarioIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	membresias, err := ctrl.getByUsuarioIDUseCase.Run(int32(usuarioID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las membresías del usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, membresias)
}
