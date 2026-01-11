package routes

import (
	"notification-service/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handlers.NotificationHandler) {

	api_v1 := r.Group("/api/v1")
	{
		api_v1.POST("/notifications", h.HandleNotification)

		api_v1.GET("/templates", h.GetTemplates)
	}
}
