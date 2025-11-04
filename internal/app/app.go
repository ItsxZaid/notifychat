package app

import (
	"itsxzaid/notifychat/internal/config"
	"log/slog"
)

type Application struct {
	Config *config.Config
	Logger *slog.Logger
}
