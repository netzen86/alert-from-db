package main

import (
	"alert-from-db/internal/frontend"
	"alert-from-db/internal/router"
	"alert-from-db/internal/tmpl"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go app... on port", frontend.HTTPport)

	gCli, err := frontend.GetgRPCCli()
	if err != nil {
		log.Fatalf("error when getting gRPC client %v", err)
	}

	template, err := tmpl.CreateTemplate(tmpl.TemplatePathFilm)
	if err != nil {
		log.Fatalf("error when create template %v", err)
	}

	// получаем роутер
	gw := router.GetGateway(gCli, template, "film")
	httpServer := &http.Server{Addr: ":" + frontend.HTTPport, Handler: gw}

	err = httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("error when start server %v ", err)
	}
}
