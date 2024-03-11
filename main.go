package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, request *http.Request) {

	if request.URL.Path != "/" {
		http.NotFound(w, request)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", displaySnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
