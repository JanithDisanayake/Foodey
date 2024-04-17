package main

import (
	"fmt"
	"go-backend/controllers/graphql"
	"go-backend/controllers/rest"
	"go-backend/models"
	"net/http"

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

	db := models.New()

	r.GET("/users", rest.GetAllUsers)
	r.GET("/users/:id", rest.GetUserById)
	r.POST("/users", rest.CreateUser)
	r.PATCH("/users", rest.UpdateUser)
	r.DELETE("/users/:id", rest.RemoveUser)

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

	r.POST("/query", graphql.GraphqlHandler())
	r.GET("/", graphql.PlaygroundHandler())

	r.Run(":3000")
}
