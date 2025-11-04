package app

import (
	"log/slog"
	"solution-for-x/notifychat/internal/config"
)

type Application struct {
	Config *config.Config
	Logger *slog.Logger
}
