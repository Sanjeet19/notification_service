package service

import (
	"fmt"
	"notification-service/internal/models"
	"notification-service/internal/repository"
)

type NotificationService struct {
	repo *repository.MemoryRepo
}

func NewNotificationService(repo *repository.MemoryRepo) *NotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) SendInstant(req models.NotificationRequest) error {
	finalPayload := req

	if req.Type != "" {
		template, exists := s.repo.GetTemplate(req.Type)
		if exists {
			finalPayload = s.mergeTemplateWithRequest(template, req)
		}
	}

	if finalPayload.Email != nil {
		s.sendEmail(finalPayload.Email)
	}
	if finalPayload.Slack != nil {
		s.sendSlack(finalPayload.Slack)
	}
	if finalPayload.InApp != nil {
		s.sendInApp(finalPayload.InApp)
	}

	return nil
}

func (s *NotificationService) GetAvailableTemplates() map[string]models.NotificationRequest {
	return s.repo.Templates
}

func (s *NotificationService) mergeTemplateWithRequest(tpl, req models.NotificationRequest) models.NotificationRequest {
	if req.Email != nil {
		tpl.Email = req.Email
	}
	if req.Slack != nil {
		tpl.Slack = req.Slack
	}
	if req.InApp != nil {
		tpl.InApp = req.InApp
	}
	return tpl
}

// Mock delivery functions (As I don't have the fare details)
func (s *NotificationService) sendEmail(p *models.EmailPayload) {
	fmt.Printf("[EMAIL] Sending to %s: Subject %s: Body: %s\n", p.Receiver, p.Subject, p.Body)
}

func (s *NotificationService) sendSlack(p *models.SlackPayload) {
	fmt.Printf("[SLACK] Channel %s: %s\n", p.Channel, p.Message)
}

func (s *NotificationService) sendInApp(p *models.InAppPayload) {
	fmt.Printf("[IN-APP] User %s: %s\n", p.UserID, p.Content)
}
