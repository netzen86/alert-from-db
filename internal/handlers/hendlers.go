// Package handlers - пакет содержит функции для обработки http запросов
package handlers

import (
	"alert-from-db/internal/data/films"
	"alert-from-db/internal/tmpl"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// handler function #1 - returns the index.html template, with film data
func IndexFilmHandle(f []films.Film, templates *template.Template, tmplname string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.RenderTemplate(w, templates, tmplname, f)
	}
}

// handler function #2 - returns the template block with the newly added film, as an HTMX response
func AddFilmHandle(f *[]films.Film, templates *template.Template, tmplname string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		*f = append(*f, films.Film{ID: len(*f) + 1, Title: title, Director: director})
		tmpl.RenderTemplate(w, templates, tmplname, *f)
	}
}

// handler function #3 - returns the template block with the deleted film, as an HTMX response
func DelFilmHandler(f *[]films.Film, templates *template.Template, tmplname string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PostFormValue("id"))
		if err != nil {
			log.Fatalf("error parse id %v", err)
		}
		f_tmp := *f
		*f = append(f_tmp[:id], f_tmp[id+1:]...)
		tmpl.RenderTemplate(w, templates, tmplname, *f)
	}
}
