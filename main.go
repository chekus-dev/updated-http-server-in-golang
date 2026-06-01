package main 

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type  UserResponse struct {
	Message string `json:"message"`
	Status int `json:"status"`
}
func HomePage(w http.ResponseWriter, res *http.Request) {
	w.Header().Set("content-type", "application/json")

	data := UserResponse{
		Message: "Welcome to updated http server in go with JSON, this is JSON data!",
		Status:200,
	} 

	json.NewEncoder(w).Encode(data)

}
func main(){
	http.HandleFunc("/", HomePage)
	fmt.Println("starting updated server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
