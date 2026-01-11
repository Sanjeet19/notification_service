package main

import (
	"log"
	"notification-service/internal/api/handlers"
	"notification-service/internal/api/routes"
	"notification-service/internal/repository"
	"notification-service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	//Set up Assets
	repo := repository.NewMemoryRepo()
	notifService := service.NewNotificationService(repo)
	schedulerService := service.NewSchedulerService(notifService)
	handler := handlers.NewNotificationHandler(notifService, schedulerService)
	routes.SetupRoutes(r, handler)

	log.Println("Notification Service starting on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
