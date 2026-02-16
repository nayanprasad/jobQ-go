package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/nayanprasad/jobq-go/internal/domain/job"
	"github.com/nayanprasad/jobq-go/internal/domain/job/command"
	"github.com/nayanprasad/jobq-go/internal/repository"
)

type JobService interface {
	CreateJob(ctx context.Context, args command.CreateJobCommand) (job.Job, error)
}

type svc struct {
	repo repository.JobRepository
}

func NewService(repo repository.JobRepository) *svc {
	return &svc{
		repo: repo,
	}
}

func (s *svc) CreateJob(ctx context.Context, cmd command.CreateJobCommand) (job.Job, error) {

	j := &job.Job{
		Type:        cmd.JobType,
		Payload:     cmd.Payload,
		Status:      job.StatusPending,
		MaxRetries:  cmd.MaxRetries,
		RetryCount:  0,
		Priority:    cmd.Priority,
		AvailableAt: determineAvailableAt(cmd.ScheduledAt),
	}

	createdJob, err := s.repo.Create(ctx, j)
	if err != nil {
		slog.Error("Error while creating job", "error", err.Error())
		return job.Job{}, err
	}

	slog.Info("created job", "job", createdJob)

	return *createdJob, nil
}

func determineAvailableAt(t *time.Time) time.Time {
	if t != nil && t.After(time.Now()) {
		return *t
	}

	return time.Now()
}
