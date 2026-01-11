package handlers

import (
	"net/http"
	"notification-service/internal/models"
	"notification-service/internal/service"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	notifService     *service.NotificationService
	schedulerService *service.SchedulerService
}

func NewNotificationHandler(ns *service.NotificationService, ss *service.SchedulerService) *NotificationHandler {
	return &NotificationHandler{
		notifService:     ns,
		schedulerService: ss,
	}
}

func (h *NotificationHandler) HandleNotification(c *gin.Context) {
	var req models.NotificationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload: " + err.Error()})
		return
	}

	if req.CronSchedule != "" {
		err := h.schedulerService.Schedule(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to schedule notification"})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"message": "Notification scheduled successfully", "schedule": req.CronSchedule})
	} else {
		err := h.notifService.SendInstant(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send notification"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Notification sent instantly"})
	}
}

func (h *NotificationHandler) GetTemplates(c *gin.Context) {
	templates := h.notifService.GetAvailableTemplates()
	c.JSON(http.StatusOK, templates)
}
