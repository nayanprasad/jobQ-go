package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/nayanprasad/jobq-go/internal/appConfig"
	"github.com/nayanprasad/jobq-go/internal/handler/job"
	"github.com/nayanprasad/jobq-go/internal/worker"
)

const (
	configPath = "config/config.yaml"
)

func main() {
	ctx := context.Background()

	//logger
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)

	slog.Debug("worker")

	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found, using system environment variables")
	} else {
		slog.Info("loaded .env file")
	}

	appConfig, err := appConfig.LoadConfig(configPath)
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	conn, err := pgx.Connect(ctx, appConfig.Server.DB.DSN)
	if err != nil {
		slog.Error("unable to connect to db", "db", appConfig.Server.DB.DSN, "error", err.Error())
		os.Exit(1)
	}
	defer conn.Close(ctx)

	//worker setup
	cnf := worker.Config{
		PollInterval: appConfig.Worker.PollInterval,
		JobTimeout:   appConfig.Worker.JobTimeout,
		RetryBackoff: appConfig.Worker.RetryBackoff,
		MaxRetries:   appConfig.Worker.MaxRetries,
		Concurrency:  appConfig.Worker.Concurrency,
	}

	jobRegistry := job.NewRegistry()
	jobRegistry.Register(job.NewEmailHanlder())
	jobRegistry.Register(job.NewWebhookHanlder())

	w := worker.NewWorker(cnf, conn, jobRegistry)

	fmt.Print(w)

	//hanler worker
}
