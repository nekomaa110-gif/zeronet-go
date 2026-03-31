package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ZERO NET BACKEND")
	})
	fmt.Println("Server Jalan di :8080")
	http.ListenAndServe(":8080", nil)
}