package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialise(DBUser string, DBPassword string, DBName string) error {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DBUser, DBPassword, DBName)
	var err error
	app.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	app.Router = mux.NewRouter().StrictSlash(true)
	app.handleRoutes()
	return nil // returns nil if no error
}

func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}

func sendResponse(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func sendError(w http.ResponseWriter, status int, err string) {
	error_message := map[string]string{"error": err}
	sendResponse(w, status, error_message)
}

func (app *App) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := getProducts(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, products)
}

func (app *App) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	p := product{ID: key}
	err = p.getProduct(app.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			sendError(w, http.StatusNotFound, "Product not found")
		default:
			sendError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	sendResponse(w, http.StatusOK, p)
}

func (app *App) createProduct(w http.ResponseWriter, r *http.Request) {
	var p product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err = p.createProduct(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusCreated, p)

}

func (app *App) updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	var p product
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	p.ID = key
	err = p.updateProduct(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, p)

}

func (app *App) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	p := product{ID: key}
	err = p.deleteProduct(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, map[string]string{"result": "successfully deleted"})
}

func (app *App) handleRoutes() {
	app.Router.HandleFunc("/products", app.getProducts).Methods("GET")
	app.Router.HandleFunc("/product/{id}", app.getProduct).Methods("GET")
	app.Router.HandleFunc("/product", app.createProduct).Methods("POST")
	app.Router.HandleFunc("/product/{id}", app.updateProduct).Methods("PUT")
	app.Router.HandleFunc("/product/{id}", app.deleteProduct).Methods("DELETE")
}

// ******** OLD CODE ********
//
// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql" // _ means we are not using this package directly
// )

// func checkError(e error) {
// 	if e != nil {
// 		log.Fatalln(e)
// 	}
// }

// type Data struct {
// 	id   int
// 	name string
// }

// func main() {
// 	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DBUser, DBPassword, DBName)
// 	db, err := sql.Open("mysql", connectionString)
// 	checkError(err)
// 	defer db.Close()

// 	result, err := db.Exec("insert into data values(5, 'xyz')")
// 	checkError(err)
// 	lastInsertedId, err := result.LastInsertId()
// 	fmt.Println("lastInsertedId: ", lastInsertedId)
// 	checkError(err)
// 	rowsAffected, err := result.RowsAffected()
// 	fmt.Println("rowsAffected: ", rowsAffected)
// 	checkError(err)

// 	rows, err := db.Query("SELECT * FROM data")
// 	checkError(err)

// 	for rows.Next() {
// 		var data Data
// 		err = rows.Scan(&data.id, &data.name)
// 		checkError(err)
// 		fmt.Println(data)

// 	}
// } //

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
