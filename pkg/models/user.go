package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserID    uint   `json:"user_id" gorm:"primaryKey"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func init() {
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	u.Password = HashPassword(u.Password)
	db.Create(&u)
	return u
}

func GetUsers() []User {
	var users []User
	db.Find(&users)
	return users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var user User
	db := db.Where("user_id=?", Id).Find(&user)
	return &user, db
}

func GetUserByUsername(Username string) (*User, *gorm.DB) {
	var user User
	db := db.Where("username=?", Username).Find(&user)
	return &user, db
}

func DeleteUser(Id int64) User {
	var user User
	db.Where("user_id=?", Id).Delete(user)
	return user
}

func (user *User) UpdateUser(Id int64) *User {
	db.Save(&user)
	return user
}

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
