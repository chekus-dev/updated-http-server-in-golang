package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func HomePage(w http.ResponseWriter, res *http.Request) {
	w.Header().Set("content-Type", "application/json")

	user := Message{
		ID:    1,
		Name:  "chekus-dev",
		Email: "okaforchekus@gmail.com",
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}

func main() {
	http.HandleFunc("/", HomePage)
	log.Println("starting application/json on localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
