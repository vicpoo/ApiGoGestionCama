// GetUsuarioByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/application"
)

type GetUsuarioByIdController struct {
	getByIdUseCase *application.GetUsuarioByIdUseCase
}

func NewGetUsuarioByIdController(getByIdUseCase *application.GetUsuarioByIdUseCase) *GetUsuarioByIdController {
	return &GetUsuarioByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetUsuarioByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	usuario, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, usuario)
}
