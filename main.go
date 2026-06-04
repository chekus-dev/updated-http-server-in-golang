package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type UserMessage struct {
	Message string `json:"messgae"`
	Reply   string `json:"Reply"`
	Answer  string `json:"message"`
	Status  int    `json:"status"`
}

func HomePage(w http.ResponseWriter, res *http.Request) {
	w.Header().Set("content-type", "application/json")

	data := UserMessage{
		Message: "welcome dear user to updated http server in go, JSON data!",
		Reply:   "thank you, JSON data!",
		Answer:  "How can we be of help, JSON data!",
		Status:  200,
	}

	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/", HomePage)
	log.Println("starting application/json on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
