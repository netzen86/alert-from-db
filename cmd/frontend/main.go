package main

import (
	"alert-from-db/internal/data/films"
	"alert-from-db/internal/router"
	"alert-from-db/internal/tmpl"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := "8000"
	fmt.Println("Go app... on port", port)

	f := []films.Film{
		{ID: 0, Title: "The Godfather", Director: "Francis Ford Coppola"},
		{ID: 1, Title: "Blade Runner", Director: "Ridley Scott"},
		{ID: 2, Title: "The Thing", Director: "John Carpenter"},
	}

	template, err := tmpl.CreateTemplate(tmpl.TemplatePathFilm)
	if err != nil {
		log.Fatalf("error when create template %v", err)
	}

	// получаем роутер
	gw := router.GetGateway(&f, template, "film")
	httpServer := &http.Server{Addr: ":" + port, Handler: gw}

	err = httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("error when start server %v ", err)
	}
}
