package repository

import (
	"context"

	"github.com/nayanprasad/jobq-go/internal/domain/job"
)

// this define contract not storing logic here
type JobRepository interface {
	Create(ctx context.Context, job *job.Job) (*job.Job, error)
}
