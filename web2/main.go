package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Page not found", http.StatusNotFound)
		fmt.Println("404 Page not found", http.StatusNotFound)

		return
	} else {
		fmt.Println(http.StatusAccepted, "Path Success")
	}

	if r.Method != "GET" {
		http.Error(w, "404 invalid Request", http.StatusNotFound)
		fmt.Println("404 Invlid Request", http.StatusNotFound)
		return
	} else {
		fmt.Println(http.StatusAccepted, "Request Accepted")
	}

	fmt.Fprintf(w, "hello there")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() Error: %v", err)
			return
		}
		fName := r.FormValue("fName")
		addr := r.FormValue("address")
		fmt.Fprintf(w, "Post request successful\n")
		fmt.Fprintf(w, "First Name: %v\n", fName)
		fmt.Fprintf(w, "Address: %v\n", addr)
		return
	}

	tmpl, err := template.ParseFiles(filepath.Join("static", "form.html"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}


func main() {
	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	if err := http.ListenAndServe(":8100", nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting server at port 8100")
}
