// CreateTipoCamaController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/application"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/domain/entities"
)

type CreateTipoCamaController struct {
	createUseCase *application.CreateTipoCamaUseCase
}

func NewCreateTipoCamaController(createUseCase *application.CreateTipoCamaUseCase) *CreateTipoCamaController {
	return &CreateTipoCamaController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateTipoCamaController) Run(c *gin.Context) {
	var tipoCamaRequest struct {
		Nombre string `json:"nombre"`
		Clima  string `json:"clima"`
	}

	if err := c.ShouldBindJSON(&tipoCamaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	tipoCama := entities.NewTipoCama(
		tipoCamaRequest.Nombre,
		tipoCamaRequest.Clima,
	)

	createdTipoCama, err := ctrl.createUseCase.Run(tipoCama)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el tipo de cama",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdTipoCama)
}
