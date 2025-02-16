package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"go-gin-app/config"
	"go-gin-app/handlers"
)

func main() {
	log.Println("🚀 Starting the application...")

	// Initialize AWS connection
	config.InitAWS()

	// Create Gin router
	router := gin.Default()

	// Routes
	router.POST("/upload", handlers.UploadFile)
	router.GET("/files/:id", handlers.GetFile)
	router.DELETE("/files/:id", handlers.DeleteFile)
	router.GET("/list-files", handlers.ListFiles)

	// Start server
	log.Println("✅ Server is running on http://localhost:8080")
	router.Run(":8080")
}
