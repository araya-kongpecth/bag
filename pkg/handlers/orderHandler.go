package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/araya-kongpecth/mux-miniproject/pkg/database"
)

var dbOrder *gorm.DB
var errOrder error

type Order struct {
	gorm.Model
	OrderID uint `json:"order_id"`
	ItemID  uint `gorm:"references:item_id"`
}

func InitializeMigrationOrder() {

	dbOrder, errOrder = database.InitializeDB()
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

	// var item Item
	// if err := dbOrder.First(&item, order.ItemID).Error; err != nil {
	// 	// If the item does not exist, return an error message
	// 	http.Error(w, "Store does not have this item", http.StatusBadRequest)
	// 	return
	// }

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
