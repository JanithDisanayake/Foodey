package main

import (
	"fmt"
	"go-backend/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uint `gorm:"primaryKey;unique"`
	Name string
	Age  uint
}

type Order struct {
	gorm.Model
	ID    uint `gorm:"primaryKey;unique"`
	Name  string
	Desc  string
	Image string
}

func main() {
	r := gin.Default()
	fmt.Print(" 🚀 Server is Up and Running \n\n")

	db := model.New()

	r.GET("/users", func(context *gin.Context) {
		var users []User
		db.Find(&users)

		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.JSON(http.StatusOK, users)
	})

	r.GET("/users/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))
		var user User
		db.First(&user, "id = ?", id) // find product with integer primary key

		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.JSON(http.StatusOK, user)
	})

	r.POST("/users", func(context *gin.Context) {
		var user User
		if err := context.ShouldBind(&user); err != nil {
			context.String(http.StatusBadRequest, "bad request %v", err)
		}
		db.Create(&User{Name: user.Name, Age: user.Age})

		context.JSON(http.StatusOK, user)
	})

	r.PATCH("/users", func(context *gin.Context) {
		var user User
		if err := context.ShouldBind(&user); err != nil {
			context.String(http.StatusBadRequest, "bad request: %v", err)
		}
		db.Save(user)

		context.JSON(http.StatusOK, user)
	})

	r.DELETE("/users/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))
		var user User
		db.Take(&user, "id = ?", id)
		db.Delete(&user)

		context.String(http.StatusOK, "user deleted")
	})

	r.GET("/orders/", func(context *gin.Context) {
		var orders []Order
		db.Find(&orders)

		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.JSON(http.StatusOK, orders)
	})

	r.POST("/orders/", func(context *gin.Context) {
		var order Order
		if err := context.ShouldBind(&order); err != nil {
			context.String(http.StatusBadRequest, "bad request %v", err)
		}
		db.Create(&Order{Name: order.Name, Desc: order.Desc, Image: order.Image})

		context.JSON(http.StatusOK, order)
	})

	r.Run(":3000")
}
