package main

import (
	"fmt"
	"net/http"

	"github.com/jsur/go-web-helloworld/pkg/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", port))
	_ = http.ListenAndServe(port, nil)
}
