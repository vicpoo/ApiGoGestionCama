// usuario_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type UsuarioRouter struct {
	engine *gin.Engine
}

func NewUsuarioRouter(engine *gin.Engine) *UsuarioRouter {
	return &UsuarioRouter{
		engine: engine,
	}
}

func (router *UsuarioRouter) Run() {
	createController, getByIdController, updateController,
		deleteController, getAllController, getByEmailController,
		authController := InitUsuarioDependencies()

	usuarioGroup := router.engine.Group("/usuarios")
	{
		usuarioGroup.POST("/", createController.Run)
		usuarioGroup.POST("/login", authController.Login)
		usuarioGroup.GET("/:id", getByIdController.Run)
		usuarioGroup.PUT("/:id", updateController.Run)
		usuarioGroup.DELETE("/:id", deleteController.Run)
		usuarioGroup.GET("/", getAllController.Run)
		usuarioGroup.GET("/by-email", getByEmailController.Run)
	}
}
