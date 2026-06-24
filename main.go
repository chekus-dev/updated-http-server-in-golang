package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type personProfile struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type userProfile struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}
func WriteJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	profile := personProfile{
		ID:    2,
		Name:  "joseph",
		Email: "okaforchekus@gmail.com",
	}
	if err := json.NewEncoder(w).Encode(profile); err != nil {
		log.Printf("failed to encode Json: %v", err)
		return
	}
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	data := userProfile{
		ID:    1,
		Name:  "chekus",
		Age:   26,
		Email: "chekusjoseph@gmail.com",
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to encode Json: %v", err)
		return

	}
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", homeHandler)
	mux.HandleFunc("/person", WriteJson)
	log.Println("starting the server on port:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
