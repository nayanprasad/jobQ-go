package job

import (
	"context"

	"github.com/nayanprasad/jobq-go/internal/domain/job"
	"github.com/nayanprasad/jobq-go/internal/handler/job/email"
	"github.com/nayanprasad/jobq-go/internal/handler/job/webhook"
)

type Handler interface {
	Type() job.JobType
	Execute(ctx context.Context, payload []byte) error
}

func NewEmailHanlder() *email.EmailHandler {
	return &email.EmailHandler{}
}

func NewWebhookHanlder() *webhook.WebhookHandler {
	return &webhook.WebhookHandler{}
}
