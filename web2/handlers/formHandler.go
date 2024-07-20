package handlers

import (
	"os"
	"fmt"
	"net/http"
	"text/template"
	"path/filepath"

 )

 func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() Error: %v", err)
			os.Exit(1)
		}
		fName := r.FormValue("fName")
		addr := r.FormValue("address")
		fmt.Fprintf(w, "Post request successful\n")
		fmt.Fprintf(w, "First Name: %v\n", fName)
		fmt.Fprintf(w, "Address: %v\n", addr)
		os.Exit(1)
	}

	tmpl, err := template.ParseFiles(filepath.Join("static", "form.html"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		os.Exit(1)
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}