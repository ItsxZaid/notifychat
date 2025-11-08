package api

import (
	"encoding/json"
	"errors"
	"itsxzaid/notifychat/internal/app"
	"itsxzaid/notifychat/internal/domain"
	"itsxzaid/notifychat/internal/store/sqlc_generated"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type TopicHandler struct {
	*Handler
}

type createTopicRequest struct {
	Name        string  `json:"name" validate:"required,min=3,max=100"`
	Description *string `json:"description" validate:"omitempty,min=8,max=150"`
}

type updateTopicRequest struct {
	Name        string  `json:"name" validate:"omitempty,min=3,max=100"`
	Description *string `json:"description" validate:"omitempty,min=8,max=150"`
}

func NewTopicHandler(app *app.Application) *TopicHandler {
	return &TopicHandler{
		Handler: NewHandler(app),
	}
}

func (th *TopicHandler) RegisterRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", th.ListTopics)
	router.Post("/", th.CreateTopic)
	router.Patch("/{topic_id}", th.UpdateTopic)
	router.Delete("/{topic_id}", th.DeleteTopic)
	return router
}

func (th *TopicHandler) ListTopics(w http.ResponseWriter, r *http.Request) {
	logger := GetLogger(r.Context())

	topics, err := th.app.Repo.TopicStore.GetAllTopics(r.Context())
	if err != nil {
		logger.Error("[ListTopic] validation failed", "err", err)
		th.Error(w, r, http.StatusUnprocessableEntity, "validation_error", err.Error())
		return
	}

	th.Success(w, r, 200, topics)
}

func (th *TopicHandler) CreateTopic(w http.ResponseWriter, r *http.Request) {
	logger := GetLogger(r.Context())

	var req createTopicRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("[CreateTopic] failed to decode request body", "err", err)
		th.Error(w, r, http.StatusBadRequest, "invalid_json", "Invalid JSON body")
		return
	}

	if err := th.app.Validator.Validate(&req); err != nil {
		logger.Error("[CreateTopic] validation failed", "err", err)
		th.Error(w, r, http.StatusUnprocessableEntity, "validation_error", err.Error())
		return
	}

	var pgDescription pgtype.Text

	if req.Description != nil {
		pgDescription = pgtype.Text{
			String: *req.Description,
			Valid:  true,
		}
	}

	topic, err := th.app.Repo.TopicStore.CreateTopic(r.Context(), sqlc_generated.CreateTopicParams{
		Name:        req.Name,
		Description: pgDescription,
	})

	if err != nil {
		logger.Error("[CreateTopic] validation failed", "err", err)
		th.Error(w, r, http.StatusUnprocessableEntity, "validation_error", err.Error())
		return
	}

	th.Success(w, r, 201, topic)
}

func (th *TopicHandler) UpdateTopic(w http.ResponseWriter, r *http.Request) {
	logger := GetLogger(r.Context())

	topicID := chi.URLParam(r, "topic_id")
	if topicID == "" {
		th.Error(w, r, http.StatusBadRequest, "invalid_id", "Missing topic ID")
		return
	}

	logger.Info("topicID", "topicID", topicID)

	var req updateTopicRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("[UpdateTopic] failed to decode request body", "err", err)
		th.Error(w, r, http.StatusBadRequest, "invalid_json", "Invalid JSON body")
		return
	}

	if err := th.app.Validator.Validate(&req); err != nil {
		logger.Error("[UpdateTopic] validation failed", "err", err)
		th.Error(w, r, http.StatusUnprocessableEntity, "validation_error", err.Error())
		return
	}

	var pgDescription pgtype.Text

	if req.Description != nil {
		pgDescription = pgtype.Text{
			String: *req.Description,
			Valid:  true,
		}
	}

	var pgUUID pgtype.UUID

	if err := pgUUID.Scan(topicID); err != nil {
		th.Error(w, r, http.StatusBadRequest, "invalid_id", "The topic ID is not a valid UUID")
		return
	}

	updatedTopic, err := th.app.Repo.TopicStore.UpdateTopic(r.Context(), sqlc_generated.UpdateTopicParams{
		ID:          pgUUID,
		Name:        req.Name,
		Description: pgDescription,
	})

	if err != nil {
		logger.Error("[UpdateTopic] error while updating topic", "err", err.Error())
		th.Error(w, r, http.StatusInternalServerError, "server_error", "Internal server error")
		return
	}

	th.Success(w, r, 200, updatedTopic)
}

func (th *TopicHandler) DeleteTopic(w http.ResponseWriter, r *http.Request) {
	logger := GetLogger(r.Context())

	topicID := chi.URLParam(r, "topic_id")
	if topicID == "" {
		th.Error(w, r, http.StatusBadRequest, "invalid_id", "Missing topic ID")
		return
	}

	logger.Info("topicID", "topicID", topicID)

	err := th.app.Repo.TopicStore.DeleteTopic(r.Context(), topicID)

	if err != nil {
		logger.Error("[DeleteTopic] error while deleting topic", "err", err.Error())

		if errors.Is(err, domain.ErrInvalidInput) {
			th.Error(w, r, http.StatusBadRequest, "invalid_id", "The topic ID is not a valid UUID")
			return
		}

		th.Error(w, r, http.StatusInternalServerError, "server_error", "Internal server error")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
