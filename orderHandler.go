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
	OrderID uint `json:"order_id"`
	ItemID  uint `gorm:"references:item_id"`
}

func initializeMigrationOrder() {
	dbOrder, errOrder = gorm.Open(mysql.Open(dsnOrder), &gorm.Config{})
	if errOrder != nil {
		fmt.Println(errOrder.Error())
		panic("Cannot connect to dbOrder")
	}
	dbOrder.Migrator().CreateTable(&Order{})

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
	var orders []Order
	json.NewDecoder(r.Body).Decode(&orders)
	dbOrder.Where("order_id = ?", params["id"]).Find(&orders)
	json.NewEncoder(w).Encode(orders)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var item Item
	if err := dbOrder.First(&item, order.ItemID).Error; err != nil {
		// If the item does not exist, return an error message
		http.Error(w, "Store does not have this item", http.StatusBadRequest)
		return
	}

	// Create the order
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
