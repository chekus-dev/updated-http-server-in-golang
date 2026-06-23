package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", homeHandler)
	mux.HandleFunc("/person", personHanlder)
	log.Println("starting the server on port:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}

func personHanlder(w http.ResponseWriter, res *http.Request) {
	w.Header().Set("content-Type", "application/json")

	coding := personProfile{
		ID:    2,
		Name:  "joseph",
		Email: "okaforchekus@gmail.com",
	}
	if err := json.NewEncoder(w).Encode(coding); err != nil {
		log.Printf("failed to encode coding: %v", err)
	}
}


type personProfile struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}


func homeHandler(w http.ResponseWriter, res *http.Request) {
	w.Header().Set("content-Type", "application/json")

	data := userProfile{
		ID:    1,
		Name:  "chekus",
		Age:   26,
		Email: "chekusjoseph@gmail.com",
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to encode data: %v", err)

	}
}

type userProfile struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}
