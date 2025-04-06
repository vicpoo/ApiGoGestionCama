// GetTipoCamaByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/application"
)

type GetTipoCamaByIdController struct {
	getByIdUseCase *application.GetTipoCamaByIdUseCase
}

func NewGetTipoCamaByIdController(getByIdUseCase *application.GetTipoCamaByIdUseCase) *GetTipoCamaByIdController {
	return &GetTipoCamaByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetTipoCamaByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	tipoCama, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el tipo de cama",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tipoCama)
}
