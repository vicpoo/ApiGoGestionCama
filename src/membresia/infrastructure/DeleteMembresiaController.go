// DeleteMembresiaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/application"
)

type DeleteMembresiaController struct {
	deleteUseCase *application.DeleteMembresiaUseCase
}

func NewDeleteMembresiaController(deleteUseCase *application.DeleteMembresiaUseCase) *DeleteMembresiaController {
	return &DeleteMembresiaController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteMembresiaController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar la membresía",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Membresía eliminada exitosamente",
	})
}
