package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jsur/go-web-helloworld/pkg/config"
	"github.com/jsur/go-web-helloworld/pkg/handlers"
	"github.com/jsur/go-web-helloworld/pkg/render"
)

const port = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// pass pointer to app
	// this way render package gets access to app-wide AppConfig
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", port))

	serve := &http.Server{
		Addr:    port,
		Handler: Routes(&app),
	}

	err = serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
