package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbOrder *gorm.DB
var errOrder error

const dsnOrder = "root:Ice@0623@tcp(127.0.0.1:3306)/bag?charset=utf8&parseTime=True&loc=Local"

type Order struct {
	gorm.Model
	ItemID string
	Item   Item `gorm:"foreignKey:ItemID;references:item_id"`
}

func initializeMigrationOrder() {
	dbOrder, errOrder = gorm.Open(mysql.Open(dsnOrder), &gorm.Config{})
	if errOrder != nil {
		fmt.Println(errOrder.Error())
		panic("Cannot connect to dbOrder")
	}
	// AutoMigrate the models
	if err := dbOrder.AutoMigrate(&Order{}); err != nil {
		fmt.Println(err.Error())
		panic("Error while migrating Order model")
	}
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var orders []Order
	json.NewDecoder(r.Body).Decode(&orders)
	dbOrder.Find(&orders)
	json.NewEncoder(w).Encode(orders)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	dbOrder.First(&order, params["id"])
	json.NewEncoder(w).Encode(order)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract the ItemID from the request body (assuming it's a JSON field)
	itemID := r.FormValue("item_id")

	if itemID == "" {
		http.Error(w, "ItemID is required in the request body.", http.StatusBadRequest)
		return
	}

	// Set the ItemID in the order
	order.ItemID = itemID

	dbOrder.Create(&order)
	json.NewEncoder(w).Encode(order)
}

func UpdateOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var order Order
	dbOrder.First(&order, params["id"])
	json.NewDecoder(r.Body).Decode(&order)
	dbOrder.Save(&order)
	json.NewEncoder(w).Encode(order)
}

func DeleteOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var order Order
	var orders []Order
	dbOrder.Delete(&order, params["id"])
	dbOrder.Find(&orders)
	json.NewEncoder(w).Encode(orders)

}
