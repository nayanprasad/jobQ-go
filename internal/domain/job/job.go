package job

import (
	"database/sql"
	"time"
)

type Job struct {
	ID           int32          `json:"id"`
	Type         JobType        `json:"type"`
	Payload      []byte         `json:"payload"`
	Status       JobStatus      `json:"status"`
	AvailableAt  time.Time      `json:"available_at"`
	MaxRetries   int            `json:"max_retries"`
	RetryCount   int            `json:"RetryCount"`
	Priority     int            `json:"priority"`
	WorkerId     sql.NullString `json:"worker_id,omitempty"`
	ClaimedAt    sql.NullTime   `json:"claimed_at,omitempty"`
	CompletedAt  sql.NullTime   `json:"completed_at,omitempty"`
	FailedAt     sql.NullTime   `json:"failed_at,omitempty"`
	ErrorMessage sql.NullString `json:"error_message,omitempty"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    sql.NullTime   `json:"deleted_at,omitempty"`
}
