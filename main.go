package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8080"

// home page handler
func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

// about page handler
func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsin template:", err)
		return
	}

}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	fmt.Println(fmt.Sprintf("Starting application on port %s", port))
	_ = http.ListenAndServe(port, nil)
}
