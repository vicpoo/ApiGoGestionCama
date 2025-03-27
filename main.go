package main

import (
	"log"

	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/core"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/rol/infrastructure"

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

	// Middleware adicional (ejemplo)
	router.Use(gin.Logger())   // Logger middleware
	router.Use(gin.Recovery()) // Recovery middleware

	// Inicializar rutas de roles
	rolRouter := infrastructure.NewRolRouter(router)
	rolRouter.Run()

	// Iniciar el servidor
	log.Println("API de Roles inicializada en http://localhost:8000")
	if err := router.Run(":8000"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
