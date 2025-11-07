package api

import (
	"itsxzaid/notifychat/internal/app"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ChannelHandler struct {
	*Handler
}

func NewChannelHandler(app *app.Application) *ChannelHandler {
	return &ChannelHandler{
		Handler: NewHandler(app),
	}
}

func (ch *ChannelHandler) RegisterRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{topic_id}", ch.ListAllChannels)

	return router
}

func (ch *ChannelHandler) ListAllChannels(w http.ResponseWriter, r *http.Request) {
	// TODO BUDDDY
}
