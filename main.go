package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// Swagger docs
	_ "facturama-api/docs"

	// Swagger handler
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Rutas
	"facturama-api/routes"
)

// @title Facturama API
// @version 1.0
// @description API para generar y consultar facturas CFDI usando Facturama
// @termsOfService http://swagger.io/terms/
// @contact.name Emmanuel Rdz
// @contact.email emmanuelrdz@example.com
// @license.name MIT
// @host localhost:8080
// @BasePath /

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Archivo .env no encontrado, se usarán valores por defecto si aplica")
	}

	// Inicializar router
	r := gin.Default()

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Registrar rutas de la API
	routes.RegisterRoutes(r)

	// Puerto configurable desde .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Servidor iniciado en http://localhost:%s/", port)
	r.Run(":" + port)
}
