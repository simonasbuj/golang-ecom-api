package main

import (
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
	}

	app := &app{
		config: cfg,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	err := app.run(app.mount())
	if err != nil {
		slog.Error("server has failed to start", "error", err)
		os.Exit(1)
	}
}
