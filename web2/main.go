package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii-art-web/handlers"
)

// type pageData struct {
// 	AsciiArt string
// }

const gv int = 9

func main() {
	// fileserver := http.FileServer(http.Dir("./templates"))

	http.HandleFunc("/", handlers.FormHandler)
	http.HandleFunc("/hello", handlers.HelloHandler)
	// http.HandleFunc("/form", fileserver)
	http.HandleFunc("/ascii-art", handlers.AsciiArtHandler)
	// http.HandleFunc("/results", handlers.ResultsHandler)

	fmt.Println("Starting server at port 8100")
	fmt.Println("Page running on http://localhost:8100")
	if err := http.ListenAndServe(":8100", nil); err != nil {
		log.Fatal(err)
	}
}
