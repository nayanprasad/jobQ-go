package job

import (
	"context"

	"github.com/nayanprasad/jobq-go/internal/domain/job"
)

type Handler interface {
	Type() job.JobType
	Execute(ctx context.Context, payload []byte) error
}
