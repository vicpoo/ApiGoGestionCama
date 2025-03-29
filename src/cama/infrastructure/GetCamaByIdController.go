package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/application"
)

type GetCamaByIdController struct {
	getByIdUseCase *application.GetCamaByIdUseCase
}

func NewGetCamaByIdController(getByIdUseCase *application.GetCamaByIdUseCase) *GetCamaByIdController {
	return &GetCamaByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetCamaByIdController) Run(c *gin.Context) {
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
			"message": "No se pudo obtener la cama",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, cama)
}
