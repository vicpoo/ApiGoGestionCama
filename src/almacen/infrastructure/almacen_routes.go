// almacen_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type AlmacenRouter struct {
	engine *gin.Engine
}

func NewAlmacenRouter(engine *gin.Engine) *AlmacenRouter {
	return &AlmacenRouter{
		engine: engine,
	}
}

func (router *AlmacenRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController,
		deleteController, getAllController, incrementarController := InitAlmacenDependencies()

	// Grupo de rutas para almacén
	almacenGroup := router.engine.Group("/almacen")
	{
		almacenGroup.POST("/", createController.Run)
		almacenGroup.GET("/:id", getByIdController.Run)
		almacenGroup.PUT("/:id", updateController.Run)
		almacenGroup.DELETE("/:id", deleteController.Run)
		almacenGroup.GET("/", getAllController.Run)
		almacenGroup.PATCH("/:id/incrementar", incrementarController.Run)
	}
}
