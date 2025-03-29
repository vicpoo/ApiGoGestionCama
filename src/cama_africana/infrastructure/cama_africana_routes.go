package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type CamaAfricanaRouter struct {
	engine *gin.Engine
}

func NewCamaAfricanaRouter(engine *gin.Engine) *CamaAfricanaRouter {
	return &CamaAfricanaRouter{
		engine: engine,
	}
}

func (router *CamaAfricanaRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController,
		deleteController, getAllController, getByCamaIDController,
		getByUsuarioIDController, assignController := InitCamaAfricanaDependencies()

	// Grupo de rutas base
	camaGroup := router.engine.Group("/camas-africanas")
	{
		// Rutas CRUD básicas
		camaGroup.POST("/", createController.Run)
		camaGroup.GET("/", getAllController.Run)
		camaGroup.GET("/:id", getByIdController.Run)
		camaGroup.PUT("/:id", updateController.Run)
		camaGroup.DELETE("/:id", deleteController.Run)

		// Rutas específicas
		camaGroup.GET("/por-cama/:camaId", getByCamaIDController.Run)
		camaGroup.GET("/por-usuario/:usuarioId", getByUsuarioIDController.Run)
		camaGroup.POST("/:camaId/asignar", assignController.Run)
	}
}
