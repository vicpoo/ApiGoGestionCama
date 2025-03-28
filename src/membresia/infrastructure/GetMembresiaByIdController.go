// GetMembresiaByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/application"
)

type GetMembresiaByIdController struct {
	getByIdUseCase *application.GetMembresiaByIdUseCase
}

func NewGetMembresiaByIdController(getByIdUseCase *application.GetMembresiaByIdUseCase) *GetMembresiaByIdController {
	return &GetMembresiaByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetMembresiaByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	membresia, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener la membresía",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, membresia)
}
