package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/application"
)

type DeleteTipoCamaController struct {
	deleteUseCase *application.DeleteTipoCamaUseCase
}

func NewDeleteTipoCamaController(deleteUseCase *application.DeleteTipoCamaUseCase) *DeleteTipoCamaController {
	return &DeleteTipoCamaController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteTipoCamaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	errDelete := ctrl.deleteUseCase.Run(int32(id))
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo eliminar el tipo de cama",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Tipo de cama eliminado exitosamente",
	})
}
