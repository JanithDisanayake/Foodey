package main

import (
	"fmt"
	"go-backend/controllers/graphql"
	"go-backend/controllers/rest"

	"github.com/gin-contrib/cors"

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

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	r.Use(cors.Default())
	fmt.Print(" 🚀 Server is Up and Running \n\n")

	// user endpoints
	r.GET("/users", rest.GetAllUsers)
	r.GET("/users/:id", rest.GetUserById)
	r.POST("/users", rest.CreateUser)
	r.PATCH("/users", rest.UpdateUser)
	r.DELETE("/users/:id", rest.RemoveUser)

	// order endpoints
	r.GET("/orders/", rest.GetAllOrders)
	r.POST("/orders/", rest.CreateOrder)

	// graphql endpoints
	r.POST("/query", graphql.GraphqlHandler())
	r.GET("/", graphql.PlaygroundHandler())

	r.Run(":3000")
}
