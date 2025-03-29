package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/application"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/domain/entities"
)

type UpdateAlmacenController struct {
	updateUseCase *application.UpdateAlmacenUseCase
}

func NewUpdateAlmacenController(updateUseCase *application.UpdateAlmacenUseCase) *UpdateAlmacenController {
	return &UpdateAlmacenController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateAlmacenController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

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
	almacen.ID = int32(id)

	updatedAlmacen, err := ctrl.updateUseCase.Run(almacen)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el registro de almacén",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedAlmacen)
}
