package handler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/nayanprasad/jobq-go/internal/domain/job/command"
	"github.com/nayanprasad/jobq-go/internal/service"
	"github.com/nayanprasad/jobq-go/pkg/json"
)

type JobHandler struct {
	service service.JobService
}

func NewHandler(s service.JobService) *JobHandler {
	return &JobHandler{
		service: s,
	}
}

func (h JobHandler) CreateJob(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateJobRequest
	if err := json.ReadRequest(r, &reqBody); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmd := command.CreateJobCommand{
		JobType:    reqBody.JobType,
		Payload:    reqBody.Payload,
		Priority:   reqBody.Priority,
		MaxRetries: reqBody.MaxRetries,
	}

	if reqBody.ScheduledAt != nil {
		t, _ := time.Parse(time.RFC3339, *reqBody.ScheduledAt)
		cmd.ScheduledAt = &t
	}

	j, err := h.service.CreateJob(r.Context(), cmd)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := CreateJobResponse{
		Success: true,
		Data: data{
			Job: j,
		},
	}

	json.WriteResponse(w, http.StatusCreated, res)
}
