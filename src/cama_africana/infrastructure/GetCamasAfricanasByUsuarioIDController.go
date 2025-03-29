package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/application"
)

type GetCamasAfricanasByUsuarioIDController struct {
	getByUsuarioIDUseCase *application.GetCamasAfricanasByUsuarioIDUseCase
}

func NewGetCamasAfricanasByUsuarioIDController(getByUsuarioIDUseCase *application.GetCamasAfricanasByUsuarioIDUseCase) *GetCamasAfricanasByUsuarioIDController {
	return &GetCamasAfricanasByUsuarioIDController{
		getByUsuarioIDUseCase: getByUsuarioIDUseCase,
	}
}

func (ctrl *GetCamasAfricanasByUsuarioIDController) Run(c *gin.Context) {
	usuarioIDParam := c.Param("usuarioId")
	usuarioID, err := strconv.Atoi(usuarioIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	camas, err := ctrl.getByUsuarioIDUseCase.Run(int32(usuarioID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las camas africanas del usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, camas)
}
