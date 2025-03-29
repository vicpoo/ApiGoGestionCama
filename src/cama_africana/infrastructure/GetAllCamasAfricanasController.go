package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/application"
)

type GetAllCamasAfricanasController struct {
	getAllUseCase *application.GetAllCamasAfricanasUseCase
}

func NewGetAllCamasAfricanasController(getAllUseCase *application.GetAllCamasAfricanasUseCase) *GetAllCamasAfricanasController {
	return &GetAllCamasAfricanasController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllCamasAfricanasController) Run(c *gin.Context) {
	camas, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las camas africanas",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, camas)
}
