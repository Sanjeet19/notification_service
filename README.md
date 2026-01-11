# Notification Service

A lightweight internal notification service designed to be used by microservices to send instant and scheduled notifications across Email, Slack, and In-App channels.

# Overview

- **Intended use:** Internal API for microservices to trigger notifications (instant or scheduled).
- **Assumptions & logic:**
  - Few kinds of messages (e.g., error, deployment, update) are common and often delivered across multiple channels.
  - Requests are JSON payloads describing which channels to use.
  - Templates (common payloads) are available and merged with incoming requests by `type`.
  - Scheduling supported via cron-style `cron_schedule`.
  - As this is a test project, did not making it more complex buy adding a db for managing the templates and crons, as it would only increase complexity.

## Key Files

- `cmd/server/main.go` — server startup (runs on `:8080`)
- `internal/api/routes/routes.go` — HTTP route setup
- `internal/api/handlers/notification_handler.go` — request handling
- `internal/service/notification_service.go` — delivery logic and template merging
- `internal/service/scheduler_service.go` — cron schedule integration (uses `robfig/cron`)
- `internal/repository/memory_repo.go` — in-memory templates (default templates loaded)
- `internal/models/notification.go` — request/payload models

---

## API Endpoints

### POST /api/v1/notifications

- Use to send a notification instantly or schedule one.
- Request Content-Type: `application/json`
- Request body: `NotificationRequest` (see schema below)
- Responses:
  - `200 OK` — Sent instantly
  - `202 Accepted` — Successfully scheduled
  - `400` — Invalid payload
  - `500` — Internal error (send or schedule failure)

### GET /api/v1/templates

- Returns available templates (map of template type → payload).

---

## Payloads (JSON)

NotificationRequest (simplified)

```json
{
  "type": "error", // can be a blank string
  "email": {
    // can be an empty dict
    "receiver": "ops@example.com",
    "subject": "Alert",
    "body": "Details..."
  },
  "slack": {
    // can be an empty dict
    "channel": "#alerts",
    "message": "Critical!"
  },
  "in_app": {
    // can be an empty dict
    "user_id": "user123",
    "title": "Notice",
    "content": "..."
  },
  "cron_schedule": "0 0 * * *" // can be a blank string
}
```

- If `cron_schedule` is present, the request is scheduled (cron format).
- If `type` is present and matches a template, the template is merged with the request; explicit fields in the request override the template fields.

---

## Templates

Default templates (from `memory_repo.go`):

- `error` — Email, Slack (#alerts), In-App title/content
- `deployment` — Slack to `#deployments`
- `update` — Weekly email + in-app summary

Use GET `/api/v1/templates` to list all currently loaded templates.

---

## Scheduler

- Payload includes `cron_schedule` (string). Example: `"*/5 * * * *"` runs every 5 minutes.
- The schedule payload is the same `NotificationRequest`. Blank channel objects are ignored (only non-nil channels are delivered).
- The scheduler uses `robfig/cron` to schedule tasks and invokes the same delivery logic as instant sends.

---

## Behavior & Notes

- Delivery functions are currently mocked and output to stdout. For production use, the following integrations can be implemented:
  - Email: SMTP, SendGrid, or Amazon SES
  - Slack: Webhooks or the Slack API
  - In-app: Persist messages to a database and deliver via the user service
- Templates are currently stored in memory. For production deployments, they should be persisted in a database.
- The service does not currently implement authentication. As this is intended to be an internal service, authentication can be handled by the surrounding infrastructure or gateway.
- Logging and metrics can be added to improve observability.

---

## Quickstart (local)

Prereqs: Go installed (1.18+)

```bash
# 1) clone and change into repo
mkdir notification-service
cd notification_service
# copy the code here

# 2) install dependencies
go mod tidy

# 3) Start the server(local)
go run cmd/server/main.go

# Service will start on :8080
```

Example requests:

- Send Instant Notification on all channel

```bash
curl -X POST http://localhost:8080/api/v1/notifications \
  -H 'Content-Type: application/json' \
  -d '{
    "type" : "",
    "email": {"receiver": "admin@test.com", "subject": "Tri-Channel", "body": "Testing all 3"},
    "slack": {"channel": "#general", "message": "Testing all 3"},
    "in_app": {"user_id": "user123", "content": "Testing all 3"}
  }'
```

- Schedule a deployment notification every 1 minute

```bash
curl -X POST http://localhost:8080/api/v1/notifications \
  -H 'Content-Type: application/json' \
  -d '{
    "type": "deployment",
    "cron_schedule": "*/1 * * * *"
  }'
```

- Fetch all available templates

```bash
curl -X GET http://localhost:8080/api/v1/templates
```

---

## Suggestions / TODOs 💡

- Add real delivery integrations (SMTP provider, Slack API).
- Add persistent template store (DB) and template management endpoints (CRUD).
- Add authentication/authorization (internal service auth).
- Add unit/integration tests and CI pipeline.
- Add graceful shutdown and scheduled-job persistence across restarts.

---
