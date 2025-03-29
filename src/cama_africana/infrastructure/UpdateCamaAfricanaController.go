package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/application"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain/entities"
)

type UpdateCamaAfricanaController struct {
	updateUseCase *application.UpdateCamaAfricanaUseCase
}

func NewUpdateCamaAfricanaController(updateUseCase *application.UpdateCamaAfricanaUseCase) *UpdateCamaAfricanaController {
	return &UpdateCamaAfricanaController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateCamaAfricanaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var request struct {
		CamaID    int32  `json:"cama_id"`
		UsuarioID *int32 `json:"usuario_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	cama := entities.CamaAfricana{
		ID:        int32(id),
		CamaID:    request.CamaID,
		UsuarioID: request.UsuarioID,
	}

	updatedCama, err := ctrl.updateUseCase.Run(&cama)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar la cama africana",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedCama)
}
