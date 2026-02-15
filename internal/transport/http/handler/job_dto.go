package handler

import (
	"encoding/json"

	"github.com/nayanprasad/jobq-go/internal/domain/job"
)

type CreateJobRequest struct {
	JobType     job.JobType     `json:"job_type"`
	Payload     json.RawMessage `json:"payload"`
	Priority    int             `json:"priority"`
	MaxRetries  int             `json:"max_retries"`
	ScheduledAt *string         `json:"scheduled_at"` // string because to avoid decerialization issue. * because its optional
}

type CreateJobResponse struct {
	Success bool `json:"success"`
}
