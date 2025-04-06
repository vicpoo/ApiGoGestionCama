// CreateAlmacenController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/application"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/domain/entities"
)

type CreateAlmacenController struct {
	createUseCase *application.CreateAlmacenUseCase
}

func NewCreateAlmacenController(createUseCase *application.CreateAlmacenUseCase) *CreateAlmacenController {
	return &CreateAlmacenController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateAlmacenController) Run(c *gin.Context) {
	var almacenRequest struct {
		TipoCamaID int32 `json:"tipo_cama_id"`
		Cantidad   int32 `json:"cantidad"`
	}

	if err := c.ShouldBindJSON(&almacenRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	almacen := entities.NewAlmacen(
		almacenRequest.TipoCamaID,
		almacenRequest.Cantidad,
	)

	createdAlmacen, err := ctrl.createUseCase.Run(almacen)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el registro en almacén",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdAlmacen)
}
