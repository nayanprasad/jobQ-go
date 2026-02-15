package postgres

import (
	"context"

	"github.com/nayanprasad/jobq-go/internal/domain/job"
)

type JobRepository struct {
	db any
}

func NewJobRepository(db any) *JobRepository {
	return &JobRepository{
		db: db,
	}
}

func (r *JobRepository) Create(ctx context.Context, j *job.Job) (*job.Job, error) {

	//query here

	return &job.Job{}, nil
}
