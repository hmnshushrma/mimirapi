package main

import (
	"log"
	"os"

	"mirmiapi/config"
	"mirmiapi/platform/handler"
	"mirmiapi/platform/repository"
	"mirmiapi/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	platformRoutes "mirmiapi/routes/platform"
)

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Init DB
	database := config.ConnectDB()

	// Setup Gin
	router := gin.Default()

	// Register routes
	routes.PingRoute(router)

	platformRepo := repository.NewPlatformUserRepo(database)
	platformHandler := handler.NewPlatformHandler(platformRepo)
	platformRoutes := platformRoutes.NewPlatformHandler(platformHandler)

	platformRoutes.RegisterPlatformRoutes(router)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server running at http://localhost:%s", port)
	router.Run(":" + port)
}
