// GetAllAlmacenesController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/application"
)

type GetAllAlmacenesController struct {
	getAllUseCase *application.GetAllAlmacenesUseCase
}

func NewGetAllAlmacenesController(getAllUseCase *application.GetAllAlmacenesUseCase) *GetAllAlmacenesController {
	return &GetAllAlmacenesController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllAlmacenesController) Run(c *gin.Context) {
	almacenes, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los registros de almacén",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, almacenes)
}
