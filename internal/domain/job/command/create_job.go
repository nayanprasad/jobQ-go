package command

import (
	"encoding/json"
	"time"

	"github.com/nayanprasad/jobq-go/internal/domain/job"
)

type CreateJobCommand struct {
	JobType     job.JobType     `json:"job_type"`
	Payload     json.RawMessage `json:"payload"`
	Priority    int             `json:"priority"`
	MaxRetries  int             `json:"max_retries"`
	ScheduledAt *time.Time      `json:"scheduled_at"`
}
