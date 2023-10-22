package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {

	r := mux.NewRouter()

	//Get All Items and item by id
	r.HandleFunc("/items", GetItems).Methods("GET")
	r.HandleFunc("/items/{item_id}", GetItem).Methods("GET")

	//Get All Orders and order by id
	r.HandleFunc("/orders", GetOrders).Methods("GET")
	r.HandleFunc("/orders/{id}", GetOrder).Methods("GET")

	//Create User
	r.HandleFunc("/users", CreateUsers).Methods("POST")

	//Get JWT Token and Refresh
	r.HandleFunc("/getToken", GetToken).Methods("POST")
	r.HandleFunc("/refresh", GetRefreshToken).Methods("GET")

	//This route need to have JWT to launch function!
	protectedRoutes := r.PathPrefix("").Subrouter()
	//Middleware function check
	protectedRoutes.Use(AuthMiddleware)

	//Create, Update, Delete Items (JWT check)
	protectedRoutes.HandleFunc("/items", CreateItems).Methods("POST")
	protectedRoutes.HandleFunc("/items/{item_id}", UpdateItems).Methods("PUT")
	protectedRoutes.HandleFunc("/items/{item_id}", DeleteItems).Methods("DELETE")

	//Create, Update, Delete Orders (JWT check)
	protectedRoutes.HandleFunc("/orders", CreateOrder).Methods("POST")
	protectedRoutes.HandleFunc("/orders/{id}", UpdateOrders).Methods("PUT")
	protectedRoutes.HandleFunc("/orders/{id}", DeleteOrders).Methods("DELETE")

	//Delete User
	protectedRoutes.HandleFunc("/users/{id}", DeleteUsers).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	initializeMigration()
	initializeMigrationOrder()
	initializeMigrationUser()
	initializeRouter()
}
