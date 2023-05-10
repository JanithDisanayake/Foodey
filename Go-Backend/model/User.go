package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID   uint `gorm:"primaryKey;unique"`
	Name string
	Age  uint
}
