// Package router - пакет в котором описаны endpoints
package router

import (
	"alert-from-db/internal/data/films"
	"alert-from-db/internal/handlers"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func GetGateway(f *[]films.Film, templates *template.Template, tmplname string) chi.Router {

	gw := chi.NewRouter()

	gw.Route("/", func(gw chi.Router) {
		gw.Get("/", handlers.IndexFilmHandle(*f, templates, tmplname))
		gw.Post("/add-film/", handlers.AddFilmHandle(f, templates, tmplname))
		gw.Post("/del-film/", handlers.DelFilmHandler(f, templates, tmplname))

		// Define the routes for serving profiling data
		gw.Mount("/debug", middleware.Profiler())
	},
	)
	return gw
}
