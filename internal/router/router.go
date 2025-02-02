// Package router - пакет в котором описаны endpoints
package router

import (
	"alert-from-db/internal/handlers"
	"alert-from-db/proto/backend"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func GetGateway(gCli backend.BackendClient, templates *template.Template, tmplname string) chi.Router {

	gw := chi.NewRouter()

	gw.Route("/", func(gw chi.Router) {
		gw.Get("/", handlers.IndexFilmHandle(gCli, templates, tmplname))
		gw.Post("/add-film/", handlers.AddFilmHandle(gCli, templates, tmplname))
		gw.Post("/del-film/", handlers.DelFilmHandler(gCli, templates, tmplname))

		// Define the routes for serving profiling data
		gw.Mount("/debug", middleware.Profiler())
	},
	)
	return gw
}
