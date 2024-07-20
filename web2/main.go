package main

import (
	"fmt"
	"web2/handlers"
	"log"
	"net/http"
)

func main() {
	fileserver := http.FileServer(http.Dir("./templates"))

	http.Handle("/", fileserver)
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/form", handlers.FormHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	if err := http.ListenAndServe(":8100", nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Starting server at port 8100")
}
