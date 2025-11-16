package main

import (
	"database/sql"
	"log/slog"
	"os"

	_ "modernc.org/sqlite"
)

func main() {
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dbType: "sqlite",
			dsn:    "ecom.db",
		},
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	db, err := sql.Open(cfg.db.dbType, cfg.db.dsn)
	if err != nil {
		logger.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		logger.Error("failed to ping database", "error", err)
		os.Exit(1)
	}

	logger.Info("connected to database")

	app := &app{
		config: cfg,
		db: db,
	}

	err = app.run(app.mount())
	if err != nil {
		slog.Error("server has failed to start", "error", err)
		os.Exit(1)
	}
}
