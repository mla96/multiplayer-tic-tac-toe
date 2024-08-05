package main

import (
	"fmt"
	"net/http"
)

func main() {
	// register handler for root URL "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("test")
	})

	// start HTTP server
	fmt.Println("Now running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
