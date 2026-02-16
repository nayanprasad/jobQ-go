package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/nayanprasad/jobq-go/internal/appConfig"
	"github.com/nayanprasad/jobq-go/internal/transport/http"
	"gopkg.in/yaml.v3"
)

const (
	configPath = "config/config.yaml"
)

func main() {

	ctx := context.Background()

	//logger
	opts := &slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)

	slog.Debug("ping")

	//load .env
	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found, using system environment variables")
	} else {
		slog.Info("loaded .env file")
	}

	//load config
	appCnf, err := loadConfig(configPath)
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	//connect db
	conn, err := pgx.Connect(ctx, appCnf.Server.DB.DSN)
	if err != nil {
		slog.Error("unable to connecto to db", "db", appCnf.Server.DB.DSN, "error", err.Error())
		os.Exit(1)
	}
	defer conn.Close(ctx)

	slog.Info("connected to db", "dsn", appCnf.Server.DB.DSN)

	// server setup
	cnf := http.Config{
		Addr: fmt.Sprintf(":%d", appCnf.Server.Port),
		DSN:  appCnf.Server.DB.DSN,
	}

	server := http.New(cnf, conn)

	h := server.Mount()
	if err := server.Run(h); err != nil {
		slog.Error("failied start the server", "error", err.Error())
		os.Exit(1)
	}
}

func loadConfig(path string) (*appConfig.AppConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// this will replace ${VAR} with actual values from the env
	expanded := os.ExpandEnv(string(data))

	var cfg appConfig.AppConfig
	if err := yaml.Unmarshal([]byte(expanded), &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
