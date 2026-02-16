package job

import (
	"errors"
	"log/slog"

	"github.com/nayanprasad/jobq-go/internal/domain/job"
)

type HandlerRegistry struct {
	handlers map[job.JobType]Handler
}

func NewRegistry() *HandlerRegistry {
	return &HandlerRegistry{
		handlers: make(map[job.JobType]Handler),
	}
}

func (r *HandlerRegistry) Register(h Handler) error {
	jobType := h.Type()

	if !jobType.IsValid() {
		slog.Debug("Invalid job type", "job", jobType)
		return errors.New("Invalid job type " + string(jobType))
	}

	r.handlers[h.Type()] = h

	return nil
}

func (r *HandlerRegistry) Get(j job.JobType) Handler {
	return r.handlers[j]
}
