package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

type pageStruct struct {
	Content template.HTML
}

func Pages(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	filepath := "web/md/" + vars["page"] + ".md"

	html, err := readMarkdownFileAsHTML(filepath)
	if err != nil {
		serveNotFoundPage(w)
		return
	}

	var page pageStruct
	page = pageStruct{Content: template.HTML(html)}

	servePage(page, w)
}

func servePage(page pageStruct, w http.ResponseWriter) {
	t := template.New("page.html")
	t, err := t.ParseFiles("web/tmpl/page.html")
	if err != nil {
		fmt.Fprintf(w, "Failed to read page template\n")
	}
	t.Execute(w, page)
	return
}

func serveNotFoundPage(w http.ResponseWriter) {
	t := template.New("404.html")
	t, err := t.ParseFiles("web/tmpl/404.html")
	if err != nil {
		fmt.Fprintf(w, "Failed to read page template\n")
	}
	// w.WriteHeader(http.StatusNotFound)
	t.Execute(w, nil)
	return
}

func readMarkdownFileAsHTML(filename string) (htmlstring string, err error) {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("Failed to open markdown file '%s'", filename)
	}

	var buf bytes.Buffer
	if err := md.Convert(source, &buf); err != nil {
		return "", fmt.Errorf("Failed to convert markdown file '%s' to html", filename)
	}

	return buf.String(), nil
}
