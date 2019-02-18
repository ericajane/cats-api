package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Cat Struct

type Cat struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   string `json:"age"`
	Breed string `json:"breed"`
	Sex   string `json:"sex"`
}

var cats []Cat

// Get all cats

func getCats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cats)
}

func getCat(w http.ResponseWriter, r *http.Request) {

}

func createCat(w http.ResponseWriter, r *http.Request) {

}

func updateCat(w http.ResponseWriter, r *http.Request) {

}

func deleteCat(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	//Fake Data

	cats = append(cats, Cat{ID: "1", Name: "Axel", Age: "9", Breed: "Siamese", Sex: "M"})

	cats = append(cats, Cat{ID: "2", Name: "Riley", Age: "15", Breed: "Buff Tabby", Sex: "F"})

	cats = append(cats, Cat{ID: "3", Name: "Vladimir", Age: "1", Breed: "Grey Tabby", Sex: "M"})

	r.HandleFunc("/api/cats", getCats).Methods("GET")
	r.HandleFunc("/api/cats/{id}", getCat).Methods("GET")
	r.HandleFunc("/api/cats", createCat).Methods("POST")
	r.HandleFunc("/api/cats/{id}", updateCat).Methods("PUT")
	r.HandleFunc("/api/cats/{id}", deleteCat).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))

}
