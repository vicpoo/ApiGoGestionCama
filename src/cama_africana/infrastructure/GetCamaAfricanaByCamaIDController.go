package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/application"
)

type GetCamaAfricanaByCamaIDController struct {
	getByCamaIDUseCase *application.GetCamaAfricanaByCamaIDUseCase
}

func NewGetCamaAfricanaByCamaIDController(getByCamaIDUseCase *application.GetCamaAfricanaByCamaIDUseCase) *GetCamaAfricanaByCamaIDController {
	return &GetCamaAfricanaByCamaIDController{
		getByCamaIDUseCase: getByCamaIDUseCase,
	}
}

func (ctrl *GetCamaAfricanaByCamaIDController) Run(c *gin.Context) {
	camaIDParam := c.Param("camaId")
	camaID, err := strconv.Atoi(camaIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de cama inválido",
			"error":   err.Error(),
		})
		return
	}

	cama, err := ctrl.getByCamaIDUseCase.Run(int32(camaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener la cama africana",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, cama)
}
