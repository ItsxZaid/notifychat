package app

import (
	"itsxzaid/notifychat/internal/config"
	"itsxzaid/notifychat/internal/store"
	"itsxzaid/notifychat/internal/validator"
	"log/slog"
)

type Repository struct {
	TopicStore *store.TopicStore
}

type Application struct {
	Config    *config.Config
	Logger    *slog.Logger
	Repo      *Repository
	Validator *validator.Validator
}
