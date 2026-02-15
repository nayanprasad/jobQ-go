package webhook

import (
	"context"
	"errors"

	"github.com/nayanprasad/jobq-go/internal/domain/job"
	"github.com/nayanprasad/jobq-go/pkg/json"
)

type WebhookHandler struct {
}

func (*WebhookHandler) Type() job.JobType {
	return job.JobTypeWebhook
}

func (*WebhookHandler) Execute(ctx context.Context, payload []byte) error {
	var webhookPayload WebhookPayload
	if err := json.Read(payload, webhookPayload); err != nil {
		return errors.New("failed to read payload " + err.Error())
	}

	//tigger webhook

	return nil
}
