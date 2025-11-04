package api

import (
	"itsxzaid/notifychat/internal/app"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CampaignHandler struct {
	*Handler
}

func NewCampaignHandler(app *app.Application) *CampaignHandler {
	return &CampaignHandler{
		Handler: NewHandler(app),
	}
}

func (ch *CampaignHandler) RegisterRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/list", ch.ListCampaigns)
	return router
}

func (ch *CampaignHandler) ListCampaigns(w http.ResponseWriter, r *http.Request) {

}
