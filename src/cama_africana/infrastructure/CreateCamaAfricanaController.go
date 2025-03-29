package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/application"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain/entities"
)

type CreateCamaAfricanaController struct {
	createUseCase *application.CreateCamaAfricanaUseCase
}

func NewCreateCamaAfricanaController(createUseCase *application.CreateCamaAfricanaUseCase) *CreateCamaAfricanaController {
	return &CreateCamaAfricanaController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateCamaAfricanaController) Run(c *gin.Context) {
	var request struct {
		CamaID    int32  `json:"cama_id" binding:"required"`
		UsuarioID *int32 `json:"usuario_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	cama := entities.NewCamaAfricana(request.CamaID, request.UsuarioID)

	createdCama, err := ctrl.createUseCase.Run(cama)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear la cama africana",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdCama)
}
