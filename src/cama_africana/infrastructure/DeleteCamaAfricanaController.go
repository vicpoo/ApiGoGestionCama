package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/application"
)

type DeleteCamaAfricanaController struct {
	deleteUseCase *application.DeleteCamaAfricanaUseCase
}

func NewDeleteCamaAfricanaController(deleteUseCase *application.DeleteCamaAfricanaUseCase) *DeleteCamaAfricanaController {
	return &DeleteCamaAfricanaController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteCamaAfricanaController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar la cama africana",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Cama africana eliminada exitosamente",
	})
}
