package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, request *http.Request) {

	if request.URL.Path != "/" {
		http.NotFound(w, request)
		return
	}

	files := []string{
		"./ui/html/layouts/app.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "app", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}

func displaySnippet(writer http.ResponseWriter, request *http.Request) {

	id, err := strconv.Atoi(request.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(writer, request)
		return
	}

	fmt.Fprintf(writer, "Display a specific snippet %d", id)
}

func createSnippet(writer http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		writer.Header().Set("Allow", "POST")
		writer.Header().Set("Content-Type", "application/json")
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	writer.Write([]byte("Create a new snippet"))
}
