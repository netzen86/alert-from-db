// Package handlers - пакет содержит функции для обработки http запросов
package handlers

import (
	"alert-from-db/internal/data/films"
	"alert-from-db/internal/tmpl"
	"alert-from-db/proto/backend"
	pb "alert-from-db/proto/backend"
	"net/http"
	"strconv"
	"text/template"
)

// handler function #1 - returns the index.html template, with film data
func IndexFilmHandle(gCli backend.BackendClient, templates *template.Template, tmplname string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var f []films.Film
		grpcResp, err := gCli.GetFilms(r.Context(), &backend.GetFilmsRequest{})
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}
		for _, film := range grpcResp.Films {
			f = append(f, films.Film{ID: film.Id, Title: film.Title, Director: film.Directorr})
		}
		tmpl.RenderTemplate(w, templates, tmplname, f)
	}
}

// handler function #2 - returns the template block with the newly added film, as an HTMX response
func AddFilmHandle(gCli backend.BackendClient, templates *template.Template, tmplname string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request pb.AddFilmRequest
		var fs []films.Film
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		*request.Film = backend.Film{Title: title, Directorr: director}

		response, err := gCli.AddFilm(r.Context(), &request)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}

		for _, movie := range response.Films {
			fs = append(fs, films.Film{ID: movie.Id, Title: movie.Title, Director: movie.Directorr})
		}

		tmpl.RenderTemplate(w, templates, tmplname, fs)
	}
}

// handler function #3 - returns the template block with the deleted film, as an HTMX response
func DelFilmHandler(f *[]films.Film, templates *template.Template, tmplname string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PostFormValue("id"))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}
		f_tmp := *f
		*f = append(f_tmp[:id], f_tmp[id+1:]...)
		tmpl.RenderTemplate(w, templates, tmplname, *f)
	}
}
