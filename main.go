package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // _ means we are not using this package directly
)

func checkError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

type Data struct {
	id   int
	name string
}

func main() {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DBUser, DBPassword, DBName)
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM data")
	checkError(err)

	for rows.Next() {
		var data Data
		err = rows.Scan(&data.id, &data.name)
		checkError(err)
		fmt.Println(data)

	}
}

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// type Product struct {
// 	Id       string
// 	Name     string
// 	Quantity int
// 	Price    float64
// }

// var Products []Product

// func homepage(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Endpoint Hit: homepage")
// 	fmt.Fprintf(w, "Hello, World!")
// }

// func returnAllProducts(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Endpoint Hit: returnAllProducts")
// 	json.NewEncoder(w).Encode(Products)
// }

// func getProduct(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key := vars["id"]
// 	for _, product := range Products {
// 		if string(product.Id) == key {
// 			json.NewEncoder(w).Encode(product)
// 		}
// 	}
// }

// func handleRequests() {
// 	myRouter := mux.NewRouter().StrictSlash(true)
// 	myRouter.HandleFunc("/products", returnAllProducts)
// 	myRouter.HandleFunc("/product/{id}", getProduct)
// 	myRouter.HandleFunc("/", homepage)
// 	http.ListenAndServe(":10000", myRouter)
// }

// func main() {

// 	Products = []Product{
// 		{Id: "1", Name: "Laptop", Quantity: 10, Price: 1000.00},
// 		{Id: "2", Name: "Headphones", Quantity: 50, Price: 10.00},
// 	}

// 	handleRequests()

// 	http.HandleFunc("/", homepage)
// 	http.ListenAndServe("localhost:10000", nil) // nil is the handler
// }
