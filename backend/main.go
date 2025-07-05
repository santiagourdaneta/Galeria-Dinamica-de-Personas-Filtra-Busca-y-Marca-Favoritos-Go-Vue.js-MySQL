package main

import (
	"log"
	"os"

	"hot-or-not/backend/config"
	"hot-or-not/backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors" // Ensure this is imported for CORS
)

func main() {
	// Database connection setup
	config.ConnectDB()

	// Initialize Gin router
	router := gin.Default()

	// Configure CORS middleware - ALL THIS CODE MUST BE INSIDE main()
	corsConfig := cors.DefaultConfig()
	// IMPORTANT: In production, change this to your actual frontend domain(s)
	corsConfig.AllowOrigins = []string{os.Getenv("FRONTEND_URL")}
	if corsConfig.AllowOrigins[0] == "" { // Fallback for local development
		corsConfig.AllowOrigins = []string{"http://localhost:5173"}
	}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))

	// Setup API routes
	routes.SetupAPIRoutes(router)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port
	}
	log.Printf("Servidor Go escuchando en :%s", port)
	log.Fatal(router.Run(":" + port))
}