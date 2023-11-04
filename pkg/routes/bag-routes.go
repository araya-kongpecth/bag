package routes

import (
	"github.com/gorilla/mux"

	"github.com/araya-kongpecth/mux-miniproject/pkg/controllers"
)

var RegisterBagStoreRoutes = func(r *mux.Router) {
	//Get All Items and item by id
	r.HandleFunc("/items", controllers.GetItems).Methods("GET")
	r.HandleFunc("/items/{item_id}", controllers.GetItemById).Methods("GET")

	r.HandleFunc("/items", controllers.CreateItems).Methods("POST")
	r.HandleFunc("/items/{item_id}", controllers.UpdateItems).Methods("PUT")
	r.HandleFunc("/items/{item_id}", controllers.DeleteItems).Methods("DELETE")

	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")

	r.HandleFunc("/login", controllers.GetToken).Methods("POST")
	r.HandleFunc("/refresh", controllers.GetRefreshToken).Methods("GET")

	/*
		//Create User
		r.HandleFunc("/users", handlers.CreateUsers).Methods("POST")

		//Get JWT Token and Refresh
		r.HandleFunc("/login", handlers.GetToken).Methods("POST")
		r.HandleFunc("/refresh", handlers.GetRefreshToken).Methods("GET")

		//This route need to have JWT to launch function!
		protectedRoutes := r.PathPrefix("").Subrouter()
		//Middleware function check
		protectedRoutes.Use(middleware.AuthMiddleware)

		//Create, Update, Delete Items (JWT check)
		protectedRoutes.HandleFunc("/items", controllers.CreateItems).Methods("POST")
		protectedRoutes.HandleFunc("/items/{item_id}", controllers.UpdateItems).Methods("PUT")
		protectedRoutes.HandleFunc("/items/{item_id}", controllers.DeleteItems).Methods("DELETE")

	*/

}
