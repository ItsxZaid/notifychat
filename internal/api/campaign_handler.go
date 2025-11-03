package api

import (
	"net/http"
	"solution-for-x/notifychat/cmd/app"

	"github.com/go-chi/chi/v5"
)

type CampaignHandler struct {
	app *app.Application
}

func NewCampaignHandler(app *app.Application) *CampaignHandler {
	return &CampaignHandler{
		app: app,
	}
}

func (ch *CampaignHandler) RegisterRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/list", ch.ListCampaigns)
	return router
}

func (ch *CampaignHandler) ListCampaigns(w http.ResponseWriter, r *http.Request) {

}
