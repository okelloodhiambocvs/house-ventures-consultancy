package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("House Ventures Backend Running...")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK - ATS System Live"))
	})

	http.ListenAndServe(":8080", nil)
}