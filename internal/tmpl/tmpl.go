package tmpl

import (
	"alert-from-db/internal/data/films"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

const (
	TemplatePathFilm string = "/web/template/film.html"
)

func CreateTemplate(tmplpath string) (*template.Template, error) {
	// Working Directory
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error get working dir %w", err)
	}
	return template.Must(template.ParseFiles(wd + tmplpath)), nil
}

// function for render template
func RenderTemplate(w http.ResponseWriter, templates *template.Template, tmpl string, f []films.Film) {
	err := templates.ExecuteTemplate(w, tmpl+".html", f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
