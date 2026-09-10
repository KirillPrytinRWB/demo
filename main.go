package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"sort"
)

func main() {
	// Handler for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!<br/>/version: 0.0.1a")
	})

	h2 := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	}

	h3 := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "It's alive, %q", html.EscapeString(r.URL.Path))
	}

	http.HandleFunc("/endpoint", h2)
	http.HandleFunc("/pax", h3)

	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		env := os.Environ()
		sort.Strings(env)
		for _, e := range env {
			fmt.Fprintln(w, e)
		}
	})

	fmt.Println("Starting...")
	fmt.Println("Server listening on port 8080...branched")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
