package main

import (
	"apiMetrics/models"
	"apiMetrics/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup database connection
	db = models.SetupDatabase()

	// Setup Gin router
	router := gin.Default()
	routes.SetupRoutes(router, db)

	// Start server
	log.Fatal(router.Run(":8080"))
}
