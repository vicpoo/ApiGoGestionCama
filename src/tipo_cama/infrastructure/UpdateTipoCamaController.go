package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/application"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/domain/entities"
)

type UpdateTipoCamaController struct {
	updateUseCase *application.UpdateTipoCamaUseCase
}

func NewUpdateTipoCamaController(updateUseCase *application.UpdateTipoCamaUseCase) *UpdateTipoCamaController {
	return &UpdateTipoCamaController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateTipoCamaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

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
	tipoCama.ID = int32(id)

	updatedTipoCama, err := ctrl.updateUseCase.Run(tipoCama)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el tipo de cama",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedTipoCama)
}
