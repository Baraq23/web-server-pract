package handlers

import (
	"os"
	"fmt"
	"net/http"
	"text/template"
	"path/filepath"

 )

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path !="/about" {
		http.Error(w, "404 Page not found", http.StatusNotFound)
		fmt.Println("404 Page not found", http.StatusNotFound)
		os.Exit(1)
	}
	tmpl, err := template.ParseFiles(filepath.Join("templates", "about.html"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		os.Exit(1)
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}