package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Costumer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
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

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(costumers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(costumers)

	}

}
