package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	err := a.Initialise(DBUser, DBPassword, "test")
	if err != nil {
		log.Fatal("An error occurred while initializing the database")
	}
	createTable()
	m.Run()

}

func createTable() {
	createTableQuery := `CREATE TABLE IF NOT EXISTS products(
		id int NOT NULL AUTO_INCREMENT,
		name varchar(255) NOT NULL,
		quantity int,
		price float(10,7),
		PRIMARY KEY (id)
	);`
	_, err := a.DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER TABLE products AUTO_INCREMENT = 1")
	log.Println("clearTable")
}

func addProduct(name string, quantity int, price float64) {
	query := fmt.Sprintf("INSERT INTO products(name, quantity, price) VALUES('%v', %v, %v)", name, quantity, price)
	_, err := a.DB.Exec(query)
	if err != nil {
		log.Println(err)

	}

	a.DB.Exec(query)
}

func TestGetProduct(t *testing.T) {
	// Clear the table
	clearTable()
	addProduct("keyboard", 100, 140)
	request, _ := http.NewRequest("GET", "/product/1", nil)
	response := sendRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)

}

func checkStatusCode(t *testing.T, expectedStatusCode int, actualStatusCode int) {
	if expectedStatusCode != actualStatusCode {
		t.Errorf("Expected status code %v but got %v", expectedStatusCode, actualStatusCode)
	}
}

func sendRequest(request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, request)
	return recorder

}

func TestCreateProduct(t *testing.T) {
	clearTable()
	var product = []byte(`{"name": "keyboard", "quantity": 1, "price": 140}`)

	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(product))

	req.Header.Set("Content-Type", "application/json")

	response := sendRequest(req)
	checkStatusCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["name"] != "keyboard" {
		t.Errorf("Expected name: %v, Got: %v", "keyboard", m["name"])
	}

	if m["quantity"] != 1.0 {
		t.Errorf("Expected quantity: %v, Got: %v", 1.0, m["quantity"])
	}

}
