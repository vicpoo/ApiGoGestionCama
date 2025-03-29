package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/application"
)

type AssignCamaToUsuarioController struct {
	assignUseCase *application.AssignCamaToUsuarioUseCase
}

func NewAssignCamaToUsuarioController(assignUseCase *application.AssignCamaToUsuarioUseCase) *AssignCamaToUsuarioController {
	return &AssignCamaToUsuarioController{
		assignUseCase: assignUseCase,
	}
}

func (ctrl *AssignCamaToUsuarioController) Run(c *gin.Context) {
	camaIDParam := c.Param("camaId")
	camaID, err := strconv.Atoi(camaIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de cama inválido",
			"error":   err.Error(),
		})
		return
	}

	var request struct {
		UsuarioID *int32 `json:"usuario_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	assignedCama, err := ctrl.assignUseCase.Run(int32(camaID), request.UsuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo asignar la cama al usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, assignedCama)
}
