package models

import (
	"gorm.io/gorm"

	"github.com/araya-kongpecth/mux-miniproject/pkg/config"
)

var db *gorm.DB

type Bag struct {
	ItemID   uint   `json:"item_id" gorm:"primaryKey"`
	ItemName string `json:"item_name"`
	Price    uint   `json:"price"`
	Amount   uint   `json:"amount"`
	ItemImg  string `json:"item_img"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Bag{})
}

func (b *Bag) CreateItems() *Bag {
	// db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetItems() []Bag {
	var bags []Bag
	db.Find(&bags)
	return bags
}

func GetItemById(Id int64) (*Bag, *gorm.DB) {
	var bag Bag
	db := db.Where("item_id=?", Id).Find(&bag)
	return &bag, db
}

func DeleteItems(Id int64) Bag {
	var bag Bag
	db.Where("item_id=?", Id).Delete(bag)
	return bag
}

func (bagDetails *Bag) UpdateItems(Id int64) *Bag {
	db.Save(&bagDetails)
	return bagDetails
}
