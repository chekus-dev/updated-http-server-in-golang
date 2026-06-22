package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	HouseAddress string `json:"HouseAddress"`
}

func HomePage(w http.ResponseWriter, res *http.Request) {
	w.Header().Set("content-Type", "application/json")
	user := Message{
		ID:           1,
		Name:         "chekus",
		Email:        "chekusjoseph@yahoo.com",
		HouseAddress: "no 7 mama ramatu avenue",
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("failed to encode user: %v", err)
	}
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/user", HomePage)
	log.Println("starting server")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("failed to start application/json: %v", err)
	}
}
