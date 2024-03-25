# Golang_RESTful_API

This was my first API in Golang and it's a simple RESTful API for managing an inventory of products. It also uses MySQL for data storage.

## Project Structure
The project is divided into several files:

``app.go``: This file contains the main application logic. It sets up the HTTP server, connects to the database, and defines the routes for the API.

``app_test.go``: This file contains unit tests for the application. It tests the API endpoints and database operations.

``model.go``: This file defines the ``product`` struct and the methods for performing CRUD operations on the products in the database.

``main.go``: This is the entry point of the application. It initializes and runs the app.

``constants.go``: This file contains the constants for the database name, user, and password.

## API Endpoints
The API provides the following endpoints:

``GET /products``: Fetch all products.
``GET /product/{id}``: Fetch a single product by its ID.
``POST /product``: Create a new product.
``PUT /product/{id}``: Update an existing product by its ID.
``DELETE /product/{id}``: Delete a product by its ID.

## Running the Tests
To run the tests, use the ``go test`` command in the root directory of the project.

## Running the Application
To run the application, use the ``go run main.go`` command in the root directory of the project. The application will start a server on port 10000.

## Database Configuration
The application connects to a MySQL database. The database name, user, and password are defined in the ``constants.go`` file. You can change these values to match your own database configuration.

## Dependencies
This project uses the following Go packages:

``database/sql``: For interacting with the database.
``github.com/go-sql-driver/mysql``: The MySQL driver for Go's database/sql package.
``github.com/gorilla/mux``: A powerful HTTP router and URL matcher for building Go web servers.
``net/http``: For building HTTP servers and clients.
``encoding/json``: For encoding and decoding JSON.

## Deployment
The API was later deployed via GitHub Actions.

## Notes

I can say that it's kinda nice not having to install a ton of third party packages and dependencies just to create a simple API like in Javascript. This is my fourth API project by now.
