package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/application"
)

type GetCamaAfricanaByIdController struct {
	getByIdUseCase *application.GetCamaAfricanaByIdUseCase
}

func NewGetCamaAfricanaByIdController(getByIdUseCase *application.GetCamaAfricanaByIdUseCase) *GetCamaAfricanaByIdController {
	return &GetCamaAfricanaByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetCamaAfricanaByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	cama, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener la cama africana",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, cama)
}
