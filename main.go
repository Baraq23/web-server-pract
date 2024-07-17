package main

import (
	"fmt"
	"net/http"
)

// helloHandler() function handles http requests and responds with 'Hello Web ;)'
func helloHandler(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "Hello Web ;)")
}

func main() {

	//Register the helloHandler function as the handler for the "/" path
	http.HandleFunc("/", helloHandler)

	// Start th http server on port 8080
	fmt.Println("Starting server on: 8080")
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
