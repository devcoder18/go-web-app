package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go Server: Request received on: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":3000", nil)
}
