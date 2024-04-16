package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID    uint `gorm:"primaryKey;unique"`
	Name  string
	Desc  string
	Image string
}

type OrderInterface interface {
	SaveOrder(order *Order)
	FindAllOrders() []*Order
}

func SaveOrder(db *gorm.DB, order *Order) {
	db.Create(&order)
}

func FindAllOrders(db *gorm.DB) []*Order {
	var orders []*Order
	db.Find(&orders)
	return orders
}
