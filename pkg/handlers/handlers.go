package handlers

import (
	"net/http"

	"github.com/jsur/go-web-helloworld/pkg/render"
)

// Home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
