package webhook

type WebhookPayload struct {
	Url   string `json:"url"`
	Event string `json:"event"`
	Data  any    `json:"data"`
}
