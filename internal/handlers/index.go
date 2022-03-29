package handlers

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request) {
	t := template.New("index.html")
	t, _ = t.ParseFiles("web/tmpl/index.html")
	t.Execute(w, nil)
}
