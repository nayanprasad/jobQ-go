package main

import (
	"log/slog"
	"os"

	"github.com/nayanprasad/jobQ-go/internal/server"
)

func main() {
	opts := &slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)

	slog.Debug("ping")

	cnf := server.Config{
		Addr: ":5055",
		DSN:  "",
	}

	svr := server.New(cnf)

	h := svr.Mount()
	if err := svr.Run(h); err != nil {
		slog.Error("failied start the server", "error", err.Error())
		os.Exit(1)
	}
}
