// DeleteCamaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/application"
)

type DeleteCamaController struct {
	deleteUseCase *application.DeleteCamaUseCase
}

func NewDeleteCamaController(deleteUseCase *application.DeleteCamaUseCase) *DeleteCamaController {
	return &DeleteCamaController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteCamaController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar la cama",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Cama eliminada exitosamente",
	})
}
