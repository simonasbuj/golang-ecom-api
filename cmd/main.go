package main

import (
	"database/sql"
	repo "golang-ecom-api/internal/adapters/sqlite/sqlc"
	"log/slog"
	"os"

	_ "modernc.org/sqlite"
)

func main() {
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dbType: "sqlite",
			dsn: "ecom.db",
		},
	}

	app := &app{
		config: cfg,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	
	db, err := sql.Open(cfg.db.dbType, cfg.db.dsn)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		slog.Error("failed to ping database", "error", err)
		os.Exit(1)
	}

	_ = repo.New(db)

	err = app.run(app.mount())
	if err != nil {
		slog.Error("server has failed to start", "error", err)
		os.Exit(1)
	}
}
