package main

import (
	"html/template"
	"net/http"
	"testing"
)

const portNumber = ":8080"

// Home page Handler
func HomeTemplate(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "homepage.tmpl")
}

// About Page handler
func AboutTemplate(w http.ResponseWriter, r *http.Request) {

}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		panic(err)
	}

}

func TestFunctionHandlerTemplate(t *testing.T) {
	http.HandleFunc("/", HomeTemplate)
	http.HandleFunc("/about", AboutTemplate)

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		panic(err)
	}
}
