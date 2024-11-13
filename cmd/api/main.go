package main

import (
	"go-api-gin/config"
	"go-api-gin/internal/handler"
	"go-api-gin/internal/repository"
	"go-api-gin/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la configuración de la base de datos
	err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Crear las instancias necesarias
	userRepo := repository.NewUserRepository()         // Repositorio
	userService := service.NewUserService(userRepo)    // Servicio
	authHandler := handler.NewAuthHandler(userService) // Controlador

	// Configurar el router de Gin
	r := gin.Default()

	// Definir rutas y asignar handlers
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
	}

	// Iniciar el servidor
	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
