package api

import (
	"solution-for-x/notifychat/cmd/app"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(app *app.Application) chi.Router {
	appRouter := chi.NewRouter()

	appRouter.Use(middleware.Recoverer)
	appRouter.Use(loggerMiddleware(app.Logger))

	campaignHandler := NewCampaignHandler(app)

	appRouter.Route("/api/v1", func(r chi.Router) {
		r.Mount("/campaigns", campaignHandler.RegisterRoutes())
	})
	return appRouter
}
