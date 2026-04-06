package main

import (
	"fmt"
	"net/http"

	"github.com/Pebble32/portfolio/templates"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HX-Request") == "true" {
			templates.HomePage().Render(r.Context(), w)
		} else {
			templates.HomeFullPage().Render(r.Context(), w)
		}
	})

	http.HandleFunc("/projects/", func(w http.ResponseWriter, r *http.Request) {
		slug := r.URL.Path[len("/projects/"):]
		if slug != "" {
			p, ok := templates.GetProjectBySlug(slug)
			if !ok {
				http.NotFound(w, r)
				return
			}
			if r.Header.Get("HX-Request") == "true" {
				templates.ProjectDetailPage(p).Render(r.Context(), w)
			} else {
				templates.ProjectDetailFullPage(p).Render(r.Context(), w)
			}
			return
		}
		if r.Header.Get("HX-Request") == "true" {
			templates.ProjectsPage().Render(r.Context(), w)
		} else {
			templates.ProjectsFullPage().Render(r.Context(), w)
		}
	})

	http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HX-Request") == "true" {
			templates.ProjectsPage().Render(r.Context(), w)
		} else {
			templates.ProjectsFullPage().Render(r.Context(), w)
		}
	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
