// UpdateCamaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/application"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/domain/entities"
)

type UpdateCamaController struct {
	updateUseCase *application.UpdateCamaUseCase
}

func NewUpdateCamaController(updateUseCase *application.UpdateCamaUseCase) *UpdateCamaController {
	return &UpdateCamaController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateCamaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var camaRequest struct {
		Modelo    string `json:"modelo" binding:"required"`
		TipoID    int32  `json:"tipo_id" binding:"required"`
		UsuarioID *int32 `json:"usuario_id,omitempty"`
	}

	if err := c.ShouldBindJSON(&camaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	cama := entities.NewCama(
		camaRequest.Modelo,
		camaRequest.TipoID,
		camaRequest.UsuarioID,
	)
	cama.ID = int32(id)

	updatedCama, err := ctrl.updateUseCase.Run(cama)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar la cama",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedCama)
}
