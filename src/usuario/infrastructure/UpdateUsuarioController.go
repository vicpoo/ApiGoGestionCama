package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/application"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/domain/entities"
)

type UpdateUsuarioController struct {
	updateUseCase *application.UpdateUsuarioUseCase
}

func NewUpdateUsuarioController(updateUseCase *application.UpdateUsuarioUseCase) *UpdateUsuarioController {
	return &UpdateUsuarioController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateUsuarioController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

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
	usuario.ID = int32(id)

	updatedUsuario, err := ctrl.updateUseCase.Run(usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedUsuario)
}
