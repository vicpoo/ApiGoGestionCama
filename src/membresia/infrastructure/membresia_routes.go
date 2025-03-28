// membresia_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type MembresiaRouter struct {
	engine *gin.Engine
}

func NewMembresiaRouter(engine *gin.Engine) *MembresiaRouter {
	return &MembresiaRouter{
		engine: engine,
	}
}

func (router *MembresiaRouter) Run() {
	// Inicializar todos los controladores necesarios pasando la conexión a la DB
	createController, getByIdController, updateController,
		deleteController, getAllController, getByUsuarioIDController,
		getActiveController, getExpiringController, renewController := InitMembresiaDependencies()

	membresiaGroup := router.engine.Group("/membresias")
	{
		// Rutas CRUD básicas
		membresiaGroup.POST("/", createController.Run)
		membresiaGroup.GET("/:id", getByIdController.Run)
		membresiaGroup.PUT("/:id", updateController.Run)
		membresiaGroup.DELETE("/:id", deleteController.Run)
		membresiaGroup.GET("/", getAllController.Run)

		// Rutas específicas para membresías
		membresiaGroup.GET("/usuario/:usuarioId", getByUsuarioIDController.Run)
		membresiaGroup.GET("/usuario/:usuarioId/activa", getActiveController.Run)
		membresiaGroup.GET("/por-vencer", getExpiringController.Run)
		membresiaGroup.PUT("/:id/renovar", renewController.Run)
	}
}
