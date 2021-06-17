package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	ID    uint `gorm:"primarykey"`
	Code  string
	Price uint
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Query()
	// 迁移 schema
	db.AutoMigrate(&Product{})
	if err := Create(); err != nil {
		panic(err)
	}
}

func Query() {
	result := []Product{}
	db.Find(&result)
	log.Println("result:", result)
}

func Create() (err error) {
	tx := db.Begin()
	//可用于状态机, if err == nil {} else {}
	err = tx.Create(Product{
		ID:    1,
		Code:  "1",
		Price: 10,
	}).Error
	if err != nil {
		//tx.Rollback()
		//return
	}
	err = tx.Create(Product{
		ID:    3,
		Code:  "3",
		Price: 30,
	}).Error
	if err != nil {
		tx.Rollback()
		return
	}
	return tx.Commit().Error
}
