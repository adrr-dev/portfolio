package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/adrr-dev/portfolio/internal/handlers"
)

func main() {
	tmpls, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	fragments, err := template.ParseGlob("templates/fragments/*.html")
	if err != nil {
		log.Fatal(err)
	}

	myHandling := handlers.NewHandling(tmpls, fragments)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("GET /{$}", myHandling.RootHandle)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
