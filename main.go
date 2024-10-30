package main

import (
	"fkedweb/config"
	"fkedweb/controllers"
	"fkedweb/models"
	"fkedweb/repositories"
	"fkedweb/routes"
	"fkedweb/services"
	"fkedweb/utils"
	"github.com/jinzhu/gorm"
	"log"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Connect to the database
	db, err := utils.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}(db)

	// Run migrations
	if err := db.AutoMigrate(&models.Todo{}).Error; err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize repository, service, and controller
	todoRepo := repositories.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepo)
	todoController := controllers.NewTodoController(todoService)

	// Set up routes
	r := routes.SetupRouter(todoController)

	// Start the server
	if err := r.Run("localhost:8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
