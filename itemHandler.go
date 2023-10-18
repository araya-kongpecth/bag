package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DSN = "root:Ice@0623@tcp(127.0.0.1:3306)/bag?charset=utf8&parseTime=True&loc=Local"

type Item struct {
	ItemID   string `json:"item_id"`
	ItemName string `json:"item_name"`
	Amount   int    `json:"amount"`
}

func initializeMigration() {
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
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
