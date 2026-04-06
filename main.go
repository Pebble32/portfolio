package main

import (
	"fmt"
	"net/http"

	"github.com/Pebble32/portfolio/templates"
	"github.com/a-h/templ"
)

func main(){
	component := templates.Hello("Teset")
	http.Handle("/", templ.Handler(component))

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":3000", nil)
	fmt.Println("Listening on port 3000")
}
