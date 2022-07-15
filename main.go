package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting...")
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8787", nil))
}
