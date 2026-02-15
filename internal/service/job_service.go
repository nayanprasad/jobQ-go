package service

import (
	"context"
	"log/slog"

	"github.com/nayanprasad/jobq-go/internal/domain/job"
	"github.com/nayanprasad/jobq-go/internal/domain/job/command"
	"github.com/nayanprasad/jobq-go/internal/repository"
)

type JobService interface {
	CreateJob(ctx context.Context, args command.CreateJobCommand) error
}

type svc struct {
	repo repository.JobRepository
	//db
}

func NewService(repo repository.JobRepository) *svc {
	return &svc{
		repo: repo,
	}
}

func (s *svc) CreateJob(ctx context.Context, cmd command.CreateJobCommand) error {

	j := &job.Job{
		Type:        cmd.JobType,
		Payload:     cmd.Payload,
		Status:      job.StatusPending,
		MaxRetries:  cmd.MaxRetries,
		RetryCount:  0,
		Priority:    cmd.Priority,
		AvailableAt: *cmd.ScheduledAt, // should validate this and set time
	}

	createdJob, err := s.repo.Create(ctx, j)
	if err != nil {
		slog.Error("Error while creating job", "error", err.Error())
		return err
	}

	slog.Info("created job", "job", createdJob)

	return nil
}
