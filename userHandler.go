package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbUser *gorm.DB
var errUser error

const dsnUser = "root:Ice@0623@tcp(127.0.0.1:3306)/bag?charset=utf8&parseTime=True&loc=Local"

type User struct {
	// gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func initializeMigrationUser() {
	dbUser, errUser = gorm.Open(mysql.Open(dsnUser), &gorm.Config{})
	if errUser != nil {
		fmt.Println(errUser.Error())
		panic("Cannot connect to dbUser")
	}
	dbUser.AutoMigrate(&User{})
}

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	dbUser.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	var users []User
	dbUser.Delete(&user, params["id"])
	dbUser.Find(&users)
	json.NewEncoder(w).Encode(users)

}
