package app

import (
	"itsxzaid/notifychat/internal/config"
	"itsxzaid/notifychat/internal/store"
	"itsxzaid/notifychat/internal/validator"
	"log/slog"
)

type Repository struct {
	CampaignStore *store.CampaignStore
}

type Application struct {
	Config    *config.Config
	Logger    *slog.Logger
	Repo      *Repository
	Validator *validator.Validator
}
