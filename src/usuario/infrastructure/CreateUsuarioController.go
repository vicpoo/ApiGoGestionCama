// CreateUsuarioController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/application"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/domain/entities"
)

type CreateUsuarioController struct {
	createUseCase *application.CreateUsuarioUseCase
}

func NewCreateUsuarioController(createUseCase *application.CreateUsuarioUseCase) *CreateUsuarioController {
	return &CreateUsuarioController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateUsuarioController) Run(c *gin.Context) {
	var usuarioRequest struct {
		Name      string `json:"name"`
		Lastname  string `json:"lastname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		RoleID    *int32 `json:"roleId"`
		IsPremium bool   `json:"isPremium"`
	}

	if err := c.ShouldBindJSON(&usuarioRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	usuario := entities.NewUsuario(
		usuarioRequest.Name,
		usuarioRequest.Lastname,
		usuarioRequest.Email,
		usuarioRequest.Password,
		usuarioRequest.RoleID,
		usuarioRequest.IsPremium,
	)

	createdUsuario, err := ctrl.createUseCase.Run(usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdUsuario)
}
