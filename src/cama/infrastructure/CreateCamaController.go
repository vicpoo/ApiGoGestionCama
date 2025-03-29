package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/application"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/domain/entities"
)

type CreateCamaController struct {
	createUseCase *application.CreateCamaUseCase
}

func NewCreateCamaController(createUseCase *application.CreateCamaUseCase) *CreateCamaController {
	return &CreateCamaController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateCamaController) Run(c *gin.Context) {
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

	createdCama, err := ctrl.createUseCase.Run(cama)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear la cama",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdCama)
}
