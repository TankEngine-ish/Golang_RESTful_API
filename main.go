package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Id       string
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

func homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: homepage")
	fmt.Fprintf(w, "Hello, World!")
}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnAllProducts")
	json.NewEncoder(w).Encode(Products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, product := range Products {
		if string(product.Id) == key {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/products", returnAllProducts)
	myRouter.HandleFunc("/product/{id}", getProduct)
	myRouter.HandleFunc("/", homepage)
	http.ListenAndServe(":10000", myRouter)
}

func main() {

	Products = []Product{
		{Id: "1", Name: "Laptop", Quantity: 10, Price: 1000.00},
		{Id: "2", Name: "Headphones", Quantity: 50, Price: 10.00},
	}

	handleRequests()

	http.HandleFunc("/", homepage)
	http.ListenAndServe("localhost:10000", nil) // nil is the handler
}
