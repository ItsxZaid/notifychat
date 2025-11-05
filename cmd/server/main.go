package main

import (
	"context"
	"errors"
	"fmt"
	"itsxzaid/notifychat/internal/api"
	"itsxzaid/notifychat/internal/app"
	"itsxzaid/notifychat/internal/config"
	"itsxzaid/notifychat/internal/store"
	"itsxzaid/notifychat/internal/validator"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := &http.Server{Addr: "0.0.0.0:" + cfg.Port, Handler: router}

	go func() {
		logger.Info("Server listening on http://localhost:" + cfg.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	logger.Info("awaiting shutdown signal...")
	<-ctx.Done()
	logger.Info("shutdown signal received, starting graceful shutdown...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatal(err)
	}
}
