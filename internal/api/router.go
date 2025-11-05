package api

import (
	"itsxzaid/notifychat/internal/app"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func SetupRouter(app *app.Application) chi.Router {
	appRouter := chi.NewRouter()

	appRouter.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))

	appRouter.Use(middleware.Recoverer)
	appRouter.Use(loggerMiddleware(app.Logger))

	campaignHandler := NewCampaignHandler(app)

	appRouter.Route("/api/v1", func(r chi.Router) {
		r.Mount("/campaigns", campaignHandler.RegisterRoutes())
	})
	return appRouter
}
