package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/application"
)

type GetCamasByUsuarioIDController struct {
	getByUsuarioIDUseCase *application.GetCamasByUsuarioIDUseCase
}

func NewGetCamasByUsuarioIDController(getByUsuarioIDUseCase *application.GetCamasByUsuarioIDUseCase) *GetCamasByUsuarioIDController {
	return &GetCamasByUsuarioIDController{
		getByUsuarioIDUseCase: getByUsuarioIDUseCase,
	}
}

func (ctrl *GetCamasByUsuarioIDController) Run(c *gin.Context) {
	idParam := c.Param("usuarioId")
	usuarioID, err := strconv.Atoi(idParam)
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
			"message": "No se pudieron obtener las camas del usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, camas)
}
