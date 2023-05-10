package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID    uint `gorm:"primaryKey;unique"`
	Name  string
	Desc  string
	Image string
}
