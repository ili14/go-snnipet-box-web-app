// handlers.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice.
	files := []string{
		"./ui/html/base.html",
		"./ui/html/pages/home.html",
		"./ui/html/partials/nav.html",
	}
	// Use the template.ParseFiles() function to read the files and store the
	// templates in a template set. Notice that we can pass the slice of file
	// paths as a variadic parameter?
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// Use the ExecuteTemplate() method to write the content of the "base"
	// template as the response body.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
