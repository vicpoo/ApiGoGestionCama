// cama_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type CamaRouter struct {
	engine *gin.Engine
}

func NewCamaRouter(engine *gin.Engine) *CamaRouter {
	return &CamaRouter{
		engine: engine,
	}
}

func (router *CamaRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController, deleteController,
		getAllController, getByUsuarioController, getByTipoController := InitCamaDependencies()

	// Grupo de rutas para camas
	camaGroup := router.engine.Group("/camas")
	{
		// Rutas CRUD básicas
		camaGroup.POST("/", createController.Run)
		camaGroup.GET("/:id", getByIdController.Run)
		camaGroup.PUT("/:id", updateController.Run)
		camaGroup.DELETE("/:id", deleteController.Run)
		camaGroup.GET("/", getAllController.Run)

		// Rutas adicionales específicas
		camaGroup.GET("/usuario/:usuarioId", getByUsuarioController.Run)
		camaGroup.GET("/tipo/:tipoId", getByTipoController.Run)
	}
}
