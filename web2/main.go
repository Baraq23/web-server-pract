package main

import (
	"fmt"
	
	"web2/handlers"
	"log"
	"net/http"
)

type pageData struct {
	AsciiArt string
}

type formData struct {
	Input string
	Banner string
}


func main() {
	fileserver := http.FileServer(http.Dir("./templates"))

	http.Handle("/", fileserver)
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/form", handlers.FormHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	fmt.Println("Starting server at port 8100")
	fmt.Println("Page running on http://localhost:8100")
	if err := http.ListenAndServe(":8100", nil); err != nil {
		log.Fatal(err)
	}

}
