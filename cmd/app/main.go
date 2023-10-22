package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func initializeENV() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func InitializeDB() (*gorm.DB, error) {
	initializeENV()
	var dbHost = os.Getenv("DB_HOST")
	var dbPort = os.Getenv("DB_PORT")
	var dbUsername = os.Getenv("DB_USERNAME")
	var dbPassword = os.Getenv("DB_PASSWORD")
	var dbName = os.Getenv("DB_NAME")

	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}

func main() {
	initializeMigration()
	initializeMigrationOrder()
	initializeMigrationUser()
	initializeRouter()
}
