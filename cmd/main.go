package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/araya-kongpecth/mux-miniproject/pkg/routes"
)

// import (
// 	"log"
// 	"net/http"
//
//
//

// 	"github.com/gorilla/mux"

// 	"github.com/araya-kongpecth/mux-miniproject/handlers"
// 	"github.com/araya-kongpecth/mux-miniproject/middleware"
// )

// func initializeRouter() {

// 	r := mux.NewRouter()

// 	//Get All Items and item by id
// 	r.HandleFunc("/items", handlers.GetItems).Methods("GET")
// 	r.HandleFunc("/items/{item_id}", handlers.GetItem).Methods("GET")

// 	//Get All Orders and order by id
// 	r.HandleFunc("/orders", handlers.GetOrders).Methods("GET")
// 	r.HandleFunc("/orders/{id}", handlers.GetOrder).Methods("GET")

// 	//Create User
// 	r.HandleFunc("/users", handlers.CreateUsers).Methods("POST")

// 	//Get JWT Token and Refresh
// 	r.HandleFunc("/getToken", handlers.GetToken).Methods("POST")
// 	r.HandleFunc("/refresh", handlers.GetRefreshToken).Methods("GET")

// 	//This route need to have JWT to launch function!
// 	protectedRoutes := r.PathPrefix("").Subrouter()
// 	//Middleware function check
// 	protectedRoutes.Use(middleware.AuthMiddleware)

// 	//Create, Update, Delete Items (JWT check)
// 	protectedRoutes.HandleFunc("/items", handlers.CreateItems).Methods("POST")
// 	protectedRoutes.HandleFunc("/items/{item_id}", handlers.UpdateItems).Methods("PUT")
// 	protectedRoutes.HandleFunc("/items/{item_id}", handlers.DeleteItems).Methods("DELETE")

// 	//Create, Update, Delete Orders (JWT check)
// 	protectedRoutes.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
// 	protectedRoutes.HandleFunc("/orders/{id}", handlers.UpdateOrders).Methods("PUT")
// 	protectedRoutes.HandleFunc("/orders/{id}", handlers.DeleteOrders).Methods("DELETE")

// 	//Delete User
// 	protectedRoutes.HandleFunc("/users/{id}", handlers.DeleteUsers).Methods("DELETE")

// 	log.Fatal(http.ListenAndServe(":9000", r))
// }

// func main() {
// 	handlers.InitializeMigration()
// 	handlers.InitializeMigrationOrder()
// 	handlers.InitializeMigrationUser()
// 	initializeRouter()
// }

func main() {
	c := cors.AllowAll()
	r := mux.NewRouter()
	routes.RegisterBagStoreRoutes(r)
	http.Handle("/", r)
	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe("localhost:9000", handler))
}
