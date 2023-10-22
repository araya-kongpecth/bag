package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/araya-kongpecth/mux-miniproject/database"
)

var DB *gorm.DB
var err error

type Item struct {
	ItemID   uint   `json:"item_id" gorm:"primaryKey"`
	ItemName string `json:"item_name"`
	Price    uint   `json:"price"`
	Amount   uint   `json:"amount"`
}

func InitializeMigration() {

	DB, err = database.InitializeDB()
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Item{})
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var items []Item
	DB.Find(&items)
	json.NewEncoder(w).Encode(items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var item Item
	DB.First(&item, params["item_id"])
	json.NewEncoder(w).Encode(item)
}

func CreateItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	json.NewDecoder(r.Body).Decode(&item)
	DB.Create(&item)
	json.NewEncoder(w).Encode(item)
}

func UpdateItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var item Item
	DB.First(&item, params["item_id"])
	json.NewDecoder(r.Body).Decode(&item)
	DB.Save(&item)
	json.NewEncoder(w).Encode(item)
}

func DeleteItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var item Item
	var items []Item
	DB.Delete(&item, params["item_id"])
	DB.Find(&items)
	json.NewEncoder(w).Encode(items)

}
