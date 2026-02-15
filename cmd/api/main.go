package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/nayanprasad/jobQ-go/internal/api"
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
	appConfig, err := loadConfig(configPath)
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	// server setup
	cnf := api.Config{
		Addr: fmt.Sprintf(":%d", appConfig.Server.Port),
		DSN:  appConfig.Server.DB.DSN,
	}

	app := api.New(cnf)

	h := app.Mount()
	if err := app.Run(h); err != nil {
		slog.Error("failied start the server", "error", err.Error())
		os.Exit(1)
	}
}

func loadConfig(path string) (*AppConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg AppConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
