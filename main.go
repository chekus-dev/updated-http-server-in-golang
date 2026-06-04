package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
	Reply   string `json:"Reply"`
	Answer  string `json:"answer"`
	Status  int    `json:"status"`
}

func HomePage(w http.ResponseWriter, res *http.Request) {
	w.Header().Set("content-type", "application/json")

	data := Message{
		Message: "welcome dear user to updated http server in go",
		Reply:   "thank you",
		Answer:  "How can we be of help",
		Status:  200,
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}

func main() {
	http.HandleFunc("/", HomePage)
	log.Println("starting application/json on localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server failed to stsrt: %v", err)
	}
}
