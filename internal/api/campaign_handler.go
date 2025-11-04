package api

import (
	"encoding/json"
	"itsxzaid/notifychat/internal/app"
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
	return router
}

func (ch *CampaignHandler) ListCampaigns(w http.ResponseWriter, r *http.Request) {
	// Todo
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
	return
}
