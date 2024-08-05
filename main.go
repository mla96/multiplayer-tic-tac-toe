package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

var (
	mu sync.Mutex
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/create", createGameHandler).Methods("POST")
	r.HandleFunc("/play/{id}", playGameHandler).Methods("GET")

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

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Tic-Tac-Toe!")
}

func createGameHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	id := 0
	// instantiate new game

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"gameID": id})
}

func playGameHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// gameID := vars["id"]

	mu.Lock()
	// check if game exists with that gameID and retrieve
	mu.Unlock()

	// definition and logic to make moves
}
