package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/jsur/go-web-helloworld/pkg/config"
	"github.com/jsur/go-web-helloworld/pkg/handlers"
)

// Routes contains app routes
func Routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
