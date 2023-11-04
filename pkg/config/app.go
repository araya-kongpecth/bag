package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeLoadEnv() string {
	err := godotenv.Load("/Users/araya/Desktop/bagOrders/.env")
	if err != nil {
		panic("Error loading .env file")
	}
	var dbHost = os.Getenv("DB_HOST")
	var dbPort = os.Getenv("DB_PORT")
	var dbUsername = os.Getenv("DB_USERNAME")
	var dbPassword = os.Getenv("DB_PASSWORD")
	var dbName = os.Getenv("DB_NAME")

	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	return DSN
}

func Connect() {
	DSN := InitializeLoadEnv()
	d, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
