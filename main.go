package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Kubernetes !, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
