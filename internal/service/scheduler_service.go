package service

import (
	"fmt"
	"notification-service/internal/models"

	"github.com/robfig/cron/v3"
)

type SchedulerService struct {
	cron         *cron.Cron
	notifService *NotificationService
}

func NewSchedulerService(ns *NotificationService) *SchedulerService {
	c := cron.New()
	c.Start()
	return &SchedulerService{
		cron:         c,
		notifService: ns,
	}
}

func (s *SchedulerService) Schedule(req models.NotificationRequest) error {
	_, err := s.cron.AddFunc(req.CronSchedule, func() {
		fmt.Printf("[SCHEDULER] Executing scheduled task: %s\n", req.Type)
		s.notifService.SendInstant(req)
	})

	return err
}
