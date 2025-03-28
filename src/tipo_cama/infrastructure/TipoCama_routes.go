package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type TipoCamaRouter struct {
	engine *gin.Engine
}

func NewTipoCamaRouter(engine *gin.Engine) *TipoCamaRouter {
	return &TipoCamaRouter{
		engine: engine,
	}
}

func (router *TipoCamaRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController, deleteController, getAllController := InitTipoCamaDependencies()

	// Grupo de rutas para tipos de cama
	tipoCamaGroup := router.engine.Group("/tipos-cama")
	{
		tipoCamaGroup.POST("/", createController.Run)
		tipoCamaGroup.GET("/:id", getByIdController.Run)
		tipoCamaGroup.PUT("/:id", updateController.Run)
		tipoCamaGroup.DELETE("/:id", deleteController.Run)
		tipoCamaGroup.GET("/", getAllController.Run)
	}
}
