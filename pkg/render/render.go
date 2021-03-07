package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// a map of custom functions that are needed in a template
var functions = template.FuncMap{}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	_, err := RenderTemplateTest(w)
	if err != nil {
		fmt.Println("Error getting template cache:", err)
	}

	/*
		parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
		err := parsedTemplate.Execute(w, nil)
		if err != nil {
			fmt.Println("error parsing template:", err)
			return
		}
	*/
}

func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currently:", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts
	}

	return cache, nil
}
