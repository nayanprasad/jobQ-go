package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/nayanprasad/jobq-go/internal/domain/job"
)

type JobRepository struct {
	db *pgx.Conn
}

func NewJobRepository(db *pgx.Conn) *JobRepository {
	return &JobRepository{
		db: db,
	}
}

func (r *JobRepository) Create(ctx context.Context, j *job.Job) (*job.Job, error) {

	query := `
		insert into jobs (type, payload, status, available_at, max_retries, retry_count, priority)
		values ($1, $2, $3, $4, $5, $6, $7)
		returning *
	`

	var createdJob job.Job
	err := r.db.QueryRow(ctx, query,
		j.Type,
		j.Payload,
		j.Status,
		j.AvailableAt,
		j.MaxRetries,
		j.RetryCount,
		j.Priority,
	).Scan(
		&createdJob.ID,
		&createdJob.Type,
		&createdJob.Payload,
		&createdJob.Status,
		&createdJob.AvailableAt,
		&createdJob.MaxRetries,
		&createdJob.RetryCount,
		&createdJob.Priority,
		&createdJob.CreatedAt,
		&createdJob.UpdatedAt,
		&createdJob.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &createdJob, nil
}

func (r *JobRepository) Get(ctx context.Context, id int32) (*job.Job, error) {
	query := `
		select *
		from jobs
		where id = $1
	`

	var j job.Job
	err := r.db.QueryRow(ctx, query, id).Scan(&j.ID,
		&j.Type,
		&j.Payload,
		&j.Status,
		&j.AvailableAt,
		&j.MaxRetries,
		&j.RetryCount,
		&j.Priority,
		&j.CreatedAt,
		&j.UpdatedAt,
		&j.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &j, nil
}
