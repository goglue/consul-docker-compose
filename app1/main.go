package main

import (
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "am app1")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
