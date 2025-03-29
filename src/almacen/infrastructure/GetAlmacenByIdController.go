package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/application"
)

type GetAlmacenByIdController struct {
	getByIdUseCase *application.GetAlmacenByIdUseCase
}

func NewGetAlmacenByIdController(getByIdUseCase *application.GetAlmacenByIdUseCase) *GetAlmacenByIdController {
	return &GetAlmacenByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetAlmacenByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	almacen, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el registro de almacén",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, almacen)
}
