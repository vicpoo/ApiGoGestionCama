package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/application"
)

type DeleteAlmacenController struct {
	deleteUseCase *application.DeleteAlmacenUseCase
}

func NewDeleteAlmacenController(deleteUseCase *application.DeleteAlmacenUseCase) *DeleteAlmacenController {
	return &DeleteAlmacenController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteAlmacenController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar el registro de almacén",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Registro de almacén eliminado exitosamente",
	})
}
