package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/araya-kongpecth/mux-miniproject/pkg/models"
	"github.com/araya-kongpecth/mux-miniproject/pkg/utils"
)

var NewBag models.Bag

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	newBags := models.GetItems()
	res, _ := json.Marshal(newBags)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetItemById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bagId := params["item_id"]
	ID, err := strconv.ParseInt(bagId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bagDetails, _ := models.GetItemById(ID)
	if bagDetails.ItemID == 0 {
		w.Write([]byte("No Item ID in Stock - cannot get by id"))
		return
	}
	res, _ := json.Marshal(bagDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateItems(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	createBag := &models.Bag{}
	utils.ParseBody(r, createBag)
	b := createBag.CreateItems()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteItems(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bagId := params["item_id"]
	ID, err := strconv.ParseInt(bagId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bagDetails, _ := models.GetItemById(ID)
	if bagDetails.ItemID == 0 {
		w.Write([]byte("No Item ID in Stock - cannot delete"))
		return
	}

	bag := models.DeleteItems(ID)
	if bag.ItemID != 0 {
		w.Write([]byte("Delete Fail :("))
		return
	}
	w.Write([]byte("Delete Successful!!\n"))
	newBags := models.GetItems()
	res, _ := json.Marshal(newBags)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateItems(w http.ResponseWriter, r *http.Request) {
	updateBag := &models.Bag{}
	utils.ParseBody(r, updateBag)
	params := mux.Vars(r)
	bagId := params["item_id"]
	ID, err := strconv.ParseInt(bagId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bagDetails, _ := models.GetItemById(ID)

	if bagDetails.ItemID == 0 {
		w.Write([]byte("No Item ID in Stock cannot update"))
		return
	}

	if updateBag.ItemName != "" {
		bagDetails.ItemName = updateBag.ItemName
		bagDetails.Price = updateBag.Price
		bagDetails.Amount = updateBag.Amount
	}

	b := bagDetails.UpdateItems(ID)

	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
