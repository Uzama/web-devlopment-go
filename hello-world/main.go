package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Welcome to our site")
	fmt.Println(r.URL.Query())
}
