package main

import (
	"fmt"
	"log"
	"net/http"
)

// set to false to fail the build
func PassTest() bool {
	return true
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are here: %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
