package models

import (
	"time"
)

type NotificationRequest struct {
	Type string `json:"type"`

	Email *EmailPayload `json:"email,omitempty"`
	Slack *SlackPayload `json:"slack,omitempty"`
	InApp *InAppPayload `json:"in_app,omitempty"`

	CronSchedule string `json:"cron_schedule,omitempty"`
}

type EmailPayload struct {
	Receiver string `json:"receiver"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
}

type SlackPayload struct {
	Channel string `json:"channel"`
	Message string `json:"message"`
}

type InAppPayload struct {
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

type Template struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Payload     NotificationRequest `json:"payload"`
}
