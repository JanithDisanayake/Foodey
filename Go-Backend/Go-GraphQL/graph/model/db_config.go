package model

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New() *gorm.DB{

	db, err := gorm.Open(sqlite.Open("../foodey.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	} else {
		fmt.Printf(" 🎯 Database is created \n\n")
	}
	db.AutoMigrate(&User{})

	return db
}
