package main

import (
	"fmt"
	"net/http"
	"time"

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

	http.HandleFunc("/window/minimize", func(w http.ResponseWriter, r *http.Request) {
		templates.WindowMinimized().Render(r.Context(), w)
	})

	http.HandleFunc("/window/", func(w http.ResponseWriter, r *http.Request) {
		page := r.URL.Path[len("/window/"):]
		switch page {
		case "projects":
			templates.WindowProjects().Render(r.Context(), w)
		default:
			templates.WindowHome().Render(r.Context(), w)
		}
	})

	var startOpen bool
	http.HandleFunc("/taskbar/start", func(w http.ResponseWriter, r *http.Request) {
		startOpen = !startOpen
		if startOpen {
			templates.StartMenuOpen().Render(r.Context(), w)
		} else {
			templates.StartMenuClosed().Render(r.Context(), w)
		}
	})

	http.HandleFunc("/taskbar/clock", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().Format("3:04 PM")
		templates.Clock(t).Render(r.Context(), w)
	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
