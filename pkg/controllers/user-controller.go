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

var NewUser models.User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	newUser := models.GetUsers()
	res, _ := json.Marshal(newUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["user_id"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetails, _ := models.GetUserById(ID)
	if userDetails.UserID == 0 {
		w.Write([]byte("No User ID regis"))
		return
	}
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	createUser := &models.User{}
	utils.ParseBody(r, createUser)
	u := createUser.CreateUser()
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["item_id"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	userDetails, _ := models.GetUserById(ID)
	if userDetails.UserID == 0 {
		w.Write([]byte("No User ID regis"))
		return
	}

	user := models.DeleteUser(ID)
	if user.UserID != 0 {
		w.Write([]byte("Delete User ID Fail :("))
		return
	}
	w.Write([]byte("Delete User ID  Successful!!\n"))
	newUser := models.GetUsers()
	res, _ := json.Marshal(newUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	updateUser := &models.User{}
	utils.ParseBody(r, updateUser)
	params := mux.Vars(r)
	userId := params["user_id"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	userDetails, _ := models.GetUserById(ID)

	if userDetails.UserID == 0 {
		w.Write([]byte("No User ID"))
		return
	}

	if updateUser.FirstName != "" && updateUser.LastName != "" && updateUser.Email != "" {
		userDetails.FirstName = updateUser.FirstName
		userDetails.LastName = updateUser.LastName
		userDetails.Email = updateUser.Email
	}

	u := userDetails.UpdateUser(ID)

	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
