package main

import "net/http"


func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("helooooooooooooooo web!"))
	})

	http.HandleFunc("/hello2", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("this is another page!"))
	})
	http.ListenAndServe(":8100",nil)

}