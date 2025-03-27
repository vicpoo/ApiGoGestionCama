package main

import (
	"log"

	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/core"
	rolInfra "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/rol/infrastructure"
	usuarioInfra "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/infrastructure"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la conexión a la base de datos
	core.InitDB()

	// Crear un router con Gin
	router := gin.Default()

	// Configuración de CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Permite todos los orígenes (ajusta para producción)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600, // 12 horas de caché para preflight requests
	}))

	// Middleware adicional
	router.Use(gin.Logger())   // Logger middleware
	router.Use(gin.Recovery()) // Recovery middleware

	// Inicializar rutas de roles
	rolRouter := rolInfra.NewRolRouter(router)
	rolRouter.Run()

	// Inicializar rutas de usuarios
	usuarioRouter := usuarioInfra.NewUsuarioRouter(router)
	usuarioRouter.Run()

	// Iniciar el servidor
	log.Println("API inicializada en http://localhost:8000")
	log.Println("- Rutas de roles: /roles")
	log.Println("- Rutas de usuarios: /usuarios")
	if err := router.Run(":8000"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
