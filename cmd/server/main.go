package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"solution-for-x/notifychat/internal/api"
	"solution-for-x/notifychat/internal/app"
	"solution-for-x/notifychat/internal/config"

	"github.com/jackc/pgx/v5"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v", err)
		os.Exit(1)
	}

	dbCtx := context.Background()
	conn, err := pgx.Connect(dbCtx, cfg.DatabaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(dbCtx)

	app := &app.Application{
		Config: cfg,
		Logger: logger,
	}

	router := api.SetupRouter(app)
	http.ListenAndServe(":"+cfg.Port, router)
}
