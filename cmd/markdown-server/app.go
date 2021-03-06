package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/greenbene/markdown-server/internal/handlers"
)

func main() {
	log.Println("Starting up server")

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Index)
	r.HandleFunc("/{page}", handlers.Pages)
	http.Handle("/", r)

	log.Println("Listing for requests at http://0.0.0.0:8000/")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
