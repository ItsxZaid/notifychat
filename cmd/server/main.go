package main

import (
	"context"
	"fmt"
	"itsxzaid/notifychat/internal/api"
	"itsxzaid/notifychat/internal/app"
	"itsxzaid/notifychat/internal/config"
	"itsxzaid/notifychat/internal/store"
	"itsxzaid/notifychat/internal/validator"
	"log/slog"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v", err)
		os.Exit(1)
	}

	m, err := migrate.New(
		"file://postgres/migrations",
		cfg.DatabaseURL,
	)
	if err != nil {
		logger.Error("failed to create migration instance", "err", err)
		os.Exit(1)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Error("failed to run migrations", "err", err)
		os.Exit(1)
	}

	logger.Info("database migrations finished successfully")

	dbConn, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbConn.Close()

	validator := validator.NewValidator()

	campaignStore := store.NewCampaignStore(dbConn)

	repo := &app.Repository{
		CampaignStore: campaignStore,
	}

	app := &app.Application{
		Config:    cfg,
		Logger:    logger,
		Repo:      repo,
		Validator: validator,
	}

	router := api.SetupRouter(app)
	http.ListenAndServe(":"+cfg.Port, router)
}
