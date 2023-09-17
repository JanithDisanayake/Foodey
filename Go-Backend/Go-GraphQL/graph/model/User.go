package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uint `gorm:"primaryKey;unique"`
	Name string
	Age  uint
}

type UserInterface interface {
	SaveUser(user *User)
	FindAllUsers() []*User
}


func SaveUser (db *gorm.DB, user *User) {
	db.Create(&user)
}

func FindAllUsers (db *gorm.DB) []*User {
	var users []*User
	db.Find(&users)
	return users
}