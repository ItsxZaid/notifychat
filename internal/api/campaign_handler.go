package api

import (
	"encoding/json"
	"errors"
	"itsxzaid/notifychat/internal/app"
	"itsxzaid/notifychat/internal/domain"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CampaignHandler struct {
	*Handler
}

type CreateCampaignRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}

func NewCampaignHandler(app *app.Application) *CampaignHandler {
	return &CampaignHandler{
		Handler: NewHandler(app),
	}
}

func (ch *CampaignHandler) RegisterRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", ch.ListCampaigns)
	router.Post("/", ch.CreateCampaign)
	router.Delete("/{campaign_id}", ch.DeleteCampaign)
	return router
}

func (ch *CampaignHandler) ListCampaigns(w http.ResponseWriter, r *http.Request) {
	logger := GetLogger(r.Context())

	campaigns, err := ch.app.Repo.CampaignStore.GetAllCampaigns(r.Context())
	if err != nil {
		logger.Error("[ListCampaign] validation failed", "err", err)
		ch.Error(w, r, http.StatusUnprocessableEntity, "validation_error", err.Error())
		return
	}

	ch.Success(w, r, 200, campaigns)
}

func (ch *CampaignHandler) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	logger := GetLogger(r.Context())

	var req CreateCampaignRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("[CreateCampaign] failed to decode request body", "err", err)
		ch.Error(w, r, http.StatusBadRequest, "invalid_json", "Invalid JSON body")
		return
	}

	if err := ch.app.Validator.Validate(&req); err != nil {
		logger.Error("[CreateCampaign] validation failed", "err", err)
		ch.Error(w, r, http.StatusUnprocessableEntity, "validation_error", err.Error())
		return
	}

	campaign, err := ch.app.Repo.CampaignStore.CreateCampaign(r.Context(), req.Name)
	if err != nil {
		logger.Error("[CreateCampaign] validation failed", "err", err)
		ch.Error(w, r, http.StatusUnprocessableEntity, "validation_error", err.Error())
		return
	}

	ch.Success(w, r, 201, campaign)
}

func (ch *CampaignHandler) DeleteCampaign(w http.ResponseWriter, r *http.Request) {
	logger := GetLogger(r.Context())

	campaignID := chi.URLParam(r, "campaign_id")
	if campaignID == "" {
		ch.Error(w, r, http.StatusBadRequest, "invalid_id", "Missing campaign ID")
		return
	}

	logger.Info("campaignID", "campaignID", campaignID)

	err := ch.app.Repo.CampaignStore.DeleteCampaign(r.Context(), campaignID)

	if err != nil {
		logger.Error("[DeleteCampaign] error while deleting campaign", "err", err.Error())

		if errors.Is(err, domain.ErrInvalidInput) {
			ch.Error(w, r, http.StatusBadRequest, "invalid_id", "The campaign ID is not a valid UUID")
			return
		}

		ch.Error(w, r, http.StatusInternalServerError, "server_error", "Internal server error")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
