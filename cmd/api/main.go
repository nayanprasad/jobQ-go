package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/nayanprasad/jobq-go/internal/appConfig"
	"github.com/nayanprasad/jobq-go/internal/transport/http"
	"gopkg.in/yaml.v3"
)

const (
	configPath = "config/config.yaml"
)

func main() {
	//logger
	opts := &slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)

	slog.Debug("ping")

	//load config
	appCnf, err := loadConfig(configPath)
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	// server setup
	cnf := http.Config{
		Addr: fmt.Sprintf(":%d", appCnf.Server.Port),
		DSN:  appCnf.Server.DB.DSN,
	}

	server := http.New(cnf)

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

	var cfg appConfig.AppConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
