package worker

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/nayanprasad/jobq-go/internal/handler/job"
)

type Worker struct {
	id          string
	db          *pgx.Conn
	jobRegistry *job.HandlerRegistry
	config      Config
}

type Config struct {
	PollInterval int
	JobTimeout   int
	RetryBackoff int
	MaxRetries   int
	Concurrency  int
}

func NewWorker(config Config, db *pgx.Conn, jobRegistry *job.HandlerRegistry) *Worker {
	return &Worker{
		id:          generateWorkerId(),
		config:      config,
		db:          db,
		jobRegistry: jobRegistry,
	}
}

func generateWorkerId() string {
	return fmt.Sprintf("worker-%s", uuid.New().String()[:8])

}

//start

//find next

// handle success

//handle failure

//retyr
