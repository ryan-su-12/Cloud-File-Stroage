package main

import (
	"go-gin-app/config"
	"log"
	/*
	"log"
	"os"
	
	"github.com/gin-gonic/gin"
	"Backend/config"
	"Backend/auth"
	"Backend/handlers"
	*/
)

func main() {
	log.Println("🚀 Starting the application...")

	// Initialize AWS connection
	config.InitAWS()

	log.Println("✅ AWS setup complete!")
}
/*
func main() {
	// Initialize AWS and JWT setup
	config.InitAWS()

	router := gin.Default()

	// Public routes
	router.POST("/login", auth.LoginUser)

	// Protected routes (require JWT)
	protected := router.Group("/")
	protected.Use(auth.ValidateToken)
	protected.POST("/upload", handlers.UploadFile)
	protected.GET("/files/:id", handlers.GetFile)
	protected.DELETE("/files/:id", handlers.DeleteFile)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)
	router.Run(":" + port)
}
	*/
