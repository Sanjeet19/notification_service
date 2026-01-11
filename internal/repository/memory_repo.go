package repository

import (
	"notification-service/internal/models"
)

type MemoryRepo struct {
	Templates map[string]models.NotificationRequest
}

func NewMemoryRepo() *MemoryRepo {
	repo := &MemoryRepo{
		Templates: make(map[string]models.NotificationRequest),
	}
	repo.loadDefaultTemplates()
	return repo
}

func (r *MemoryRepo) GetTemplate(typeName string) (models.NotificationRequest, bool) {
	tpl, exists := r.Templates[typeName]
	return tpl, exists
}

func (r *MemoryRepo) loadDefaultTemplates() {
	r.Templates["error"] = models.NotificationRequest{
		Type: "error",
		Email: &models.EmailPayload{
			Subject: "System Error Alert",
			Body:    "A critical error has been detected in the microservice cluster.",
		},
		Slack: &models.SlackPayload{
			Channel: "#alerts",
			Message: "*Critical Error:* System health check failed.",
		},
		InApp: &models.InAppPayload{
			Title:   "System Error",
			Content: "A critical error occurred. Please check logs.",
		},
	}

	r.Templates["deployment"] = models.NotificationRequest{
		Type: "deployment",
		Slack: &models.SlackPayload{
			Channel: "#deployments",
			Message: "New version has been deployed successfully to production.",
		},
	}

	r.Templates["update"] = models.NotificationRequest{
		Type: "update",
		Email: &models.EmailPayload{
			Subject: "Weekly System Report",
			Body:    "Here is your summary of the system performance for this week.",
		},
		InApp: &models.InAppPayload{
			Title:   "Weekly Update",
			Content: "Your weekly performance report is now available.",
		},
	}
}
