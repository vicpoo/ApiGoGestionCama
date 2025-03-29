package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/application"
)

type GetCamasByTipoIDController struct {
	getByTipoIDUseCase *application.GetCamasByTipoIDUseCase
}

func NewGetCamasByTipoIDController(getByTipoIDUseCase *application.GetCamasByTipoIDUseCase) *GetCamasByTipoIDController {
	return &GetCamasByTipoIDController{
		getByTipoIDUseCase: getByTipoIDUseCase,
	}
}

func (ctrl *GetCamasByTipoIDController) Run(c *gin.Context) {
	idParam := c.Param("tipoId")
	tipoID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de tipo inválido",
			"error":   err.Error(),
		})
		return
	}

	camas, err := ctrl.getByTipoIDUseCase.Run(int32(tipoID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las camas por tipo",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, camas)
}
