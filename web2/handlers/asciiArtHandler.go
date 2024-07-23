package handlers

import (
	"fmt"
	"net/http"
	"os"
	asciiart "ascii-art-web/functions"
	"path/filepath"
	"text/template"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside asciiArtHAndler")
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "404 Page not found", http.StatusNotFound)
		fmt.Println("404 Page not found", http.StatusNotFound)
		os.Exit(1)
	}
	if r.Method == http.MethodPost {
		fmt.Println("Inside method: ", http.MethodPost)
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() Error: %v", err)
			os.Exit(1)
		}
		input := r.FormValue("tArea")
		banner := r.FormValue("banner")
		fmt.Println(input, banner)
		art := asciiart.WordDistributor(input, banner)
		data = WebData{
			ArtStr: art,
		}
		fmt.Println("input & banner", input, banner)
		fmt.Println("Data: \n", data)
	}
	fmt.Println("outside Method")
	tmpl, err := template.ParseFiles(filepath.Join("templates", "ascii-art.html"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		os.Exit(1)
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
