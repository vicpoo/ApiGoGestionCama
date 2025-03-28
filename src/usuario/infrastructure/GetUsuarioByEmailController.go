// GetUsuarioByEmailController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/application"
)

type GetUsuarioByEmailController struct {
	getByEmailUseCase *application.GetUsuarioByEmailUseCase
}

func NewGetUsuarioByEmailController(getByEmailUseCase *application.GetUsuarioByEmailUseCase) *GetUsuarioByEmailController {
	return &GetUsuarioByEmailController{
		getByEmailUseCase: getByEmailUseCase,
	}
}

func (ctrl *GetUsuarioByEmailController) Run(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email es requerido",
		})
		return
	}

	usuario, err := ctrl.getByEmailUseCase.Run(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, usuario)
}
