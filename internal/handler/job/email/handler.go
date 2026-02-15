package email

import (
	"context"
	"errors"

	"github.com/nayanprasad/jobq-go/internal/domain/job"
	"github.com/nayanprasad/jobq-go/pkg/json"
)

type EmailHandler struct {
}

func (*EmailHandler) Type() job.JobType {
	return job.JobTypeEmail
}

func (*EmailHandler) Execute(ctx context.Context, payload []byte) error {
	var emailPayload EmailPayload
	if err := json.Read(payload, emailPayload); err != nil {
		return errors.New("failed to read payload " + err.Error())
	}

	//sent email

	return nil
}
