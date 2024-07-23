package handlers

import (
	"os"
	"fmt"
	"net/http"
	

 )


 func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Page not found", http.StatusNotFound)
		fmt.Println("404 Page not found", http.StatusNotFound)
		os.Exit(1)
	} else {
		fmt.Println(http.StatusAccepted, "Path Success")
	}

	if r.Method != "GET" {
		http.Error(w, "404 invalid Request", http.StatusNotFound)
		fmt.Println("404 Invlid Request", http.StatusNotFound)
		os.Exit(1)
	} else {
		fmt.Println(http.StatusAccepted, "Request Accepted")
	}

	fmt.Fprintf(w, "hello\nthere")
}
