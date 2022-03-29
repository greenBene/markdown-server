package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type page struct {
	Content string
}

func Pages(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	log.Println("Opening Page: ", vars["page"])

	t := template.New("page.html")
	t, _ = t.ParseFiles("web/tmpl/page.html")
	t.Execute(w, nil)
}
