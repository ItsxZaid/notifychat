package api

import (
	"encoding/json"
	"itsxzaid/notifychat/internal/app"
	"itsxzaid/notifychat/internal/domain"
	"itsxzaid/notifychat/internal/store/sqlc_generated"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ChannelHandler struct {
	*Handler
}

type channelConfig struct {
	// Telegram
	BotToken string `json:"bot_token" validate:"omittype,min=40,max=50"`
}

type createChannelRequest struct {
	Type     domain.ChannelType `json:"type" validate:"required,oneof=whatsapp telegram"`
	Config   channelConfig      `json:"config" validate:"required,min=8,max=150"`
	Template string             `json:"template"`
}

func NewChannelHandler(app *app.Application) *ChannelHandler {
	return &ChannelHandler{
		Handler: NewHandler(app),
	}
}

func (ch *ChannelHandler) RegisterRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{topic_id}", ch.ListAllChannels)
	router.Post("/{topic_id}", ch.CreateChannel)

	return router
}

func (ch *ChannelHandler) ListAllChannels(w http.ResponseWriter, r *http.Request) {
	logger := GetLogger(r.Context())

	topicID := chi.URLParam(r, "topic_id")
	if topicID == "" {
		ch.Error(w, r, http.StatusBadRequest, "invalid_id", "Missing topic ID")
		return
	}

	logger.Info("topicID", "topicID", topicID)

	channels, err := ch.app.Repo.ChannelStore.GetChannelsByTopicID(r.Context(), topicID)
	if err != nil {
		logger.Error("[ListAllChannels] validation failed", "err", err)
		ch.Error(w, r, http.StatusUnprocessableEntity, "validation_error", err.Error())
		return
	}

	ch.Success(w, r, 200, channels)
}

func (ch *ChannelHandler) CreateChannel(w http.ResponseWriter, r *http.Request) {
	logger := GetLogger(r.Context())

	topicID := chi.URLParam(r, "topic_id")
	if topicID == "" {
		ch.Error(w, r, http.StatusBadRequest, "invalid_id", "Missing topic ID")
		return
	}

	logger.Info("topicID", "topicID", topicID)

	var req createChannelRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("[CreateChannel] failed to decode request body", "err", err)
		ch.Error(w, r, http.StatusBadRequest, "invalid_json", "Invalid JSON body")
		return
	}

	if err := ch.app.Validator.Validate(&req); err != nil {
		logger.Error("[CreateChannel] validation failed", "err", err)
		ch.Error(w, r, http.StatusUnprocessableEntity, "validation_error", err.Error())
		return
	}

	var pgUUID pgtype.UUID

	if err := pgUUID.Scan(topicID); err != nil {
		logger.Error("[CreateChannel] topicId validation failed", "err", err)
		ch.Error(w, r, http.StatusUnprocessableEntity, "validation_error", "Invalid topic id")
		return
	}

	config, err := json.Marshal(req.Config)
	if err != nil {
		logger.Error("[CreateChannel] config validation failed", "err", err)
		ch.Error(w, r, http.StatusUnprocessableEntity, "validation_error", "Invalid config provided")
		return
	}

	channels, err := ch.app.Repo.ChannelStore.CreateChannel(r.Context(), sqlc_generated.CreateChannelParams{
		TopicID:  pgUUID,
		Type:     string(req.Type),
		Config:   config,
		Template: req.Template,
	})
	if err != nil {
		logger.Error("[ListAllChannels] validation failed", "err", err)
		ch.Error(w, r, http.StatusUnprocessableEntity, "validation_error", err.Error())
		return
	}

	ch.Success(w, r, 200, channels)
}
