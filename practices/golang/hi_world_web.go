package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)

	http.ListenAndServe(":5000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recibiendo petición")

	io.WriteString(w, "Hi World from Web Server")
}
