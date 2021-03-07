package main

import (
	"fmt"
	"net/http"
)

// home page handler
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// about page handler
func about(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("About page, sum is %d", sum))
}

func addValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	_ = http.ListenAndServe(":8080", nil)
}
