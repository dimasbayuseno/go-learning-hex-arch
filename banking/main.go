package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Costumer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func main() {

	// define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/getAllCostumers", getAllCostumer)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func getAllCostumer(w http.ResponseWriter, r *http.Request) {
	costumers := []Costumer{
		{Name: "Dimas", City: "Jogja", Zipcode: "12345"},
		{Name: "Soultan", City: "Jogja", Zipcode: "12345"},
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(costumers)
}
