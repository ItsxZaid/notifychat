package app

import (
	"itsxzaid/notifychat/internal/config"
	"itsxzaid/notifychat/internal/service"
	"itsxzaid/notifychat/internal/validator"
	"log/slog"
)

type Service struct {
	TopicService   *service.TopicService
	ChannelService *service.ChannelService
}

type Application struct {
	Config    *config.Config
	Logger    *slog.Logger
	Service   *Service
	Validator *validator.Validator
}
