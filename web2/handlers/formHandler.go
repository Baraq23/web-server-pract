package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

)

type WebData struct {
	ArtStr string
}

var data WebData

func FormHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside formHandler")

	tmpl, err := template.ParseFiles(filepath.Join("templates", "index.html"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		os.Exit(1)
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
