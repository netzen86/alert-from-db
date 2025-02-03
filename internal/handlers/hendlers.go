// Package handlers - пакет содержит функции для обработки http запросов
package handlers

import (
	"alert-from-db/internal/data/films"
	"alert-from-db/internal/tmpl"
	pb "alert-from-db/proto/backend"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// handler function #1 - returns the index.html template, with film data
func IndexFilmHandle(gCli pb.BackendClient, templates *template.Template, tmplname string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var f []films.Film
		grpcResp, err := gCli.GetFilms(r.Context(), &pb.GetFilmsRequest{})
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}
		for _, film := range grpcResp.Films {
			f = append(f, films.Film{ID: film.Id, Title: film.Title, Director: film.Director})
		}
		tmpl.RenderTemplate(w, templates, tmplname, f)
	}
}

// handler function #2 - returns the template block with the newly added film, as an HTMX response
func AddFilmHandle(gCli pb.BackendClient, templates *template.Template, tmplname string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request pb.AddFilmRequest
		var fs []films.Film
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		request.Film = &pb.Film{Title: title, Director: director}
		response, err := gCli.AddFilm(r.Context(), &request)
		log.Println("handler", err)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}

		for _, movie := range response.Films {
			fs = append(fs, films.Film{ID: movie.Id, Title: movie.Title, Director: movie.Director})
		}

		tmpl.RenderTemplate(w, templates, tmplname, fs)
	}
}

// handler function #3 - returns the template block with the deleted film, as an HTMX response
func DelFilmHandler(gCli pb.BackendClient, templates *template.Template, tmplname string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request pb.DelFilmRequest
		var id int64
		var fs []films.Film
		var err error

		id, err = strconv.ParseInt(r.PostFormValue("id"), 10, 64)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}

		request.Id = id

		response, err := gCli.DelFilm(r.Context(), &request)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}

		for _, movie := range response.Films {
			fs = append(fs, films.Film{ID: movie.Id, Title: movie.Title, Director: movie.Director})
		}

		tmpl.RenderTemplate(w, templates, tmplname, fs)
	}
}
