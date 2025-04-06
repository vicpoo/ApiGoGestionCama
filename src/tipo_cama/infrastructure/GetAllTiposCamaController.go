// GetAllTiposCamaController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/application"
)

type GetAllTiposCamaController struct {
	getAllUseCase *application.GetAllTiposCamaUseCase
}

func NewGetAllTiposCamaController(getAllUseCase *application.GetAllTiposCamaUseCase) *GetAllTiposCamaController {
	return &GetAllTiposCamaController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllTiposCamaController) Run(c *gin.Context) {
	tiposCama, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los tipos de cama",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tiposCama)
}
